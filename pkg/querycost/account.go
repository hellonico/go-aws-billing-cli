package querycost

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
)

func FetchAccountsMap(profileName string) (map[string]string, error) {

	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profileName))

	svc := organizations.NewFromConfig(cfg)
	input := &organizations.ListAccountsInput{}

	result, err := svc.ListAccounts(context.Background(), input)

	accountsMap := make(map[string]string)
	if err == nil {
		for _, account := range result.Accounts {
			accountsMap[*account.Id] = *account.Name
		}
	}
	return accountsMap, err
}
