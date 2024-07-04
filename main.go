package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp-social-network/src/router"
	"webapp-social-network/src/utils"
)

func main() {
	utils.LoadingTemplates()
	r := router.Generate()

	fmt.Println("Listening WebAPP!")
	log.Fatal(http.ListenAndServe(":3000", r))

}
