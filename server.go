package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	router := mux.NewRouter()

	router.Handle("/api/greet", Greet()).Methods("GET")

	fmt.Printf("Available routes:\n")
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		m, err := route.GetMethods()
		if err != nil {
			return err
		}
		fmt.Printf("%s\t%s\n", m, t)
		return nil
	})
	a.Router = router
}

func Greet() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		JSONResponse(w, http.StatusOK, map[string]string{"response": "HELLO YOU FELLOW TRAVELER!"})
	})
}

func (a *App) Run(addr string) {
	handler := a.Router
	log.Printf("Server is listening on %v\n", addr)
	http.ListenAndServe(addr, handler)
}

func JSONResponse(w http.ResponseWriter, code int, output interface{}) {
	response, _ := json.Marshal(output)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
