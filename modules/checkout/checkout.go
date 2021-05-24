package checkout

import (
	"checkout-service/modules/checkout/model"
	"checkout-service/utls/entity"
)

func (DI *checkout) CheckoutAll(requestCheckout *model.RequestCheckout) (model.ResponsesCheckout, error) {

	// cek promotion
	item, totalPrice := cekPromotion(requestCheckout)

	response := model.ResponsesCheckout{
		Product:    model.ResponseCheckout{
			Quantity:    requestCheckout.Quantity,
			SkuProduct:  requestCheckout.SkuProduct,
			ProductName: requestCheckout.ProductName,
			PromoItem:   item,
		},
		TotalPrice: totalPrice,
	}

	return response, nil
}

func cekPromotion(requestCheckout *model.RequestCheckout) (string, float64) {

	if requestCheckout.ProductName == entity.MacBookPro {
		// Promo buy macbook pro free raspberry
		return entity.Raspbery, entity.PriceMacBookPro
	} else if requestCheckout.ProductName == entity.GoogleHome && requestCheckout.Quantity >= 3 {
		//	buy 3 for price 2
		return "", entity.PriceGoogleHome * 2
	} else if requestCheckout.ProductName == entity.AlexaSpeaker && requestCheckout.Quantity >= 3 {
		//	buy 3 discount 10 %
		totalPrice := entity.PriceAlexaSpeaker * float64(requestCheckout.Quantity)
		return "", totalPrice * 0.1
	}

	return "", 0
}
