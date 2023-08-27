package foreignusage

import (
	"log"

	"github.com/Kawaii-Konnections-KK-Limited/subscription/api"
)

func InitService(gupFunc func(token *string) *[]string, vtFunc func(token *string) bool, addr, certFile, keyFile *string) {

	if addr != nil && certFile != nil && keyFile != nil {
		log.Println("running with tls")
		api.InitRoutes(gupFunc, vtFunc).RunTLS(*addr, *certFile, *keyFile)
	} else {
		log.Println("running without tls")

		api.InitRoutes(gupFunc, vtFunc).Run()
	}

}
