package auth

import (
	"log"
	"src/internal/config"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(config.Config.AUTH.JWTSecretKey)

type Claims struct {
	UserID int
	jwt.StandardClaims
}

/* トークンを生成する */
func GenerateToken(userID int) string {
	// 有効期限の設定（15分）
	expirationTime := time.Now().Add(15 * time.Minute)

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// HS256のアルゴリズムで秘密鍵を利用し、JWTを署名する
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}

/* トークンを検証する */
// func ValidateToken(authHeaderValue string) (bool, int) {
// 	separatedHeaderValues := strings.Split(authHeaderValue, " ")
// 	if len(separatedHeaderValues) < 3 && separatedHeaderValues[0] != "Bearer" {
// 		return false, -1
// 	}

// 	receivedToken := separatedHeaderValues[1]
// 	claims := &Claims{}
// 	parsedToken, err := jwt.ParseWithClaims(receivedToken, claims, func(token *jwt.Token) (interface{}, error) {
// 		return config.Config.AUTH.JWTSecretKey, nil
// 	})
// 	if err != nil && err == jwt.ErrSignatureInvalid {
// 		return false, -1
// 	}

// 	if !parsedToken.Valid {
// 		return false, -1
// 	}
// 	return true, claims.UserID
// }
