package awsinventory

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

var (
	startURL     string
	orgAccountID string
	roleName     string
)

func main() {
	flag.StringVar(&startURL, "start-url", "", "AWS SSO Start URL")
	flag.StringVar(&orgAccountID, "org-account-id", "", "AWS SSO account ID")
	flag.StringVar(&roleName, "role-name", "", "Role Name to assume")

	if startURL == "" || orgAccountID == "" || roleName == "" {
		flag.Usage()
		os.Exit(1)
	}

}

func GetConfig() (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(""))
	if err != nil {
		return nil, err
	}

	sharedConfig := getSharedConfig(&cfg)
	validateSharedConfig(sharedConfig)

	reauthenticate := false

	staticCredentials, err := loadExistingCredentials()
	if err != nil {
		reauthenticate = true
	}

	if !reauthenticate {
		credsProvider := credentials.StaticCredentialsProvider{
			Value: staticCredentials,
		}

		cfg.Credentials = credsProvider

		if !checkExistingCredentials(&cfg) {
			reauthenticate = true
		}
	}
	if reauthenticate {
		err = reloginWorkflow(&cfg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return &cfg, nil
}

func getSharedConfig(cfg *aws.Config) (sharedConfig config.SharedConfig) {
	cfgSources := cfg.ConfigSources
	for _, cfgSource := range cfgSources {
		foundSharedConfig, ok := cfgSource.(config.SharedConfig)
		if ok {
			sharedConfig = foundSharedConfig
		}
	}
	return sharedConfig
}

func validateSharedConfig(sharedConfig config.SharedConfig) {

}
func loadExistingCredentials() (aws.Credentials, error) {

}
func checkExistingCredentials(cfg *aws.Config) bool {

}
func reloginWorkflow(cfg *aws.Config) error {

}
