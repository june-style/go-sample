package services

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/derrors"
)

func NewJWTer(cfg *dconfig.Config) JWTer {
	return &jwterImpl{
		config:     cfg,
		expireTime: time.Second * time.Duration(cfg.App.SessionExpirationTime),
		issuer:     cfg.App.Key,
		hmacSecret: []byte(cfg.App.HMACSecret),
	}
}

//go:generate mockgen -source=${GOFILE} -destination=./mock/${GOFILE} -package=${GOPACKAGE}_mock
type JWTer interface {
	Create(ctx context.Context) (string, error)
	Verify(ctx context.Context, token string) error
}

var (
	ErrJWTTokenIsInvalid                 = derrors.NewUnauthenticated("JWT token is invalid")
	ErrJWTTokenIsUnexpectedSigningMethod = derrors.NewUnauthenticated("JWT token is unexpected signing method")
)

type jwterImpl struct {
	config     *dconfig.Config
	expireTime time.Duration
	issuer     string
	hmacSecret []byte
}

func (j *jwterImpl) Create(ctx context.Context) (string, error) {
	now, err := dcontext.GetTime(ctx)
	if err != nil {
		return "", derrors.Wrap(err)
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Audience:  jwt.ClaimStrings{dcontext.GetAuthenticatedUserID(ctx)},
		ExpiresAt: jwt.NewNumericDate(now.Add(j.expireTime)),
		IssuedAt:  jwt.NewNumericDate(now),
		Issuer:    j.issuer,
	})
	token, err := jwtToken.SignedString(j.hmacSecret)
	if err != nil {
		return "", derrors.Wrap(err)
	}

	return token, nil
}

func (j *jwterImpl) Verify(ctx context.Context, token string) error {
	now, err := dcontext.GetTime(ctx)
	if err != nil {
		return derrors.Wrap(err)
	}

	jwtToken, err := jwt.Parse(token, func(jwtToken *jwt.Token) (any, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, derrors.Wrapf(ErrJWTTokenIsUnexpectedSigningMethod, "signing-method is %s", jwtToken.Header["alg"])
		}
		return j.hmacSecret, nil
	},
		jwt.WithTimeFunc(func() time.Time { return now }),
		jwt.WithAudience(dcontext.GetAuthenticatedUserID(ctx)),
		jwt.WithExpirationRequired(),
		jwt.WithIssuedAt(),
		jwt.WithIssuer(j.issuer),
	)
	if err != nil {
		return derrors.Wrap(err)
	}
	if !jwtToken.Valid {
		return derrors.Wrap(ErrJWTTokenIsInvalid)
	}

	return nil
}
