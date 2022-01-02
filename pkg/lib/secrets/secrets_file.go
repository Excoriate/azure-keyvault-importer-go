package secrets

import (
	"encoding/json"
	"github.com/Excoriate/azure-keyvault-importer-go/pkg/lib/config"
	"io/ioutil"
)

var logger = config.GetLogger()

type SecretFileType struct {
	KeyVaultName string `json:"keyVaultName"`
	Secrets      []struct {
		SecretName        string `json:"secretName"`
		SecretValue       string `json:"secretValue"`
		IsOverrideAllowed bool   `json:"isOverrideAllowed"`
	} `json:"secrets"`
}

func LoadSecretsJSONFile(jsonFilePath string) (SecretFileType, error) {
	// FIXME: Does not recognize filepath once it is sourced from the function argument
	// data, err := ioutil.ReadFile(fileContent)
	data, err := ioutil.ReadFile("./importer/secrets_example.json")

	if err != nil {
		logger.Error(err)
		logger.Fatal("Error. Cannot found json file with path " + jsonFilePath)
	}

	var secretFile SecretFileType

	_ = json.Unmarshal([]byte(data), &secretFile)

	return secretFile, nil
}
