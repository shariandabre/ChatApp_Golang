package UserRoutes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorldUser(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World User!!!"

	c.JSON(http.StatusOK, resp)
}
