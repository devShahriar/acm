package helper

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"magic.pathao.com/carta/carta-acm/internal/config"
	"magic.pathao.com/carta/carta-acm/internal/contract"
)

func generateRandomString(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(randomBytes), nil
}

func GenerateUniqueAPIKey(userId, email, serviceType string) (string, error) {
	randomString, err := generateRandomString(32)
	if err != nil {
		return "", err
	}

	combinedData := userId + email + serviceType + randomString
	hash := sha256.Sum256([]byte(combinedData))
	apiKey := fmt.Sprintf("%x", hash)
	return apiKey, nil
}

func GenerateAccessToken(payload contract.JwtPayload) (string, error) {

	secretKey := config.GetAppConfig().JWTSecretKey
	fmt.Println("secretKey", secretKey)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          payload.Id,
		"email":       payload.Email,
		"role_id":     payload.RoleId,
		"org_id":      payload.OrgId,
		"member_id":   payload.MemberId,
		"permissions": payload.Permissions,
		"created_at":  time.Now(),
	})

	tokenString, err := claims.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}

// func ParseJwt(token string) {

// 	secretKey := config.GetAppConfig().JWTSecretKey

// }

func ParseJWTToken(tokenString string, secretKey []byte) (*contract.JwtPayload, error) {

	claims := &contract.JwtPayload{}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Check for the expected signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key
		return secretKey, nil
	})

	// Check if there was an error in parsing the token
	if claims, ok := token.Claims.(*contract.JwtPayload); ok && token.Valid {
		fmt.Printf("user_id: %v %v\n", claims.Id, claims.RegisteredClaims.Issuer)
		return claims, nil
	}

	return nil, err
}

func GenerateUUID() string {
	return uuid.New().String()
}
