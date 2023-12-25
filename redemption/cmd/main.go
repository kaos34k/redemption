package main

import (
	"context"
	"fmt"
	"redemption/internal/adapter"
	"redemption/internal/app"
	"redemption/internal/domain"
	"redemption/internal/usecase"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type LambdaHandler struct {
	myApp app.MyApp
}

func NewLambdaHandler(pointRepository domain.PointRepository) *LambdaHandler {
	pointUsecase := usecase.NewPointUsecase(pointRepository)
	myApp := app.NewMyApp(*pointUsecase)
	return &LambdaHandler{
		myApp: *myApp,
	}
}

func (h *LambdaHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userID := request.PathParameters["id"]
	err := h.myApp.HandleRequest(userID)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("Dato eliminado con exito"),
	}, nil
}

func main() {
	tableName := "Point"
	sess := session.Must(session.NewSession())
	actualDynamoDBClient := dynamodb.New(sess)
	pointRepository := adapter.NewDynamoDBRepository(tableName, actualDynamoDBClient)
	handler := NewLambdaHandler(pointRepository)
	lambda.Start(handler.HandleRequest)
}
