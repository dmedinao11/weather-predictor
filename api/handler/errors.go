package handler

import (
	"errors"
	"github.com/dmedinao11/weather-predictor/internal/apperrrors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	errorDto struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
)

func handleError(context *gin.Context, err error) {
	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, apperrrors.ErrNotFound):
		status = http.StatusNotFound
	case errors.Is(err, apperrrors.ErrScanFailed):
		status = http.StatusInternalServerError
	case errors.Is(err, apperrrors.ErrInvalidPathParam) || errors.Is(err, apperrrors.ErrParsingPathParam):
		status = http.StatusBadRequest
	}

	dto := errorDto{
		Status:  status,
		Message: err.Error(),
	}

	context.JSON(status, dto)
}
