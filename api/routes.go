package api

import (
	"github.com/Kawaii-Konnections-KK-Limited/subscription/api/handlers"
	"github.com/Kawaii-Konnections-KK-Limited/subscription/api/middleware"
	"github.com/Kawaii-Konnections-KK-Limited/subscription/utils"
	"github.com/gin-gonic/gin"
)

func InitRoutes(gupFunc func(token *string) *[]string, vtFunc func(token *string) bool, SubUpdateSender func(token *string)) *gin.Engine {
	if gupFunc != nil {
		utils.GetUserProfiles = gupFunc
	}
	if vtFunc != nil {
		utils.VerifyToken = vtFunc
	}
	if SubUpdateSender != nil {
		utils.SubUpdateSender = SubUpdateSender
	}
	r := gin.Default()
	r.Use(middleware.AuthTokenMiddleware())
	r.GET("/:token", handlers.SubscriptionHandler)
	return r
}
