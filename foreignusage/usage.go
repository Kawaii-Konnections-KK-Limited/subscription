package foreignusage

import "github.com/Kawaii-Konnections-KK-Limited/subscription/api"

func InitService(gupFunc func(token *string) *[]string, vtFunc func(token *string) bool) {

	api.InitRoutes(gupFunc, vtFunc).Run()

}
