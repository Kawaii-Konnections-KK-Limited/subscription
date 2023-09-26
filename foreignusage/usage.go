package foreignusage

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Kawaii-Konnections-KK-Limited/subscription/api"
)

func InitService(gupFunc func(token *string) *[]string, vtFunc func(token *string) bool, SubUpdateSender func(token *string), certFile, keyFile *string) {
	port := os.Getenv("PORT")
	addr := os.Getenv("ADDRESS")
	if port == "" {
		fmt.Println(errors.New("please set PORT environment variable"))
		port = "8080"
		fmt.Println("Defaulting to port ", port)

	}
	if addr == "" {
		fmt.Println(errors.New("please set ADDRESS environment variable"))
		addr = "0.0.0.0"
		fmt.Println("Defaulting to addr ", addr)
	}
	if addr != "" && certFile != nil && keyFile != nil {
		log.Println("running with tls")
		api.InitRoutes(gupFunc, vtFunc, SubUpdateSender).RunTLS(fmt.Sprintf("%s:%s", addr, port), *certFile, *keyFile)
	} else {
		log.Println("running without tls")

		api.InitRoutes(gupFunc, vtFunc, SubUpdateSender).Run(fmt.Sprintf("%s:%s", addr, port))
	}

}
