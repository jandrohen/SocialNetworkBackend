package secretmanager

import (
	"encoding/json"
	"fmt"

	"SocialNetworkBackend/awsgo"
	"SocialNetworkBackend/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	var dataSecrets models.Secret
	fmt.Println("> Getting secret" + secretName )

	// Create a Secrets Manager client
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return dataSecrets, err
	}

	json.Unmarshal([]byte(*key.SecretString), &dataSecrets)
	fmt.Println("> Secret obtained" + secretName)

	return dataSecrets, nil
}