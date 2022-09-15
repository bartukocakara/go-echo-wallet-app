package service

import (
	"fmt"
	"os"
	"roof-stack/internal/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type IJWTService interface {
	GenerateToken(UserID entity.User) string
}


type JwtCustomClaim struct {
	UserID 				uint64	`json:"user_id"`
	Role 				string	`json:"role"`
	FirstName			string  `json:"first_name"`
	jwt.StandardClaims
}
type jwtService struct {
	secretKey 	string
	issuer 		string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer : getIssuer(),
	}
}

func (j *jwtService) GenerateToken(User entity.User) string {
	claims := &JwtCustomClaim{
		User.ID,
		"user",
		User.FirstName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	fmt.Println(getSecretKey())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(getSecretKey()))
	if err != nil {
		panic(err)
	}

	return t
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(getSecretKey()), nil
	})
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_ACCESS_SIGN_KEY")
	if secretKey == "" {
		secretKey = "greatest-secret-ever"
	}
	return secretKey
}

func getIssuer()string {
	issuer := os.Getenv("JWT_ISSUER")
	if issuer == "" {
		issuer = "my-amazing-company-name"
	}
	return issuer
}