package main

import (
	"net/http"
)

func main() {
	models.InitDB("postgres://pmutua@localhost/health_insurance")

	http.ListenAndServe(":3000", nil)
}
