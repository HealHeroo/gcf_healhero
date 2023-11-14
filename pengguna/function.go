package gcf_healhero

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HealHeroo", healHeroo_Pengguna)
}

func healHeroo_Pengguna(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://healheroo.github.io")
		w.Header().Set("Access-Control-Allow-Methods", "GET,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://healheroo.github.io")
	fmt.Fprintf(w, module.GCFHandlerLogin("PASETOPRIVATEKEY", "MONGOSTRING", "healhero_db", r))

}