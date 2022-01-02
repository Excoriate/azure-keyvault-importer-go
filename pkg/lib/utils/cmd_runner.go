package utils

import (
	"github.com/Excoriate/azure-keyvault-importer-go/pkg/lib/config"
	"os/exec"
	"strings"
)

var logger = config.GetLogger()

func RunCMD(mainCommandOrApp string, args []string, debug bool) (out string, err error) {

	cmd := exec.Command(mainCommandOrApp, args...)
	stdout, err := cmd.Output()

	if debug {
		logger.Info(strings.Join(cmd.Args[:], " "))

		if err != nil {
			logger.Errorf("RunCMD ERROR, %s", err.Error())
			return string(stdout), err
		}
	}

	return string(stdout), nil
}
