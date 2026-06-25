package auth

import "net/http"

const HeaderName = "X-Access-Token"

func VerifyToken(r *http.Request, expectedToken string, enabled bool) bool {
	if !enabled {
		return true
	}
	return r.Header.Get(HeaderName) == expectedToken
}
