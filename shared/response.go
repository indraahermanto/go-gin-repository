package shared

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, status bool, data interface{}) {
	var httpStatus int
	var res interface{}
	if status == true {
		httpStatus = http.StatusOK
		res = gin.H{
			"status": status,
			"data":   data,
		}
	} else {
		httpStatus = http.StatusBadRequest
		res = gin.H{
			"status": status,
			"error":  data,
		}
	}

	c.JSON(httpStatus, res)
	return
}
