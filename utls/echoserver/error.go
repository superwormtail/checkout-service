package echoserver

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type response struct {
	Message interface{} `json:"message"`
}

func Handler(err error, c echo.Context) {
	var (
		code        = http.StatusInternalServerError
		msg interface{}
	)

	if GetServer().Debug {
		msg = err.Error()
		switch err.(type) {
		case *echo.HTTPError:
			code = err.(*echo.HTTPError).Code
		}
	} else {
		switch e := err.(type) {
		case *echo.HTTPError:
			code = e.Code
			msg = e.Message
			if e.Internal != nil {
				msg = fmt.Sprintf("%v, %v", err, e.Internal)
			}
		default:
			msg = http.StatusText(code)
		}

		if _, ok := msg.(string); ok {
			msg = response{Message: msg}
		}
	}

	if !c.Response().Committed {
		if c.Request().Method == "HEAD" {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}

		if code == http.StatusInternalServerError {
			// Log error message
			log.Fatalf("internal server error")
		}

		if err != nil {
			log.Panicf("got an error while serve data, error: %s", err)
		}
	}
}
