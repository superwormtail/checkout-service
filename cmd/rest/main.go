package main

import (
	"checkout-service/config"
	checkout2 "checkout-service/modules/checkout"
	checkout "checkout-service/modules/checkout/transport"
	"checkout-service/utls/echoserver"
)

func main() {

	config.Initialization()

	// run echo service
	echoserver.InitServer()

	// init service
	checkout.NewRest(checkout2.Initialize())

	// start server
	go echoserver.StartServer()

	// Shutdown server gracefully
	echoserver.Shutdown()
}
