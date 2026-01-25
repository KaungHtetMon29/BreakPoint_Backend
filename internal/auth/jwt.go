package auth

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type Test_Credential struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateJWTToken(userInfo *UserInfo) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, Test_Credential{
		Name:  userInfo.Name,
		Email: userInfo.Id,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN_SECRET")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func ParseJWTToken(tokenString string) (*Test_Credential, error) {
	t, err := jwt.ParseWithClaims(tokenString, &Test_Credential{}, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := t.Claims.(*Test_Credential); ok {
		return claims, err
	} else {
		return nil, fmt.Errorf("unknown claims type, cannot proceed")
	}
}
