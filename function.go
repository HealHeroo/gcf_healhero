package HealHero

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	be_healhero "github.com/HealHeroo/be_healhero"
)

func init() {
	functions.HTTP("HealHeroo", healHerooPost)
}

func healHerooPost(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://healheroo.github.io")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://healheroo.github.io")
	fmt.Fprintf(w, be_healhero.GCFPostHandler("PASETOPRIVATEKEY", "MONGOSTRING", "healhero_db", "user", r))

}