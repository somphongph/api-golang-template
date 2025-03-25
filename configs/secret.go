package configs

import (
	"log"
	"slices"
	"strings"

	"github.com/spf13/viper"
	"github.com/tripgator/lib-golang-packages/xlogger"
	"github.com/tripgator/lib-golang-packages/xvault"
)

func GetSecret() *Secrets {
	var secret Secrets

	conf := GetConfig()
	appEnv := conf.App.Env
	envlist := []string{"dev", "sit", "uat", "prod"}
	contains := slices.Contains(envlist, appEnv)

	if contains {
		vaultUrl := conf.Service.VaultUrl
		appName := conf.App.Name
		kvName := conf.App.KvName

		vaultSecret, err := xvault.LoadVaultSecret(vaultUrl, appName, kvName, secret)
		if err != nil {
			log.Fatalf("Failed to load vault secret: %s", err)
		}

		secret = *vaultSecret

		xlogger.Infof("Fetching secrets from vault")
	} else {
		// viper
		viper.AddConfigPath(strings.Join([]string{"./configs"}, "/"))
		viper.SetConfigName("secret")
		viper.SetConfigType("json")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Failed to read secret file: %s", err)
		}

		err = viper.Unmarshal(&secret)
		if err != nil {
			log.Fatalf("Failed to unmarshal secret file: %s", err)
		}

		xlogger.Infof("Fetching secrets from env")
	}

	return &secret
}
