package lib

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func NewParamError(err error, message ...string) error {
	if err == nil {
		return nil
	}
	e, ok := err.(gin.Error)
	if ok && e.Type == gin.ErrorTypeBind {
		return e
	}
	if len(message) == 0 {
		message = []string{"param error"}
	}
	msg := strings.Join(message, ", ")
	return &gin.Error{
		Err:  err,
		Type: gin.ErrorTypeBind,
		Meta: map[string]interface{}{"message": msg},
	}
}

func NewServiceError(err error, message ...string) error {
	if err == nil {
		return nil
	}
	e, ok := err.(gin.Error)
	if ok && e.Type == gin.ErrorTypePrivate {
		return e
	}
	if len(message) == 0 {
		message = []string{"service error"}
	}
	msg := strings.Join(message, ", ")
	return &gin.Error{
		Err:  err,
		Type: gin.ErrorTypePrivate,
		Meta: map[string]interface{}{"message": msg},
	}
}
