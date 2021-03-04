//Guards throws errors or panics when error occur
package guard

import (
	"github.com/kuritka/golic/utils/log"
	"net/http"
	"os"

)

var logger = log.Log

func HttpThrowServerError(w http.ResponseWriter, err error, message string, v ...interface{}) {
	HttpThrowError(w, http.StatusInternalServerError, message, v)
	logger.Err(err).Msgf(message, v)
}

func HttpThrowError(w http.ResponseWriter, httpCode int, message string, v ...interface{}) {
	http.Error(w, message, httpCode)
	logger.Error().Msgf(message, v)
}

func FailOnError(err error, message string, v ...interface{}) {
	if err != nil {
		logger.Panic().Err(err).Msgf(message, v)
	}
}

// Must exit on error.
func Must(err error) {
	if err == nil {
		return
	}

	logger.Error().Msgf("ERROR: %+v", err)
	os.Exit(1)
}