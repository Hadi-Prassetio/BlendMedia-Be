package routes

import (
	"ibox/handlers"
	"ibox/pkg/middleware"
	"ibox/pkg/mysql"
	"ibox/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/carts", middleware.Auth(h.FindCarts)).Methods("GET")
	r.HandleFunc("/cart-id", middleware.Auth(h.GetCart)).Methods("GET")
	r.HandleFunc("/cart", middleware.Auth(h.CreateCart)).Methods("POST")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.DeleteCart)).Methods("DELETE")
	r.HandleFunc("/cartID", middleware.Auth(h.UpdateCart)).Methods("PATCH")
	r.HandleFunc("/cart-status", middleware.Auth(h.FindbyIDCart)).Methods("GET")
}
