package http

import (
	"net/http"

	"github.com/georace/recentLocations/router"
)

// Health Details the health of a service
type Health struct {
	Status string
}

// HealthHandler godoc
// @Summary Service Health
// @Description Service Health
// @Produce  json
// @Success 200 {object} http.Health
// @Failure 404 {object} http.ErrorResponse
// @Router /healthz [get]
// @Tags System
func HealthHandler(w http.ResponseWriter, r *http.Request) {

	NewResponse(w, Health{Status: "OK"})
}

// Routes from the http package
var Routes = router.RoutePrefix{
	"",
	[]router.Route{
		router.Route{"Health", "GET", "/healthz", HealthHandler, false},
	},
}
