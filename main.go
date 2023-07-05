package main

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/matelang/jwt-go-aws-kms/v2/jwtkms"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

var KmsKeyID = os.Getenv("KMS_JWT_KEY_ID")

func HandleRequest(ctx context.Context, request events.APIGatewayCustomAuthorizerRequestTypeRequest) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
	awsConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Errorf("aws config failed to load %s", err)

		return events.APIGatewayV2CustomAuthorizerSimpleResponse{
			IsAuthorized: false,
		}, nil
	}

	kmsConfig := jwtkms.NewKMSConfig(kms.NewFromConfig(awsConfig), KmsKeyID, false)

	claims := jwt.RegisteredClaims{}

	_, err = jwt.ParseWithClaims(request.Headers["authorization"],
		&claims, func(token *jwt.Token) (interface{}, error) {
			return kmsConfig, nil
		})
	if err != nil {
		log.Errorf("can not parse/verify token %s", err)

		return events.APIGatewayV2CustomAuthorizerSimpleResponse{
			IsAuthorized: false,
		}, nil
	}

	log.Infof("validated token with claims: %v", claims)

	return events.APIGatewayV2CustomAuthorizerSimpleResponse{
		IsAuthorized: true,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
