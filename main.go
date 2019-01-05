package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	uuid "github.com/satori/go.uuid"
)

var (
	ct    = "Content-Type"
	ctVal = "application/json"
)

type account struct {
	ID          uuid.UUID `json:"id" schema:"id"`
	Name        string    `json:"name" schema:"name"`
	Description string    `json:"description" schema:"description"`
	AccountType string    `json:"accountType" schema:"account-type"`
	OwnerID     uuid.UUID `json:"ownerID" schema:"owner-id"`
	ParentID    uuid.UUID `json:"parentID" schema:"parent-id"`
	Email       string    `json:"email" schema:"email"`
	CreatedBy   uuid.UUID `json:"CreatedBy,omitempty" schema:"-"`
	UpdatedBy   uuid.UUID `json:"UpdatedBy,omitempty" schema:"-"`
	CreatedAt   time.Time `json:"createdAt,omitempty" schema:"-"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty" schema:"-"`
}

func findAll() (events.APIGatewayProxyResponse, error) {
	accounts := sampleAccounts()
	response, err := json.Marshal(accounts)
	// Error
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	// Ok
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			ct: ctVal,
		},
		Body: string(response),
	}, nil
}

func sampleAccounts() []account {
	accounts := make([]account, 0)
	if u1, ok := toUUID("568aabd8-c431-485f-845a-c447083ab287"); ok {
		a1 := account{ID: u1, Name: "Account1"}
		accounts = append(accounts, a1)
	}
	if u2, ok := toUUID("338681c2-fb4b-4448-957a-297729eab4a8"); ok {
		a2 := account{ID: u2, Name: "Account2"}
		accounts = append(accounts, a2)
	}
	return accounts
}

func toUUID(idStr string) (uuid.UUID, bool) {
	u, err := uuid.FromString(idStr)
	if err != nil {
		u, _ = uuid.FromString("00000000-0000-0000-0000-000000000000")
		return u, false
	}
	return u, true
}

func main() {
	lambda.Start(findAll)
}
