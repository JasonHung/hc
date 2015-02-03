package endpoint

import (
	"github.com/brutella/hap/netio"
	"github.com/brutella/hap/netio/pair"
	"github.com/brutella/log"
	"io"
	"net/http"
)

// Handles the /pairings endpoint
//
// This endpoint is not session based and the same for all connections
type Pairing struct {
	http.Handler

	controller *pair.PairingController
}

func NewPairing(controller *pair.PairingController) *Pairing {
	handler := Pairing{
		controller: controller,
	}

	return &handler
}

func (handler *Pairing) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("[VERB] POST /pairings")
	response.Header().Set("Content-Type", netio.HTTPContentTypePairingTLV8)

	res, err := pair.HandleReaderForHandler(request.Body, handler.controller)

	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		io.Copy(response, res)
	}
}
