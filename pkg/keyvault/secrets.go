package keyvault

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/keyvault/keyvault"
	"github.com/Excoriate/azure-keyvault-importer-go/pkg/lib/config"
	"github.com/Excoriate/azure-keyvault-importer-go/pkg/lib/secrets"
	"time"
)

var logger = config.GetLogger()

type AzureKeyVaultSecret struct {
	id    string
	name  string
	value string
}

func CreateSecrets(vaultUri string, keyvaultClient keyvault.BaseClient, secrets secrets.SecretFileType) []string {
	var secretIdsCreated []string

	for i, secret := range secrets.Secrets {
		logger.Warnf("Attempting to create/import secret $s, named %s in keyVault -> %s", i, secret.SecretName, secrets.KeyVaultName)
		time.Sleep(2 * time.Second)

		if !secret.IsOverrideAllowed {
			_, result := GetSecretByName(vaultUri, keyvaultClient, secret.SecretName)
			if result {
				// DeleteSecretByName(vaultUri, keyvaultClient, secret.SecretName)
				logger.Warnf("Ignoring secret %s", secret.SecretName)
				continue
			}
		}

		// Set operation in KeyVault
		res, err := keyvaultClient.SetSecret(context.Background(), vaultUri, secret.SecretName, keyvault.SecretSetParameters{Value: &secret.SecretValue})

		if err != nil {
			logger.Errorf("Error. Cannot create secret %s , %v", secret.SecretName, err.Error())
		}

		logger.Infof("Added Secret : %s , Id : %s", secret.SecretName, *res.ID)
		secretIdsCreated = append(secretIdsCreated, *res.ID)
	}

	logger.Info("All secrets created successfully")

	return secretIdsCreated
}

func GetSecretByName(vaultUri string, keyvaultClient keyvault.BaseClient, secretName string) (AzureKeyVaultSecret, bool) {
	res, err := keyvaultClient.GetSecret(context.Background(), vaultUri, secretName, "") //always the latest

	secretFromKv := AzureKeyVaultSecret{}
	if err != nil {
		logger.Warnf("Error. Cannot get secret %s , %v", secretName, err.Error())
		return secretFromKv, false
	}

	secretFromKv.id = *res.ID
	secretFromKv.value = *res.Value
	secretFromKv.name = secretName

	return secretFromKv, true
}
