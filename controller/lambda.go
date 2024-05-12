package controller

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/openkrafter/anytore-backend/logger"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger.Logger.Debug("Processing request", logger.Attr("req", req))

	r := gin.Default()
	SetCors(r)
	SetCSP(r)
	RegisterRoutes(r)

	bodyReader := bytes.NewBufferString(req.Body)
	httpReq, err := http.NewRequest(req.HTTPMethod, req.Path, bodyReader)
	if err != nil {
		logger.Logger.Error("Failed to create HTTP request", logger.ErrAttr(err))
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	for key, values := range req.MultiValueHeaders {
		for _, value := range values {
			httpReq.Header.Add(key, value)
		}
	}

	logger.Logger.Debug("HTTP request created", logger.Attr("httpReq Method", httpReq.Method))
	logger.Logger.Debug("HTTP request created", logger.Attr("httpReq URL", httpReq.URL))
	logger.Logger.Debug("HTTP request created", logger.Attr("httpReq Header", httpReq.Header))
	logger.Logger.Debug("HTTP request created", logger.Attr("httpReq Body", bodyReader.String()))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)
	logger.Logger.Debug("HTTP request processed", logger.Attr("httpResp Code", w.Code))
	logger.Logger.Debug("HTTP request processed", logger.Attr("httpResp Header", w.Header()))
	logger.Logger.Debug("HTTP request processed", logger.Attr("httpResp Body", w.Body.String()))

	return events.APIGatewayProxyResponse{
		StatusCode: w.Code,
		Body:       w.Body.String(),
		Headers:    convertHeaders(w.Header()),
	}, nil
}

func convertHeaders(h http.Header) map[string]string {
	headers := map[string]string{}
	for k, v := range h {
		headers[k] = strings.Join(v, ",")
	}
	return headers
}
