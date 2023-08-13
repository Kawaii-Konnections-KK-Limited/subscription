package main

import (
	"github.com/Kawaii-Konnections-KK-Limited/subscription/api"
)

func main() {

	api.InitRoutes(nil, nil).Run()

}
