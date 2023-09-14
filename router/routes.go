package router

import (
	"go/postgres-go/middleware"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/gorilla/mux"
)

func HelloMessage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("server up")
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", HelloMessage).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", middleware.GetALLStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/delete/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")
	return router
}
