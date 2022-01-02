package keyvault

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/keyvault/keyvault"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Excoriate/azure-keyvault-importer-go/pkg/azure"
)

func GetKeyVaultClient(servicePrincipal azure.ServicePrincipalType) keyvault.BaseClient {
	keyvaultClient := keyvault.New()

	clientCredentialConfig := auth.NewClientCredentialsConfig(servicePrincipal.AppID, servicePrincipal.Password, servicePrincipal.Tenant)

	// From SDK NewClientCredentialsConfig generates a object to azure control plane
	// (By default Resource is set to management.azure.net)
	// There below line was added to access the azure data plane
	// Which is required to access secrets in keyvault

	clientCredentialConfig.Resource = "https://vault.azure.net"
	authorizer, err := clientCredentialConfig.Authorizer()

	if err != nil {
		fmt.Printf("Error occured while creating azure KV authroizer %v ", err.Error())

	}
	keyvaultClient.Authorizer = authorizer

	return keyvaultClient
}
