package token

import (
	"errors"
	"gaoMall/app"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    uint      `json:"userId"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return errors.New("token has expired")
	}
	return nil
}

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token has invalid")
)

func CreateToken(userId uint) (string, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	payload := &Payload{
		ID:        tokenId,
		UserID:    userId,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(app.Config.Token.Duration),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(app.Config.Token.SymmetricKey))
}

func VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(app.Config.Token.SymmetricKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		var vErr *jwt.ValidationError
		ok := errors.As(err, &vErr)
		if ok && errors.Is(vErr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
