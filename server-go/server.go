package server

import (
	"go_pill/headers-go"
	"go_pill/html-go"
	"go_pill/json-go"
	"go_pill/plaintext-go"
	"go_pill/xml-go"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ServerCode() {
	//init router
	r := mux.NewRouter()

	//router handlers/endpoints
	r.HandleFunc("/header", headers.HttpSetHeader).Methods("GET")

	//header request to check status. It is shown by requesting to a ran
	//-dom http address
	//headers.HttpHeaderRequest()  //use it to test status request, the first part of the pill

	//returning an html object
	r.HandleFunc("/html", html.Gethtml).Methods("GET")
	//returning  xml object
	r.HandleFunc("/xml", xml.GetXmlValue).Methods("GET")
	//returning json object
	r.HandleFunc("/json", json.GetJsonObject).Methods("GET")
	//returning a plain text
	r.HandleFunc("/plaintext", plaintext.ImportPlainText).Methods("GET")
	//returning an image
	log.Fatal(http.ListenAndServe(":8000", r))

}
