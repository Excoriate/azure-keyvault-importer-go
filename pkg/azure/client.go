package azure

import (
	"github.com/Excoriate/azure-keyvault-importer-go/pkg/lib/config"
	"github.com/Excoriate/azure-keyvault-importer-go/pkg/lib/utils"
	"os"
)

var (
	logger = config.GetLogger()
)

func SetCredentialsInProcess(config config.ConfigType) {
	_ = os.Setenv("AZURE_TENANT_ID", config.TenantId)
	_ = os.Setenv("AZURE_SUBSCRIPTION_ID", config.SubscriptionId)
	_ = os.Setenv("AZURE_LOCATION", config.Location)

	args := []string{}
	output, err := utils.RunCMD("printenv", args, true)

	if err != nil {
		logger.Fatal("Error. Cannot set required AZURE environment variables. Error details: %s ", err.Error())
	}

	logger.Info(output)
}

func AzLogin() {
	args := []string{"login"}
	_, err := utils.RunCMD("az", args, true)

	if err != nil {
		logger.Fatal("Error. Cannot login in azure. Error details: %s ", err)
	}

	logger.Info("Successfully logged in into Azure")
}
