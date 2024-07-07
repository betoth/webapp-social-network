package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp-social-network/src/config"
	"webapp-social-network/src/cookie"
	"webapp-social-network/src/router"
	"webapp-social-network/src/utils"
)

func main() {
	config.Load()
	cookie.Load()
	utils.LoadingTemplates()
	r := router.Generate()

	fmt.Println("Listening WebAPP in port", config.APIPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.APIPort), r))

}
