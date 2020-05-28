package groups

import (
	"gitlab.com/bloom42/gobox/crypto"
	"gitlab.com/bloom42/gobox/uuid"
)

func SignInvitation(inviterPrivateKey crypto.PrivateKey, groupID uuid.UUID, inviteeUsername string,
	inviteePublicKey crypto.PublicKey, ephemeralPubliKey crypto.PublicKey) ([]byte, error) {

	message := make([]byte, 0)
	message = append(message, groupID.Bytes()...)
	message = append(message, []byte(inviteeUsername)...)
	message = append(message, []byte(inviteePublicKey)...)
	message = append(message, []byte(ephemeralPubliKey)...)

	signature, err := inviterPrivateKey.Sign(crypto.RandReader(), message, crypto.PrivateKeySignerOpts)
	return signature, err
}
