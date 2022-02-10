package main

import (
  "log"
	"net/http"

	"github.com/gowebexample/goreddit/postgres"
  "github.com/gowebexample/goreddit/web"
)

func main() {
  store, err := postgres.NewStore("postgres://postgres:postgres@localhost/postgres?sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }

  h := web.NewHandler(store)
  http.ListenAndServe(":3000", h)
}
