package http

import (
	"net/http"

	"github.com/santowijaya28/cleanarch-sample/pkg/response"
)

func (d *HttpDelivery) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// TODO: add logic
	rsp := response.Response{
		Data:           "register user",
		Error:          "",
		HttpStatusCode: http.StatusOK,
	}

	response.WriteSuccessResponse(w, rsp)
}
