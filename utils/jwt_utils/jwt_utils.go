package jwt_utils

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v4"
)

func ParseJWT(jwtToken string) {
	// Parse JWT
	token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) { return nil, nil })

	// Marshal header and claims to JSON
	headerJSON, _ := json.Marshal(token.Header)
	claimsJSON, _ := json.Marshal(token.Claims)

	// Prepare the result as a slice of interfaces
	result := []interface{}{
		decodeJSON(headerJSON),
		decodeJSON(claimsJSON),
		token.Signature,
	}

	// Print the result as a JSON array
	outputJSON, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error marshalling result to JSON: ", err)
	}
	fmt.Println(string(outputJSON))
}

func decodeJSON(data []byte) interface{} {
	var obj interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		log.Fatal("Error unmarshalling JSON: ", err)
	}
	return obj
}
