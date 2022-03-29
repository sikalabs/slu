package config

type SluSecrets struct {
	DigitalOcean []SluSecretsDigitalOcean
	SluVault     SluSecretsSluVault
}
