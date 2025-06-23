package keycloak_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func PasswordResetOrDie(keycloakUrl, adminUser, adminPass, realm, username, password string) {
	err := PasswordReset(keycloakUrl, adminUser, adminPass, realm, username, password)
	if err != nil {
		log.Fatalln("Password reset failed: " + err.Error())
	}
	fmt.Println("Password reset successfully!")
}

func PasswordReset(keycloakUrl, adminUser, adminPass, realm, username, password string) error {
	token, err := getToken(keycloakUrl, adminUser, adminPass, "admin-cli")
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	userID, err := getUserID(keycloakUrl, realm, token, username)
	if err != nil {
		return fmt.Errorf("failed to get user ID: %w", err)
	}

	err = resetPassword(keycloakUrl, realm, token, userID, password, true)
	if err != nil {
		return fmt.Errorf("failed to reset password: %w", err)
	}

	return nil
}

func getToken(keycloakUrl, user, pass, clientID string) (string, error) {
	data := []byte(fmt.Sprintf("username=%s&password=%s&grant_type=password&client_id=%s", user, pass, clientID))

	resp, err := http.Post(
		fmt.Sprintf("%s/realms/master/protocol/openid-connect/token", keycloakUrl),
		"application/x-www-form-urlencoded",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get access token: %v", result)
	}
	return token, nil
}

func getUserID(keycloakUrl, realm, token, username string) (string, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/admin/realms/%s/users?username=%s", keycloakUrl, realm, username), nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var users []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&users)

	if len(users) == 0 {
		return "", fmt.Errorf("user not found")
	}

	return users[0]["id"].(string), nil
}

func resetPassword(keycloakUrl, realm, token, userID, password string, temporary bool) error {
	passwordPayload := map[string]interface{}{
		"type":      "password",
		"value":     password,
		"temporary": temporary,
	}
	payloadBytes, _ := json.Marshal(passwordPayload)

	req, _ := http.NewRequest("PUT",
		fmt.Sprintf("%s/admin/realms/%s/users/%s/reset-password", keycloakUrl, realm, userID),
		bytes.NewBuffer(payloadBytes),
	)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to reset password: %s", body)
	}
	return nil
}
