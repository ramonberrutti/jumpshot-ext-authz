package storage

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"golang.org/x/crypto/bcrypt"
)

func (x *User) ValidatePassword(password string) bool {
	if x == nil {
		return false
	}

	hash := x.GetPassword().Value
	if len(hash) == 0 {
		return false
	}

	return bcrypt.CompareHashAndPassword(hash, []byte(password)) == nil
}

func (x *User) WebAuthnID() []byte {
	if x == nil {
		return nil
	}

	return []byte(x.Id)
}

func (x *User) WebAuthnName() string {
	if x == nil {
		return ""
	}

	return x.Name
}

func (x *User) WebAuthnDisplayName() string {
	if x == nil {
		return ""
	}

	return x.Name
}

func (x *User) WebAuthnCredentials() []webauthn.Credential {
	if x == nil {
		return nil
	}

	if x.Password == nil {
		return nil
	}

	credentials := x.Password.GetWebauthnCredentials()
	retCredentials := make([]webauthn.Credential, 0, len(credentials))

	for _, credential := range credentials {
		transports := make([]protocol.AuthenticatorTransport, 0, len(credential.Transports))
		for _, transport := range credential.Transports {
			transports = append(transports, protocol.AuthenticatorTransport(transport))
		}

		retCredentials = append(retCredentials, webauthn.Credential{
			ID:              credential.Id,
			PublicKey:       credential.PublicKey,
			AttestationType: credential.AttestationType,
			Transport:       transports,
			Authenticator: webauthn.Authenticator{
				AAGUID:       credential.Authenticator.Aaguid,
				SignCount:    credential.Authenticator.SignCount,
				CloneWarning: credential.Authenticator.CloneWarning,
				Attachment:   protocol.AuthenticatorAttachment(credential.Authenticator.Attachment),
			},
			Flags: webauthn.CredentialFlags{
				UserPresent:    credential.Flags.UserPresent,
				UserVerified:   credential.Flags.UserVerified,
				BackupEligible: credential.Flags.BackupEligible,
				BackupState:    credential.Flags.BackupState,
			},
		})
	}

	return retCredentials
}

func (x *User) WebAuthnIcon() string {
	return ""
}
