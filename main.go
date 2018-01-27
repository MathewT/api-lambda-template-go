package main

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
)

//Response is the return value of func handleRequest
type Response struct {
	content    string
	statusCode int
}

//Environment holds the environment variables.
type Environment struct {
	logLevel string
}

// Handler is the main Lambda function handler and is the entry point of the API
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var err error
	var env Environment

	log := logrus.New()
	log.Formatter = new(logrus.JSONFormatter)

	log.Info("======Handler begin========")
	//create response headers
	h := createResponseHeaders()

	err = setEnvironemnt(log, &env)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal error",
			Headers:    h,
			StatusCode: 500,
		}, err
	}
	//Handle the request
	r, err := handleRequest(log, env, request)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			Headers:    h,
			StatusCode: r.statusCode,
		}, nil
	}

	log.Println("======Handler end======")
	// Success
	return events.APIGatewayProxyResponse{
		Body:       r.content,
		Headers:    h,
		StatusCode: r.statusCode,
	}, nil

}

func createResponseHeaders() map[string]string {
	h := make(map[string]string)
	h["Access-Control-Allow-Origin"] = "*"
	return h
}

// processes the request.
func handleRequest(log *logrus.Logger, env Environment, request events.APIGatewayProxyRequest) (Response, error) {

	log.Println("content-6")
	return Response{
		content:    "content-6",
		statusCode: 200,
	}, nil
}

func setEnvironemnt(log *logrus.Logger, env *Environment) error {
	env.logLevel = os.Getenv("LOG_LEVEL")
	if env.logLevel == "" {
		log.Error("LOG_LEVEL is not set")
		return errors.New("LOG_LEVEL is not set")
	}
	log.Infof("LOG_LEVEL is %s", env.logLevel)
	//TODO: Other env variables get read here

	return nil
}

func main() {
	lambda.Start(Handler)
}
