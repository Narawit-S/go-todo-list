package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error)  {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid keysize: must be exactly %d character", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto: paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (m *PasetoMaker) CreateToken(email string, duration time.Duration) (string, error)  {
	payload, err := NewPayload(email, duration)
	if err != nil{
		return "", err
	}

	return m.paseto.Encrypt(m.symmetricKey, payload, nil)
}

func (m *PasetoMaker) ValidateToken(token string) (*Payload, error)  {
	payload := &Payload{}

	err := m.paseto.Decrypt(token, m.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
