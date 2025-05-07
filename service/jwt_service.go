package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adieos/imk-backend/constants"
	"github.com/adieos/imk-backend/dto"
	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateTokenEmail(email string) string
	GenerateTokenUser(userId string, role string) string
	ValidateTokenEmail(token string) (*jwt.Token, error)
	ValidateTokenUser(token string) (*jwt.Token, error)
	GetUserIDByToken(token string) (string, string, error)
	GetUserEmailByToken(token string) (string, string, error)
}

type jwtEmailClaim struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type jwtUserClaim struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "INI LHO ITS 2025",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "Template"
	}
	return secretKey
}

func (j *jwtService) GenerateTokenEmail(email string) string {
	claims := jwtEmailClaim{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * constants.JWT_EXPIRE_TIME_IN_MINS)),
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tx, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Println(err)
	}
	return tx
}

func (j *jwtService) GenerateTokenUser(userId string, role string) string {
	claims := jwtUserClaim{
		userId,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * constants.JWT_EXPIRE_TIME_IN_MINS)),
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tx, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Println(err)
	}
	return tx
}

func (j *jwtService) parseToken(t_ *jwt.Token) (any, error) {
	if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
	}
	return []byte(j.secretKey), nil
}

func (j *jwtService) ValidateTokenEmail(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &jwtEmailClaim{}, j.parseToken)
}

func (j *jwtService) ValidateTokenUser(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &jwtUserClaim{}, j.parseToken)
}

func (j *jwtService) GetUserIDByToken(token string) (string, string, error) {
	t_Token, err := j.ValidateTokenUser(token)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return "", "", dto.ErrTokenExpired
		}
		return "", "", dto.ErrTokenInvalid
	}
	if !t_Token.Valid {
		return "", "", dto.ErrTokenInvalid
	}

	claims := t_Token.Claims.(*jwtUserClaim)

	id := fmt.Sprintf("%v", claims.UserID)
	role := fmt.Sprintf("%v", claims.Role)
	return id, role, nil
}

func (j *jwtService) GetUserEmailByToken(token string) (string, string, error) {
	t_Token, err := j.ValidateTokenEmail(token)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return "", "", dto.ErrTokenExpired
		}
		return "", "", dto.ErrTokenInvalid
	}

	claims, ok := t_Token.Claims.(*jwtEmailClaim)
	if !ok || !t_Token.Valid {
		return "", "", fmt.Errorf("invalid token claims")
	}

	email := claims.Email
	expiration := claims.ExpiresAt.Time.Format("2006-01-02 15:04:05")

	return email, expiration, nil
}
