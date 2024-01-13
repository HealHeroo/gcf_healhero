package gcf_healhero

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/HealHeroo/be_healhero/module"
)

func init() {
	functions.HTTP("HealHeroo", healHero_Pengguna)
}

func healHero_Pengguna(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://healhero.my.id")
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method == http.MethodPut {
		fmt.Fprintf(w, module.GCFHandlerUpdatePengguna("PASETOPUBLICKEY", "MONGOSTRING", "healhero_db", r))
		return
	}
	if r.Method == http.MethodDelete {
		fmt.Fprintf(w, module.GCFHandlerDeletePengguna("PASETOPUBLICKEY", "MONGOSTRING", "healhero_db", r))
		return
	}
	// Set CORS headers for the main request.
	fmt.Fprintf(w, module.GCFHandlerGetPengguna("PASETOPUBLICKEY", "MONGOSTRING", "healhero_db", r))
}