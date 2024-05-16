package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r := gin.Default()

	r.POST("/urls", Generate)

	//r.Run(":8080")
	//return events.APIGatewayProxyResponse{}, nil

	ginLambda = ginadapter.New(r)
	return ginLambda.Proxy(request)
}

func main() {
	lambda.Start(Handler)
	//req := events.APIGatewayProxyRequest{
	//	MultiValueHeaders:               nil,
	//	QueryStringParameters:           nil,
	//	MultiValueQueryStringParameters: nil,
	//	PathParameters:                  map[string]string{"url": "40c55c4a-269a-45be-9e13-ab0c2f424b36"},
	//	StageVariables:                  nil,
	//	RequestContext:                  events.APIGatewayProxyRequestContext{},
	//	Body:                            "",
	//	IsBase64Encoded:                 false,
	//}
	//_, _ = Handler(req)
}
