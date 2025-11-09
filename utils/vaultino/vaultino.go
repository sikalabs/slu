package vaultino

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/tidwall/gjson"
	"golang.org/x/crypto/argon2"
	"golang.org/x/term"
	"gopkg.in/yaml.v3"
)

const (
	saltSize  = 16
	nonceSize = 12
)

type vaultData struct {
	header     string
	salt       []byte
	nonce      []byte
	ciphertext []byte
}

func CreateVault(name string, file string) error {
	plaintext, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("error reading source file: %w", err)
	}

	password, err := readPassword()
	if err != nil {
		return fmt.Errorf("error reading password: %w", err)
	}

	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return fmt.Errorf("error generating salt: %w", err)
	}

	nonce := make([]byte, nonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return fmt.Errorf("error generating nonce: %w", err)
	}

	ciphertext, err := encrypt(plaintext, password, salt, nonce)
	if err != nil {
		return err
	}

	header := buildHeader(file)
	vaultFile := fmt.Sprintf("%s.vault", name)

	return writeVaultFile(vaultFile, header, salt, nonce, ciphertext)
}

func DecryptVaultToFile(vaultFile string) error {
	vd, err := readVaultFile(vaultFile)
	if err != nil {
		return err
	}

	password, err := readPassword()
	if err != nil {
		return fmt.Errorf("error reading password: %w", err)
	}

	plaintext, err := decrypt(vd.ciphertext, password, vd.salt, vd.nonce)
	if err != nil {
		return err
	}

	fileType := parseHeaderValue(vd.header, "type")
	name := parseHeaderValue(vd.header, "name")
	outFile := fmt.Sprintf("%s.%s", name, fileType)

	return os.WriteFile(outFile, plaintext, 0600)
}

func EditVault(vaultFile string) error {
	return editVaultInternal(vaultFile, false)
}

func EditVaultWithPasswordChange(vaultFile string) error {
	return editVaultInternal(vaultFile, true)
}

func editVaultInternal(vaultFile string, changePassword bool) error {
	vd, err := readVaultFile(vaultFile)
	if err != nil {
		return fmt.Errorf("error reading vault: %w", err)
	}

	password, err := readPassword()
	if err != nil {
		return fmt.Errorf("error reading password: %w", err)
	}

	plaintext, err := decrypt(vd.ciphertext, password, vd.salt, vd.nonce)
	if err != nil {
		return err
	}

	tmpFile, err := ioutil.TempFile("", "myvault-edit-*.tmp")
	if err != nil {
		return fmt.Errorf("error creating temporary file: %w", err)
	}
	tmpFileName := tmpFile.Name()

	if _, err := tmpFile.Write(plaintext); err != nil {
		tmpFile.Close()
		os.Remove(tmpFileName)
		return fmt.Errorf("error writing to temp file: %w", err)
	}
	tmpFile.Close()

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	cmd := exec.Command(editor, tmpFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		os.Remove(tmpFileName)
		return fmt.Errorf("error editing vault: %w", err)
	}

	// Read edited content
	editedContent, err := os.ReadFile(tmpFileName)
	if err != nil {
		os.Remove(tmpFileName)
		return fmt.Errorf("error reading edited file: %w", err)
	}
	os.Remove(tmpFileName)

	// Preserve original name and type, update user and time
	name := parseHeaderValue(vd.header, "name")
	fileType := parseHeaderValue(vd.header, "type")

	// Create new header with original name and type
	user := os.Getenv("USER")
	if user == "" {
		user = os.Getenv("USERNAME")
	}
	timestamp := time.Now().Format(time.RFC3339)
	header := fmt.Sprintf("$VAULTINO;1.2;AES256-GCM;argon2id;name=%s;type=%s;user=%s;time=%s",
		name, fileType, user, timestamp)

	// Determine which password to use for re-encryption
	encryptPassword := password
	if changePassword {
		fmt.Println("Enter new password for the vault:")
		newPassword, err := readPassword()
		if err != nil {
			return fmt.Errorf("error reading new password: %w", err)
		}
		encryptPassword = newPassword
	}

	// Re-encrypt with new salt and nonce
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return fmt.Errorf("error generating salt: %w", err)
	}

	nonce := make([]byte, nonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return fmt.Errorf("error generating nonce: %w", err)
	}

	ciphertext, err := encrypt(editedContent, encryptPassword, salt, nonce)
	if err != nil {
		return err
	}

	return writeVaultFile(vaultFile, header, salt, nonce, ciphertext)
}

func GetSecretFromVault(vaultFile, key, password string) (string, error) {
	data, err := decryptVaultInMemory(vaultFile, password)
	if err != nil {
		return "", err
	}

	fileType := parseHeaderValueFromFile(vaultFile, "type")
	return extractValue(data, key, fileType)
}

func encrypt(plaintext, password, salt, nonce []byte) ([]byte, error) {
	key := deriveKey(password, salt)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error creating cipher: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("error creating GCM: %w", err)
	}

	return aesGCM.Seal(nil, nonce, plaintext, nil), nil
}

func decrypt(ciphertext, password, salt, nonce []byte) ([]byte, error) {
	key := deriveKey(password, salt)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error creating cipher: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("error creating GCM: %w", err)
	}

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("decryption failed: incorrect password or corrupted data")
	}

	return plaintext, nil
}

func deriveKey(password, salt []byte) []byte {
	return argon2.IDKey(password, salt, 3, 64*1024, 4, 32)
}

func readVaultFile(vaultFile string) (*vaultData, error) {
	content, err := os.ReadFile(vaultFile)
	if err != nil {
		return nil, fmt.Errorf("error reading vault file: %w", err)
	}

	parts := strings.SplitN(string(content), "\n", 2)
	if len(parts) < 2 {
		return nil, errors.New("invalid vault format")
	}

	data, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 content: %w", err)
	}

	if len(data) < saltSize+nonceSize {
		return nil, errors.New("invalid encrypted data")
	}

	return &vaultData{
		header:     parts[0],
		salt:       data[:saltSize],
		nonce:      data[saltSize : saltSize+nonceSize],
		ciphertext: data[saltSize+nonceSize:],
	}, nil
}

func writeVaultFile(filename, header string, salt, nonce, ciphertext []byte) error {
	data := append(salt, nonce...)
	data = append(data, ciphertext...)
	encoded := base64.StdEncoding.EncodeToString(data)

	content := fmt.Sprintf("%s\n%s", header, encoded)
	return os.WriteFile(filename, []byte(content), 0600)
}

func decryptVaultInMemory(vaultFile string, providedPassword string) ([]byte, error) {
	vd, err := readVaultFile(vaultFile)
	if err != nil {
		return nil, err
	}

	password, err := getPassword(providedPassword)
	if err != nil {
		return nil, fmt.Errorf("error reading password: %w", err)
	}

	return decrypt(vd.ciphertext, password, vd.salt, vd.nonce)
}

func buildHeader(file string) string {
	user := os.Getenv("USER")
	if user == "" {
		user = os.Getenv("USERNAME")
	}
	timestamp := time.Now().Format(time.RFC3339)
	ext := strings.TrimPrefix(filepath.Ext(file), ".")
	filename := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))

	return fmt.Sprintf("$VAULTINO;1.2;AES256-GCM;argon2id;name=%s;type=%s;user=%s;time=%s",
		filename, ext, user, timestamp)
}

func extractValue(data []byte, key, fileType string) (string, error) {
	switch fileType {
	case "yaml", "yml":
		return getFromYAML(data, key)
	case "json":
		return getFromJSON(data, key)
	case "env":
		return getFromENV(data, key)
	default:
		return "", fmt.Errorf("unsupported file type: %s", fileType)
	}
}

func getFromYAML(data []byte, key string) (string, error) {
	// Convert YAML to JSON first, then use gjson for consistent parsing
	var yamlData interface{}
	if err := yaml.Unmarshal(data, &yamlData); err != nil {
		return "", fmt.Errorf("error parsing YAML: %w", err)
	}

	jsonData, err := json.Marshal(yamlData)
	if err != nil {
		return "", fmt.Errorf("error converting YAML to JSON: %w", err)
	}

	return getValueUsingGjson(string(jsonData), key)
}

func getFromJSON(data []byte, key string) (string, error) {
	return getValueUsingGjson(string(data), key)
}

func getValueUsingGjson(jsonStr string, key string) (string, error) {
	result := gjson.Get(jsonStr, key)
	if !result.Exists() {
		return "", fmt.Errorf("key %s not found", key)
	}
	return result.String(), nil
}

func getFromENV(data []byte, key string) (string, error) {
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, key+"=") {
			return strings.TrimPrefix(line, key+"="), nil
		}
	}
	return "", fmt.Errorf("key %s not found", key)
}

func parseHeaderValue(header, key string) string {
	parts := strings.Split(header, ";")
	for _, p := range parts {
		if strings.HasPrefix(p, key+"=") {
			return strings.TrimPrefix(p, key+"=")
		}
	}
	return ""
}

func parseHeaderValueFromFile(vaultFile, key string) string {
	content, err := os.ReadFile(vaultFile)
	if err != nil {
		return ""
	}

	parts := strings.SplitN(string(content), "\n", 2)
	if len(parts) < 1 {
		return ""
	}

	return parseHeaderValue(parts[0], key)
}

func getPassword(providedPassword string) ([]byte, error) {
	if providedPassword != "" {
		return []byte(providedPassword), nil
	}
	return readPassword()
}

func readPassword() ([]byte, error) {
	fmt.Print("Enter password: ")
	pass, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return pass, err
}
