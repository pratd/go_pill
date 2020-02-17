package xml

import (
	"encoding/xml"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string `xml:"Hobbies>hobby"`
}

func GetXmlValue(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Pratt", []string{"golang", "python"}}

	x, err := xml.MarshalIndent(profile, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}
