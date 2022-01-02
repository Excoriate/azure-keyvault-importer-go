package secrets

func IsSecretFileValid(fileType SecretFileType) bool {
	if fileType.KeyVaultName == "" {
		logger.Error("Error. The Azure KeyVault name is required to continue.")
		return false
	}

	if len(fileType.Secrets) == 0 {
		logger.Error("Error. The secrets [] collection was detected empty.")
		return false
	}

	secrets := fileType.Secrets

	// Checks secret collection
	for _, secret := range secrets {
		if secret.SecretValue == "" || secret.SecretName == "" {
			logger.Error("Error. Secret configuration inconsistent. Secret Value and Secret Names are mandatory fields.")
		}

		if secret.IsOverrideAllowed {
			// TODO: No prompt needed here, perhaps for future releases?
			logger.Warnf("Warning: secret %s will be overwritten in azure KeyVault %s ", secret.SecretName, fileType.KeyVaultName)
		}
	}

	return true
}
