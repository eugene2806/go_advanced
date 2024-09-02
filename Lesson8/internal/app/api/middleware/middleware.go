package middleware

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

var (
	SecretKey      []byte      = []byte("UltraRestApiSecurityKey9000")
	emptyValidFunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	}
)

var JwtMiddleWare = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: emptyValidFunc,
	SigningMethod:       jwt.SigningMethodHS256,
})
