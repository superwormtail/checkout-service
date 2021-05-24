package checkout

import "checkout-service/modules/checkout/model"

type checkout struct {
}

func Initialize() *checkout {
	return &checkout{}
}

type Usecase interface {
	CheckoutAll(requestCheckout *model.RequestCheckout) (model.ResponsesCheckout, error)
}
