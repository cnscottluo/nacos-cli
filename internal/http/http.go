package http

import (
	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
)

var HttpClient = resty.New()

func Get(url string) {

}

type LoginResult struct {
	AccessToken string `json:"accessToken"`
	TokenTtl    uint64 `json:"tokenTtl"`
	GlobalAdmin bool   `json:"globalAdmin"`
	Username    string `json:"username"`
}

func Login() {
	var loginResult LoginResult
	_, err := HttpClient.R().
		SetDebug(true).
		EnableGenerateCurlOnDebug().
		SetResult(&loginResult).
		SetFormData(map[string]string{
			"username": "nacos",
			"password": "nacos",
		}).
		Post("http://127.0.0.1:8848/nacos/v1/auth/login")
	if err == nil {
		println(loginResult.AccessToken)

		jwt.Parse(loginResult.AccessToken, func(token *jwt.Token) (interface{}, error) {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				println(claims["exp"].(float64))
			}
			return nil, nil
		})
	} else {
		println(err)
	}
}
