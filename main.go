package main

import (
	"encoding/json"
	"os"

	rz "github.com/adrianpk/rezerw/core"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	ct    = "Content-Type"
	ctVal = "application/json"
)

// Account - Account
type Account struct {
	ID          string `schema:"id"`
	Name        string `json:"Name" schema:"name"`
	Description string `json:"Description" schema:"description"`
	AccountType string `json:"AccountType" schema:"account-type"`
}

func findAll() (events.APIGatewayProxyResponse, error) {
	// Config
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return rz.RenderError(err, "Error while retrieving credentials")
	}
	// DynamoDB
	svc := dynamodb.New(cfg)
	req := svc.ScanRequest(&dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	})
	dbRes, err := req.Send()
	if err != nil {
		return rz.RenderError(err, "Error while scanning data")
	}
	// Populate
	accounts := make([]Account, 0)
	for _, item := range dbRes.Items {
		accounts = append(accounts, Account{
			ID:   *item["ID"].S,
			Name: *item["Name"].S,
		})
	}
	// Marshall
	res, err := json.Marshal(accounts)
	if err != nil {
		return rz.RenderError(err, "Error while decoding response")
	}
	// Response
	return rz.RenderOk(res)
}

func main() {
	lambda.Start(findAll)
}
