// this is complete
package main

import (
	"fmt"
	"net/http"

	"github.com/pluralsight/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
	fmt.Println("this works")
}
