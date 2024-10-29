package http

import (
	"net/http"

	"github.com/santowijaya28/cleanarch-sample/pkg/response"
)

func (d *HttpDelivery) Ping(w http.ResponseWriter, r *http.Request) {
	rsp := response.Response{
		Data:           "pong",
		Error:          "",
		HttpStatusCode: http.StatusOK,
	}

	response.WriteSuccessResponse(w, rsp)
}
