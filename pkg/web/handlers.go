package web

import (
	"net/http"

	"github.com/hellofresh/janus/pkg/errors"
	"github.com/hellofresh/janus/pkg/response"
	log "github.com/sirupsen/logrus"
)

// Home handler is just a nice home page message
func Home(version string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, "Welcome to Janus v"+version)
	}
}

// NotFound handler is called when no route is matched
func NotFound(w http.ResponseWriter, r *http.Request) {
	notFoundError := errors.ErrRouteNotFound
	response.JSON(w, notFoundError.Code, notFoundError)
}

// RecoveryHandler handler is used when a panic happens
func RecoveryHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	switch internalErr := err.(type) {
	case *errors.Error:
		log.WithFields(log.Fields{"code": internalErr.Code, "error": internalErr.Error()}).
			Info("Internal error handled")
		response.JSON(w, internalErr.Code, internalErr.Error())
	default:
		log.WithField("error", err).Error("Internal server error handled")
		response.JSON(w, http.StatusInternalServerError, err)
	}
}
