package auth

import (
	"fmt"
	"log"
	"net/http"
	"src/internal/config"
	"src/internal/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var jwtKey = []byte(config.Config.AUTH.JWTSecretKey)

type JWTCustomClaims struct {
	UserID int    `json:"userID"`
	Name   string `json:"name"`
	jwt.StandardClaims
}

/* アクセストークンの生成 */
func GenerateAccessToken(user *models.User) (string, time.Time, error) {
	// 有効期限の設定（15分）
	expirationTime := time.Now().Add(1 * time.Minute)

	return generateToken(user, expirationTime, []byte(config.Config.AUTH.JWTSecretKey))
}

/* JWTトークン生成の主要なロジック */
func generateToken(user *models.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {

	// ユーザ名と有効期限を含むJWTクレームを作成する
	claims := &JWTCustomClaims{
		UserID: user.ID,
		Name:   user.Name,
		StandardClaims: jwt.StandardClaims{
			// 有効期限はUnixミリ秒で表される
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// 署名に使用したHS256アルゴリズムと、クレームでトークンを宣言する
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// JWT文字列の作成
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Fatal(err)
	}
	return tokenString, expirationTime, nil
}

/* リフレッシュトークの生成 */
func GenerateRefreshToken(user *models.User) (string, time.Time, error) {
	// 有効期限の設定(1時間)
	expirationTime := time.Now().Add(1 * time.Hour)

	return generateToken(user, expirationTime, []byte(config.Config.AUTH.JWTRefreshSecretKey))
}

/* 有効なJWTトークンを保存する新しいクッキーを作成する */
func SetTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"

	// スクリプトから保護する
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

/* ユーザ名を保存するクッキー */
func SetUserCookie(user *models.User, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = user.Name
	cookie.Expires = expiration
	cookie.Path = "/"
	c.SetCookie(cookie)
}

/* クッキーをチェックする */
func CheckCookieMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenCookie, _ := c.Cookie(config.Config.AUTH.AccessTokenCookieName)
		accessToken := tokenCookie.Value

		token, err := jwt.ParseWithClaims(accessToken, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.AUTH.JWTSecretKey), nil
		})
		if err != nil {
			fmt.Println(err)
		}

		claims := token.Claims.(*JWTCustomClaims)
		// log.Infof("claims.Name: %+v", claims.Name)

		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": claims.Name,
			})
		}

		return next(c)
	}
}

func WriteCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

func ReadCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "read a cookie")
}

/* ユーザが制限されたパスにアクセスしようとしたときに実行する */
// func JWTErrorChecker(err error, c echo.Context) error {
// 	// signinページにリダイレクトさせる
// 	fmt.Println("----------エラーチェッカー----------------")
// 	// return c.Redirect(http.StatusFound, "/sign_in")
// 	return c.Redirect(http.StatusSeeOther, c.Echo().Reverse("/sign_in"))
// }

// ミドルウェア制限検証用
// func Restricted(c echo.Context) error {
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(*JWTCustomClaims)
// 	userID := claims.UserID
// 	name := claims.Name
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"userID":   userID,
// 		"userName": name,
// 	})
// }
