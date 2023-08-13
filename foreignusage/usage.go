package foreignusage

import "github.com/Kawaii-Konnections-KK-Limited/subscription/api"

func InitService() {

	api.InitRoutes(nil, nil).Run()

}
