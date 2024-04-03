package main

import (
	"net/http"

	"github.com/fermyon/enterprise-architectures-and-patterns/http-crud-go-sqlite/pkg/api"
	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		router := api.New()
		router.ServeHTTP(w, r)
	})
}

func main() {}
