package writer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/usecase"
)

func ToLoginOutputType(c *gin.Context, output *usecase.LoginResponse) {
	c.JSON(http.StatusOK, output)
}

func ToRefreshTokenOutputType(c *gin.Context, output *usecase.RefreshTokenResponse) {
	c.JSON(http.StatusOK, output)
}
