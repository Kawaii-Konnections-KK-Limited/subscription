package handlers

import (
	"github.com/Kawaii-Konnections-KK-Limited/subscription/utils"
	"github.com/gin-gonic/gin"
)

const plaintextContentType = "text/plain"

func SubscriptionHandler(c *gin.Context) {
	token, _ := c.Get("token")
	subs := utils.GetUserProfiles(token.(*string))
	c.Writer.Header().Set("Content-Type", plaintextContentType)
	c.String(200, utils.SubscriptionBuilder(subs))
	utils.SubUpdateSender(token.(*string))
	return
}
