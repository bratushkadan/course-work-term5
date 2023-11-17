package api

import (
	"net/http"
	"time"

	"floral/internal/db"

	floralApi "floral/generated/api"

	"github.com/gin-gonic/gin"
)

type Impl struct{}

func (*Impl) GetPing(c *gin.Context) {
	ping := floralApi.Ping{}
	ping.Ts = time.Now().UnixMilli()

	c.JSON(http.StatusOK, ping)
}
func (api *Impl) GetHealth(c *gin.Context) {
	api.GetPing(c)
}
func (*Impl) GetV1Users(c *gin.Context) {
	users, err := db.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	if users == nil {
		c.JSON(http.StatusOK, []struct{}{})
		return
	}
	c.JSON(http.StatusOK, users)
}
func (*Impl) GetV1UsersId(c *gin.Context, id int64) {
}
