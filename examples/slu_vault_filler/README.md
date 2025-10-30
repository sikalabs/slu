# slu vault filler example

Run Vault

```
vault server -dev --dev-root-token-id=root
```

Login to Vault

```
vault login -address=http://127.0.0.1:8200 root
```

Run the slu vault filler

```
slu vault filler
```

Check the secrets in Vault

- http://127.0.0.1:8200/ui/vault/secrets/secret/kv/slu_vault_filler_example%2Fapp1/details?version=1
- http://127.0.0.1:8200/ui/vault/secrets/secret/kv/slu_vault_filler_example%2Fapp2/details?version=1
