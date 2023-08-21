package handlers

import "github.com/go-chi/jwtauth/v5"

var tokenAuth *jwtauth.JWTAuth

const Secret = "f6fe4j33jf2fdqDv"

func init() {
	tokenAuth = jwtauth.New("HS256", []byte(Secret), nil)
}

func MakeToken(name string) string {
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"username": name})
	return tokenString
}
