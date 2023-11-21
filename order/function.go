package gcf_healhero

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/HealHeroo/be_healhero/module"
)

func init() {
	functions.HTTP("HealHeroo", healHero_Order)
}

func healHero_Order(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://healheroo.github.io")
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, module.GCFHandlerInsertObat("PASETOPUBLICKEY", "MONGOSTRING", "healhero_db", r))
		return
	}
	if r.Method == http.MethodDelete {
		fmt.Fprintf(w, module.GCFHandlerDeleteOrder("PASETOPUBLICKEY", "MONGOSTRING", "healhero_db", r))
		return
	}
	// Set CORS headers for the main request.
	fmt.Fprintf(w, module.GCFHandlerGetOrder("PASETOPUBLICKEY", "MONGOSTRING", "healhero_db", r))
}