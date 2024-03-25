package restutil

import "encoding/base64"

const (
	HeaderAuthorization = "Authorization"
	HeaderContentType   = "Content-Type"
	HeaderAccept        = "Accept"

	Bearer = "bearer"
)

func BasicAuthHeader(username, password string) (string, string) {
	return HeaderAuthorization, "Basic " + BasicAuth(username, password)
}

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func AccessTokenAuthHeader(token string) (string, string) {
	return HeaderAuthorization, Bearer + " " + token
}
