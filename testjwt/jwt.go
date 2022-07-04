package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func main() {
	sign := []byte("ahut_test")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": "司大帅",
		"exp":  time.Now().Unix() + 5,
		"iss":  "sywdebug",
	})
	tokenString, err := token.SignedString(sign)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("加密后的token字符串", tokenString)
	token, _ = jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return sign, nil
	})
	time.Sleep(6 * time.Second)
	fmt.Println(token.Claims)

}
