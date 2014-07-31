package auth

import (
	"encoding/base64"
	"net/http"

	"github.com/codegangsta/negroni"
)

// Basic returns a negroni.HandlerFunc that authenticates via Basic Auth. Writes a http.StatusUnauthorized
// if authentication fails
func Basic(username string, password string) negroni.HandlerFunc {
	var siteAuth = base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	return func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		auth := req.Header.Get("Authorization")
		if !SecureCompare(auth, "Basic "+siteAuth) {
			res.Header().Set("WWW-Authenticate", "Basic realm=\"Authorization Required\"")
			http.Error(res, "Not Authorized", http.StatusUnauthorized)
		}
		r := res.(negroni.ResponseWriter)
		if r.Status() != http.StatusUnauthorized {
			next(res, req)
		}
	}
}
