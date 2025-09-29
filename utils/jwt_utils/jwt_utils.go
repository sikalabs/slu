package jwt_utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/coreos/go-oidc"
	"github.com/golang-jwt/jwt/v4"
)

func ParseJWT(jwtToken string, formatDates bool) {
	// Parse JWT
	token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) { return nil, nil })

	// Marshal header and claims to JSON
	headerJSON, _ := json.Marshal(token.Header)
	claimsJSON, _ := json.Marshal(token.Claims)

	// Prepare the result as a slice of interfaces
	result := []interface{}{
		decodeJSON(headerJSON, formatDates),
		decodeJSON(claimsJSON, formatDates),
		token.Signature,
	}

	// Print the result as a JSON array
	outputJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal("Error marshalling result to JSON: ", err)
	}
	fmt.Println(string(outputJSON))
}

func decodeJSON(data []byte, formatDates bool) interface{} {
	var obj map[string]interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		log.Fatal("Error unmarshalling JSON: ", err)
	}

	if formatDates {
		for key, value := range obj {
			if timestamp, ok := value.(float64); ok && isTimestamp(timestamp) {
				obj[key] = formatTimestamp(timestamp)
			}
		}
	}
	return obj
}

func ValidateJWT(issuer, rawToken string) error {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return err
	}

	_, err = provider.Verifier(&oidc.Config{SkipClientIDCheck: true}).Verify(ctx, rawToken)
	if err != nil {
		return err
	}

	return nil
}

// Check if the value is a valid timestamp
func isTimestamp(value float64) bool {
	// The timestamp should be greater than 0 and less than the current time + 100 years
	now := time.Now().Unix()
	return value > 0 && value < float64(now+100*365*24*60*60)
}

// Format timestampt to reader-friendly string
func formatTimestamp(timestamp float64) string {
	loc := time.Now().Location()
	return time.Unix(int64(timestamp), 0).In(loc).Format("2006-01-02 15:04:05")
}
