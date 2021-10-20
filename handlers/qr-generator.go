package handlers

import (
	"fmt"
	"image/png"
	"net/http"

	"github.com/joaomarcuslf/qr-generator/usecases"
)

func GenerateQr(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("dataString")

	qrCode, err := usecases.GenerateQr(data)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		w.Write([]byte(fmt.Sprintf("error: %s", err)))

		return
	}

	png.Encode(w, *qrCode)
}
