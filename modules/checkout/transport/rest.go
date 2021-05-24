package transport

import (
	"checkout-service/modules/checkout"
	"checkout-service/modules/checkout/model"
	"checkout-service/utls/echoserver"
	"checkout-service/utls/echoserver/response"
	"checkout-service/utls/entity"
	"checkout-service/utls/errors"
	"github.com/labstack/echo"
	"net/http"
)

type Rest struct {
	usecase checkout.Usecase
}

func NewRest(usecase checkout.Usecase) {
	transport := Rest{usecase}

	checkoutGroup := echoserver.GetServer().Group("/v1/checkout")

	checkoutGroup.POST("/checkingout", transport.checkingout)
}

func (DI *Rest) checkingout(c echo.Context) error {
	var (
		err error
		req = new(model.RequestCheckout)
	)

	if err = c.Bind(req); err != nil {
		return response.Error(c, errors.New(errors.Badrequest, errors.Message(entity.TransportBindError)))
	}

	result, err := DI.usecase.CheckoutAll(req)
	if err != nil {
		return err
	}

	return response.Render(c, http.StatusOK, response.MessageEmpty, result)
}