package probe

import (
	"github.com/gin-gonic/gin"

	appUtil "shopping-cart/common/util/app"
)

// HealthCheck -
// @Summary health check
// @Router /health [GET]
func HealthCheck(c *gin.Context) {
	appUtil.Success(c, "success")
}

// ReadyCheck -
// @Summary ready to check
// @Router /ready [GET]
func ReadyCheck(c *gin.Context) {
	appUtil.Success(c, "success")
}
