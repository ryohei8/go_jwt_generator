package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// client_id
	clientID := ""
	// client_secret
	clientSecret := ""

	// JWTペイロード（クレーム）
	claims := jwt.MapClaims{
		"client_id": clientID,
		"iat":       time.Now().Unix(),                // 発行時刻
		"exp":       time.Now().Add(time.Hour).Unix(), // 有効期限（1時間後）
	}

	// トークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名
	signedToken, err := token.SignedString([]byte(clientSecret))
	if err != nil {
		log.Fatalf("Failed to sign token: %v", err)
	}

	fmt.Printf("Generated JWT Token:\n%s\n", signedToken)
}
