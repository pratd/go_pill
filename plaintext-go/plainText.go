package plaintext

import (
	"net/http"
)

func ImportPlainText(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}
