package keyvault

import "fmt"

func GetKeyVaultUri(keyVaultName string) string {
	return fmt.Sprintf("https://%s.vault.azure.net", keyVaultName)
}
