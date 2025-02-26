package middlewares

import (
	"net/http"

	"github.com/solrac97gr/yendoapi/utilities"
)

/*CheckToken : Check the validate of the jwt*/
func CheckToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := utilities.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error in the jwt !"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
