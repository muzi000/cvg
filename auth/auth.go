package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/muzi000/vuln-go/module"
)

var jwtSecret []byte = []byte("1247sksjanbsajha.,&sb")

func Auth(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	if path == "/login" {
		ctx.Next()
		return
	}
	userJwt, err := ctx.Cookie("jwt")
	if err != nil {
		ctx.Abort()
		ctx.HTML(http.StatusUnauthorized, "auth-redirect.html", gin.H{
			"redictUrl": "/login?return=" + path,
		})
		return
	}
	u, err := ParseToken(userJwt)
	if err != nil {
		ctx.Abort()
		ctx.SetCookie("jwt", "", -1, "/", "localhost", false, true)
		ctx.HTML(http.StatusUnauthorized, "auth-redirect.html", gin.H{
			"redictUrl": "/login?return=" + path,
		})
		return
	}
	ctx.Set("user", u)
	ctx.Next()

}

type Claims struct {
	User module.User `json:"user"`
	jwt.RegisteredClaims
}

func GenerateToken(name string, role int) (string, error) {
	claims := Claims{
		module.User{
			Id:   0,
			Name: name,
			Role: role,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSecret)
	return ss, err
}

func ParseToken(tokenstring string) (module.User, error) {
	token, err := jwt.ParseWithClaims(tokenstring, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if c, ok := token.Claims.(*Claims); ok && token.Valid {
		return c.User, nil
	} else {
		return module.User{}, err
	}

}
