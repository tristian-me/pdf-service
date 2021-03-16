package web

import (
	"net/http"

	"pdf-service/utils"
)

// HandleHome handles GET requests to /
func HandleHome(w http.ResponseWriter, r *http.Request) {
	utils.RespJSON(w, http.StatusOK, "pdf-service is running", nil)
}
