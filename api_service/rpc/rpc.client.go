package rpc

import (
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	"sync"
	"log"
	"net/http"
)

const addr = "localhost:9020"

var client *SubscriptionServiceClient

func GetRpcClientInstance() SubscriptionServiceClient {
	sync.Once.Do(func() {
		client = &newRpcClient()
	})

	return client
}

func newRpcClient() *SubscriptionServiceClient {
	// set up a connection to the server
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	return &NewSubscriptionServiceClient(conn)
}

func PerformForId(ctx echo.Context, handler func(c SubscriptionServiceClient, sID *SubscriptionId)) error {
	sID := new(SubscriptionId)
	if err := ctx.Bind(sID); err != nil {
		return err
	}

	c := GetRpcClientInstance()
	res, err := handler(&c, sID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusAccepted, res)
}
