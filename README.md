# Azure Importer for Azure keyVault 🔑
## Description
This simple project aims to create secrets (in batches) into a valid `Azure keyVault` destination.

## Prerequisites
* Valid **Azure CLI** installation. See this [link](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli) for more detailed instructions
* Go
---

## Configuration 
### Files
There are two main configuration files:
* `import.config`: contains necessary configuration to reach Azure. It means, it's required to build the azure client, and operates against azure KeyVault.
```txt
SUBSCRIPTION_ID=44b2w4ebb-8442-0000-4w56-9k39146bd
TENANT_ID=555555-7dd4-8v3a-b87d-7777777777
RESOURCE_GROUP_NAME=rg-name-in-azure
SOURCE_SECRETS_PATH="importer/secrets.json"
LOCATION="West Europe"
 ```

All these fields are **mandatory** — ensure they are added properly and corresponds to valid Azure values.

* `importer/secrets.json`: This file holds the `keyvault` name and the `secrets` (names and values) that will be created in AzurekeyVault. Please, use the `secrets_template.json` template as a starting point.

```json
{
  "keyVaultName": "key-vault-name-that-exists",
  "secrets": [
    {
      "secretName": "mySecret",
      "secretValue": "ExampleOfASuperSecretValue",
      "isOverrideAllowed": false
    },
    {
      "secretName": "mySecret2",
      "secretValue": "ExampleOfASuperSecretValueWhichWillOverrideWhatExistsAlreadyInAzureKeyVault",
      "isOverrideAllowed": false
    }
  ]
}
```

This configuration aims to be as explicit as possible, which means it'll fail whether there's an invalid configuration or there are some secret fields detected as empty.
**Note** : The `IsOverrideAllowed` isn't working yet. In Azure KeyVault, a deleted secret is actually soft-deleted. A purge operation need to take place to definitively erase it.

### Azure Configuration
* Ensure that a valid `az login` has been performed without errors. 
* Ensure that you're pointing to the correct subscription before execute this application. E.g.: `az account set --subscription <my-subscription-id>`
* In the very first execution, it'll create a `service principal` with minimal permissions to operate against the target azure KeyVault.
* If there's a `sp-creds.json` file detected, it'll be loaded instead of creating a new `service principal`. This file will be placed in the `importer/` folder.
sp-creds.json file example:
```json
{
 "appId": "....",
 "displayName": "....",
 "name": "....",
 "password": "....",
 "tenant": "...."
}

```

---
## Run
Build the program first, running:
```bash
make build
```
After the above step, you're able to execute the binary, ensuring that the configuration files are filled with proper values, as it was indicated above
```bash
make run

```