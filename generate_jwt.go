package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// クライアント情報
	clientID := "06546c62-f625-40b8-b834-7f6a81b73dd0"
	clientSecret := "95a8612f4e08298d66a5cde88b20d00392d2ce49be425486b1a8e8c88dc918fad4e6ceb066cca0aff9eef4fb6326fefb18ef4211431e4b31734210502bc82163adc4ec41290a79372497ffcf75a5ba78b8c30c539ce8854dc5c264ea6c584578a6b687a2e3a729d6a9bb2c9830c490e6c7cfd0cfe00596ec3520e711aeaa9d3dae6faba54fa5789f465d50f73d4712bc0ddb58fa50327338a62f693fae516ac2221b39ccbb1d9c156cecc80f2759ae80f179341b0415bac8b37a7de51457dc58ad316f74717601e92f78efda444c293c06ce241c6dd6a27553a690e7091603151420903c29ce9d23797cf5ed5de9e3922d81f1b072bb491041eb5d4f1b342da5"

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
