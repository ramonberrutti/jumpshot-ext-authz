package storage

import (
	"context"
	"fmt"

	"github.com/ramonberrutti/jumpshot-ext-authz/internal/storage/engine"
	"google.golang.org/protobuf/proto"

	storagepb "github.com/ramonberrutti/jumpshot/protogen/storage"
)

type Sessions interface {
	Get(ctx context.Context, id string) (*storagepb.Session, error)
}

type sessions struct {
	e engine.Engine
}

func (s *sessions) getKey(id string) string {
	return fmt.Sprintf("sessions/%s", id)
}

func (s *sessions) Get(ctx context.Context, id string) (*storagepb.Session, error) {
	key := s.getKey(id)
	b, err := s.e.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var session storagepb.Session
	if err := proto.Unmarshal(b, &session); err != nil {
		return nil, err
	}

	return &session, nil
}
