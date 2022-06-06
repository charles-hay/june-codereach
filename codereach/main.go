package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type SlackBotRequest struct {
	Challenge string `json:"challenge"`
}

var (
	username = os.Getenv("DBUSERNAME")
	password = os.Getenv("DBPASSWORD")
	dbname   = os.Getenv("DBNAME")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// resp, err := http.Get(DefaultHTTPGetAddress)
	var slackBotRequest SlackBotRequest

	err := json.Unmarshal([]byte(request.Body), &slackBotRequest)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}

	output, err := json.Marshal(slackBotRequest)
	if err != nil {
		fmt.Println(err)
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprint(string(output)),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
