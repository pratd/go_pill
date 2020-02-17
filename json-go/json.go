package json

import (
	"encoding/json"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func GetJsonObject(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Pratt", []string{"golang", "python"}}

	js, err := json.MarshalIndent(profile, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
