package token

import (
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	"log"
	"time"
)

// PsetoMaker is a PASETO token maker
type PasetoMaker struct {
	paseto      *paseto.V2
	symetricKey []byte
}

func NewPasetoMaker(symetricKey string) (Maker, error) {
	if len(symetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("Invalid key size:  must me exactly %s characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:      paseto.NewV2(),
		symetricKey: []byte(symetricKey),
	}

	return maker, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", nil
	}
	return maker.paseto.Encrypt(maker.symetricKey, payload, nil)
}

// VerifyToken checks if the token is valid or not
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := maker.paseto.Decrypt(token, maker.symetricKey, &payload, nil)
	if err != nil {
		log.Printf("Failed to decrypt the token: %v", err)
		return nil, ErrInvalidaToken
	}
	log.Printf("Token decrypted successfully. Payload: %v", payload)

	err = payload.Valid()
	if err != nil {
		log.Printf("Payload is invalid: %v", err)
		return nil, err
	}
	log.Printf("Payload validated. Payload: %v", payload)

	return payload, nil
}
