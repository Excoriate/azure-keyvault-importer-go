package config

import (
	"bufio"
	"os"
	"strings"
)

var logger = GetLogger()

const SettingKeyTenantId = "TENANT_ID"
const SettingKeySubscriptionId = "SUBSCRIPTION_ID"
const SettingsKeyResourceGroup = "RESOURCE_GROUP_NAME"
const SettingsKeySecretsSourcePath = "SOURCE_SECRETS_PATH"
const SettingsKeyLocation = "LOCATION"

type ConfigType struct {
	TenantId          string
	SubscriptionId    string
	ResourceGroup     string
	SecretsSourcePath string
	Location          string
}

func GetConfigFileLines(settingsFile *os.File) []string {

	var linesInConfigFile []string
	scanner := bufio.NewScanner(settingsFile)
	for scanner.Scan() {
		linesInConfigFile = append(linesInConfigFile, scanner.Text())

		logger.Info(linesInConfigFile)
	}

	return linesInConfigFile
}

func GetConfigFile() *os.File {
	file, err := os.Open("./import.config")
	currentPath, _ := os.Getwd()

	if err != nil {
		logger.Fatal("Cannot file configuration file [import.config] in path" + currentPath)
	}

	return file
}

func GetSettingsAndValidate(lines []string) ConfigType {
	settingsSize := len(lines)
	config := ConfigType{}

	for i := 0; i < settingsSize; i++ {
		var setting = lines[i]
		if setting != "" {
			settingKey := strings.Split(setting, "=")[0]
			settingValue := strings.Split(setting, "=")[1]
			if settingKey != SettingKeySubscriptionId && settingKey != SettingKeyTenantId && settingKey != SettingsKeyResourceGroup && settingKey != SettingsKeySecretsSourcePath && settingKey != SettingsKeyLocation {
				logger.Fatal("Invalid settings. Could not found either TENANT_ID, SOURCE_SECRETS_PATH, LOCATION, RESOURCE_GROUP_NAME or SUBSCRIPTION_ID")
			}

			if settingKey == SettingKeyTenantId {
				config.TenantId = settingValue
				logger.Info("Tenant Id found: " + config.TenantId)
			}

			if settingKey == SettingsKeyLocation {
				config.Location = settingValue
				logger.Info("Location found: " + config.Location)
			}

			if settingKey == SettingKeySubscriptionId {
				config.SubscriptionId = settingValue
				logger.Info("Subscription Id found: " + config.SubscriptionId)
			}

			if settingKey == SettingsKeyResourceGroup {
				config.ResourceGroup = settingValue
				logger.Info("Resource Group found: " + config.ResourceGroup)
			}

			if settingKey == SettingsKeySecretsSourcePath {
				config.SecretsSourcePath = settingValue
				logger.Info("Secrets source path found: " + config.SecretsSourcePath)
			}
		}
	}

	return config
}

func LoadSettings() ConfigType {
	settingsFile := GetConfigFile()
	settingsContent := GetConfigFileLines(settingsFile)

	return GetSettingsAndValidate(settingsContent)
}
