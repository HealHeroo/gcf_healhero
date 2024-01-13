package gcf_healhero

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/HealHeroo/be_healhero/module"
)

func init() {
	functions.HTTP("HealHeroo", healHero_Password)
}

func healHero_Password(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://healhero.my.id")
		w.Header().Set("Access-Control-Allow-Methods", "PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://healhero.my.id")
	fmt.Fprintf(w, module.GCFHandlerUpdatePasswordUser("PASETOPUBLICKEY", "MONGOSTRING", "healhero_db", r))

}