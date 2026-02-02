# Vault Unseal from 1Password

## Requirements

### Tools

- [1Password CLI](https://developer.1password.com/docs/cli/) (`op`)
- `kubectl` configured with access to the Vault cluster

### Vault

You must be connected to the cluster where Vault is installed.

Vault must be installed in namespace `vault` with 3 replicas (pod names are `vault-0`, `vault-1`, `vault-2` ) and initialized but sealed.

You have Vault init JSON file (`vault_init.local.json`) containing unseal keys, created by `vault operator init -format=json` command.

## Save Vault init JSON to 1Password

Save `vault_init.local.json` to 1Password:

```
slu vault save-vault-init-json-to-1password \
  --file vault_init.local.json \
  --vault-group my-example-client \
  --vault-name MY_EXAMPLE_CLIENT_VAULT_INIT_JSON
```

## Unseal Vault from 1Password

Unseal Vault pods (`vault-0`, `vault-1`, `vault-2`) using keys stored in 1Password:

```
slu vault unseal-from-1password \
  --vault-group my-example-client \
  --vault-name MY_EXAMPLE_CLIENT_VAULT_INIT_JSON
```
