package main

import (
	"github.com/Excoriate/azure-keyvault-importer-go/pkg/azure"
	"github.com/Excoriate/azure-keyvault-importer-go/pkg/keyvault"
	"github.com/Excoriate/azure-keyvault-importer-go/pkg/lib/config"
	"github.com/Excoriate/azure-keyvault-importer-go/pkg/lib/secrets"
)

var (
	logger = config.GetLogger()
)

func main() {
	logger.Info("Starting importer")
	logger.Info("Importing settings file expected in root path: ./import.config")

	//Load importer configuration
	importerConfig := config.LoadSettings()

	// Load required json file with secrets to import
	secretsFile, _ := secrets.LoadSecretsJSONFile(importerConfig.SecretsSourcePath)
	logger.Info(secretsFile)

	// Validate secret configuration file
	if !secrets.IsSecretFileValid(secretsFile) {
		logger.Fatal("Secret file with errors. Fix errors and try again.")
	}

	// Set credentials to login against azure
	azure.SetCredentialsInProcess(importerConfig)

	// AZ Login (assuming an existing and valid az cli configuration, this step is redundant)
	// azure.AzLogin()

	// Create (or load)  azure service principal
	spCredentials := azure.CreateServicePrincipal(importerConfig)
	logger.Info(spCredentials)

	// Assign permissions for the loaded/created service principal in the target KeyVault
	azure.SetServicePrincipalPermissionsInKeyVault(spCredentials, secretsFile.KeyVaultName)

	// Create KeyVault client
	kvClient := keyvault.GetKeyVaultClient(spCredentials)
	kvVaultUri := keyvault.GetKeyVaultUri(secretsFile.KeyVaultName)

	// Start importing process
	output := keyvault.CreateSecrets(kvVaultUri, kvClient, secretsFile)

	logger.Info("***********************************")
	for _, ids := range output {
		logger.Infof("Secret id: %s", ids)
	}
}
