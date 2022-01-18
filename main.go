package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/golang-jwt/jwt"
)

func main() {
	content, _ := ioutil.ReadFile("keys/private.pem")
	privKey := content
	// signer, _ := jwt.ParseRSAPrivateKeyFromPEMWithPassword(key, "1234")
	signer, _ := jwt.ParseRSAPrivateKeyFromPEM(privKey)

	content, _ = ioutil.ReadFile("keys/public.pem")
	publicKey := content
	// signer, _ := jwt.ParseRSAPrivateKeyFromPEMWithPassword(key, "1234")
	signed, _ := jwt.ParseRSAPublicKeyFromPEM(publicKey)

	now := time.Now().UTC()
	ttl, _ := time.ParseDuration("1h")

	claims := make(jwt.MapClaims)
	claims["dat"] = "text"
	claims["exp"] = now.Add(ttl).Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix()          // The time at which the token was issued.
	claims["nbf"] = now.Unix()          // The time before which the token must be disregarded.

	token, _ := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(signer)

	token = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXQiOiJ0ZXh0IiwiZXhwIjoxNjQyNDQ4MDIyLCJpYXQiOjE2NDI0NDQ0MjIsIm5iZiI6MTY0MjQ0NDQyMn0.VrYtQ2bVPzs5Bj_DlX54yzfCkAJ6XaPRLmDLeJHpaG3ofhGH9nX5J3kVzO_QSlMy8s-umHKW9A1lFhccN7uvKjpWkCkUBfoqNlxZwxZY-bAfeqYjBhaQ138-1J064FFX0WUxWKIdwEKr78XFI_BB1Ta90Vazbstf7gNUFNVHz8XRfyLjWtmNj5Y8vtkQ0EP1B0VVGqgElR5IZuIqZUzr445R-U0nCuNruQPv1nM5fpLo_5a7RPn4SYG8AqvkPXdkSetvrsjYh38xuAKyMKWnT153Ou5ud3VyZdYsk0cFS13rHmxnl9tJxwp5bmdGFFqb_AnPWo8eLpNWdb7huU14Gw`

	parsedToken, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}
		return signed, nil
	})

	if err != nil {
		panic(err)
	}

	parsedClaims, _ := parsedToken.Claims.(jwt.MapClaims)

	// if !ok || !parsedToken.Valid {
	// 	panic("Cries")
	// }

	fmt.Printf("%+v", parsedClaims)

}
