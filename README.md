# Checkout Backend API

* Content root path main : 
```sh
cmd/rest/main.go
```
* To run :
```sh
$ go run cmd/rest/main.go
```
* Script cURL for checkout with promotion :
```sh
curl --location --request POST 'localhost:8080/v1/checkout/checkingout' \
--header 'Content-Type: application/json' \
--data-raw '{
    "sku_product": "",
    "product_name": "Alexa Speaker",
    "quantity": 3
}'
```
