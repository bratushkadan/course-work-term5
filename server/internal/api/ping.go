package api

import (
	"net/http"
	"time"

	floralApi "floral/generated/api"

	"github.com/gin-gonic/gin"
)

func (*Impl) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, floralApi.PingResponse{Ts: time.Now().UnixMilli()})
}
func (impl *Impl) GetPing(c *gin.Context) {
	impl.GetHealth(c)
}
