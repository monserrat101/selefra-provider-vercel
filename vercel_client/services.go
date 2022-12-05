package vercel_client

import (
	"context"
	"errors"
	"github.com/selefra/selefra-provider-vercel/constants"
	"os"

	vercel "github.com/chronark/vercel-go"
)

func Connect(ctx context.Context, vercelConfig *Config) (*vercel.Client, error) {

	apiToken := os.Getenv(constants.VERCELAPITOKEN)
	team := os.Getenv(constants.VERCELTEAM)

	if vercelConfig.APIToken != constants.Constants_3 {
		apiToken = vercelConfig.APIToken
	}
	if vercelConfig.Team != constants.Constants_4 {
		team = vercelConfig.Team
	}

	if apiToken == constants.Constants_5 {
		return nil, errors.New(constants.Apitokenmustbeconfigured)
	}

	config := vercel.NewClientConfig{Token: apiToken}
	if team != constants.Constants_6 {
		config.Teamid = team
	}
	conn := vercel.New(config)

	return conn, nil
}
