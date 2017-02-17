package main

import (
	"fmt"

	"time"

	"github.com/robbert229/jwt"
)

func main() {
	fmt.Printf("Hello world\n")

	secret := "ThisIsMySuperDuperSecret"

	algor := jwt.HmacSha256(secret)

	claims := jwt.NewClaim()
	claims.Set("Role", "Admin")
	claims.SetTime("exp", time.Now().Add(time.Minute))

	token, err := algor.Encode(claims)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Token: %s\n", token)

}
