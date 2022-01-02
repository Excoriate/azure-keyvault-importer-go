package azure

import (
	"encoding/json"
	"fmt"
	"github.com/Excoriate/azure-keyvault-importer-go/pkg/lib/config"
	"github.com/Excoriate/azure-keyvault-importer-go/pkg/lib/utils"
	"io/ioutil"
)

type ServicePrincipalType struct {
	AppID       string `json:"appId"`
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Tenant      string `json:"tenant"`
}

func createServicePrincipalInAzure(config config.ConfigType) ServicePrincipalType {
	logger.Info("Creating service principal...")

	var (
		servicePrincipalName = "sp-prip-keyvault-importer"
		subscriptionPath     = fmt.Sprintf("/subscriptions/%s/resourceGroups/%s", config.SubscriptionId, config.ResourceGroup)
	)

	args := []string{"ad", "sp", "create-for-rbac", "--name", servicePrincipalName, "--role", "Contributor", "--scopes", subscriptionPath}
	output, err := utils.RunCMD("az", args, true)

	if err != nil {
		logger.Fatal("Error. Cannot login in azure. Error details: %s ", err)
	}

	// Parsing output
	data := []byte(output)
	var servicePrincipal ServicePrincipalType

	err = json.Unmarshal([]byte(data), &servicePrincipal)

	if err != nil {
		logger.Fatal("Error. Cannot unmarshall output from Azure API call")
	}

	return servicePrincipal
}

func SetServicePrincipalPermissionsInKeyVault(servicePrincipal ServicePrincipalType, keyVaultName string) {
	args := []string{"keyvault", "set-policy", "--name", keyVaultName, "--spn", servicePrincipal.AppID, "--secret-permissions", "get", "list", "set", "delete"}
	_, err := utils.RunCMD("az", args, true)

	if err != nil {
		logger.Fatal("Error. Cannot set required permissions for service principal: %s, error detials: %s", servicePrincipal.Name, err.Error())
	}
}

func saveServicePrincipalCredentials(servicePrincipal ServicePrincipalType) {
	file, _ := json.MarshalIndent(servicePrincipal, "", " ")
	err := ioutil.WriteFile("importer/sp-creds.json", file, 0644)

	if err != nil {
		logger.Fatal("Error. Cannot create SP empty file into local filesystem")
	}
}

func CreateServicePrincipal(config config.ConfigType) ServicePrincipalType {
	var servicePrincipal ServicePrincipalType

	if utils.IsFileExist("importer/sp-creds.json") {
		// No need to create a new service principal. Load what was already found
		spLoadedFromFs, _ := ioutil.ReadFile("./importer/sp-creds.json")

		err := json.Unmarshal([]byte(spLoadedFromFs), &servicePrincipal)

		if err != nil {
			logger.Fatal("Error. Cannot load sp-creds.json credentials from filesystem, error %s", err.Error())
		}

		return servicePrincipal
	}

	sp := createServicePrincipalInAzure(config)
	saveServicePrincipalCredentials(sp)

	return sp
}
