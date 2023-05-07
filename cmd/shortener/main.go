package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/go-chi/chi"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	// r.Get("/value/*", handlers.GetMetric)
	// r.Get("/", handlers.ShowMetrics)
	// r.Post("/", NewHandStruct.Update)
	// r.Get("/{id}", NewHandStruct.NowValueMetrics)
	return r
}
func main() {
	genString := fmt.Sprint(rand.Int63n(1000))
	fmt.Println(genString)
	r := chi.NewRouter()
	http.ListenAndServe(":8080", r)
}
