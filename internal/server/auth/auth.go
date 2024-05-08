package auth

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	"github.com/ramonberrutti/jumpshot-ext-authz/internal/storage"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	v31 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	authpb "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	typev3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"

	storagepb "github.com/ramonberrutti/jumpshot/protogen/storage"
)

type AuthorizationServiceServer struct {
	Store storage.Storage
}

func cookieHeader(rawCookies, filter string) (*http.Cookie, bool) {
	if rawCookies == "" {
		return nil, false
	}

	header := http.Header{}
	header.Add("Cookie", rawCookies)
	req := http.Request{Header: header}

	cookie, err := req.Cookie(filter)
	if err != nil {
		return nil, false
	}

	return cookie, true
}

func getToken(headers map[string]string) string {
	authorization := headers["authorization"]
	cookie, cookieExist := cookieHeader(headers["cookie"], "token")

	if cookieExist {
		return cookie.Value
	} else if len(authorization) > 7 && strings.HasPrefix(strings.ToLower(authorization[0:7]), "bearer ") {
		return authorization[7:]
	}

	return ""
}

func (a *AuthorizationServiceServer) Check(ctx context.Context, req *authpb.CheckRequest) (*authpb.CheckResponse, error) {
	headers := req.GetAttributes().GetRequest().GetHttp().GetHeaders()
	token := getToken(headers)
	var md *structpb.Struct

	authHeaders := []*v31.HeaderValueOption{}
	if token != "" {
		if session, err := a.Store.Sessions().Get(ctx, token); err == nil {
			// TODO(ramonberrutti): move to a new function to validate session.
			if session.ExpireTime.AsTime().Before(time.Now().UTC()) {
				return &authpb.CheckResponse{
					HttpResponse: &authpb.CheckResponse_DeniedResponse{
						DeniedResponse: &authpb.DeniedHttpResponse{
							Status: &typev3.HttpStatus{
								Code: typev3.StatusCode_Unauthorized,
							},
							Body: "Session expired",
						},
					},
				}, nil
			}

			authHeaders = append(authHeaders, &v31.HeaderValueOption{
				Header: &v31.HeaderValue{
					Key:   "x-extrapolate-id",
					Value: session.ExtrapolateId,
				},
				Append: wrapperspb.Bool(false),
			})

			// Include session metadata.
			// TODO(ramonberrutti): validate use case for this.
			md, _ = structpb.NewStruct(map[string]any{
				"session": session.GetMetadata(),
			})

			switch m := session.Method.(type) {
			case *storagepb.Session_User:
				userBinary, _ := proto.Marshal(m.User)
				authHeaders = append(authHeaders,
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:   "x-user-id",
							Value: m.User.Id,
						},
						Append: wrapperspb.Bool(false),
					},
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:   "x-user",
							Value: base64.StdEncoding.EncodeToString(userBinary),
						},
						Append: wrapperspb.Bool(false),
					},
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:      "x-user-bin",
							RawValue: userBinary,
						},
						Append: wrapperspb.Bool(false),
					},
				)
			case *storagepb.Session_Client:
				clientBinary, _ := proto.Marshal(m.Client)
				authHeaders = append(authHeaders,
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:   "x-client-id",
							Value: m.Client.Id,
						},
						Append: wrapperspb.Bool(false),
					},
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:   "x-client",
							Value: base64.StdEncoding.EncodeToString(clientBinary),
						},
						Append: wrapperspb.Bool(false),
					},
				)
			default:
				authHeaders = append(authHeaders,
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:   "x-user-id",
							Value: "",
						},
						Append: wrapperspb.Bool(false),
					},
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:   "x-user",
							Value: "",
						},
						Append: wrapperspb.Bool(false),
					},
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:   "x-client-id",
							Value: "",
						},
						Append: wrapperspb.Bool(false),
					},
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:   "x-client",
							Value: "",
						},
						Append: wrapperspb.Bool(false),
					},
				)
			}

			switch m := session.Impersonate.(type) {
			case *storagepb.Session_AsUser:
				userBinary, _ := proto.Marshal(m.AsUser)
				authHeaders = append(authHeaders,
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:   "x-as-user-id",
							Value: m.AsUser.Id,
						},
						Append: wrapperspb.Bool(false),
					},
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:   "x-as-user",
							Value: base64.StdEncoding.EncodeToString(userBinary),
						},
						Append: wrapperspb.Bool(false),
					},
				)
			case *storagepb.Session_AsClient:
				clientBinary, _ := proto.Marshal(m.AsClient)
				authHeaders = append(authHeaders,
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:   "x-as-client-id",
							Value: m.AsClient.Id,
						},
						Append: wrapperspb.Bool(false),
					},
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:   "x-as-client",
							Value: base64.StdEncoding.EncodeToString(clientBinary),
						},
						Append: wrapperspb.Bool(false),
					},
				)
			default:
				authHeaders = append(authHeaders,
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:   "x-as-user-id",
							Value: "",
						},
						Append: wrapperspb.Bool(false),
					},
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:   "x-as-user",
							Value: "",
						},
						Append: wrapperspb.Bool(false),
					},
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:   "x-as-client-id",
							Value: "",
						},
						Append: wrapperspb.Bool(false),
					},
					&v31.HeaderValueOption{
						Header: &v31.HeaderValue{
							Key:   "x-as-client",
							Value: "",
						},
						Append: wrapperspb.Bool(false),
					},
				)
			}
		}
	}

	return &authpb.CheckResponse{
		// Status: &rpcstatus.Status{
		// 	Code:    int32(rpc.UNAUTHENTICATED),
		// 	Message: "Authorization Header malformed or not provided",
		// },
		// HttpResponse: &authpb.CheckResponse_DeniedResponse{
		// 	DeniedResponse: &authpb.DeniedHttpResponse{
		// 		Status: &typev3.HttpStatus{
		// 			Code: typev3.StatusCode_Unauthorized,
		// 		},
		// 		Body: "Authorization Header malformed or not provided",
		// 	},
		// },
		HttpResponse: &authpb.CheckResponse_OkResponse{
			OkResponse: &authpb.OkHttpResponse{
				Headers: authHeaders,
				HeadersToRemove: []string{
					"authorization",
					"cookie",
				},
			},
		},
		DynamicMetadata: md,
	}, nil
}
