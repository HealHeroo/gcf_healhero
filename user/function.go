package gcf_healhero

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/HealHeroo/be_healhero/module"
)

func init() {
	functions.HTTP("HealHeroo", healHero_User)
}

func healHero_User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://healhero.my.id")
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	
	// Set CORS headers for the main request.
	fmt.Fprintf(w, module.GCFHandlerGetUser("PASETOPUBLICKEY", "MONGOSTRING", "healhero_db", r))
}