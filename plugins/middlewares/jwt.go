package middlewares

import "github.com/dgrijalva/jwt-go"

// JWT ...
type JWT struct {
	secret string
}

// Claims ...
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// NewJWT ...
func NewJWT() *JWT {
	return &JWT{}
}

// Sign ...
func (j *JWT) Sign() (string, error) {
	return "", nil
}

// UnSign ...
func (j *JWT) UnSign() {

}
