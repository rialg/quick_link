package utils

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Read the list of quick links and return the mapping
func GetUris() map[string]string {

	content, err := os.ReadFile("./static/redirects.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var list map[string]string
	err = json.Unmarshal(content, &list)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return list

}

// Update the list of quick links
func AddUri(new_uri string, linked_url string) {

	log.Println(new_uri, " -> ", linked_url)

	content, err := os.ReadFile("./static/redirects.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var list map[string]string
	err = json.Unmarshal(content, &list)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	list[new_uri] = linked_url
	new_json, _ := json.Marshal(list)

	// TODO: Change file for a proper DB where to
	// store the quick links
	os.Truncate("./static/redirects.json", 0)
	os.WriteFile("./static/redirects.json", new_json, os.ModeAppend)

}

// Handle adding a new URI
func AddUriTemplate(w http.ResponseWriter, new_uri string) {

	tpl, err := template.ParseFiles("./static/add_template.html")
	if err != nil {

		log.Default().Print("Error while parsing add_template.html")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	data := struct{ URI string }{URI: new_uri}
	err = tpl.Execute(w, data)

	if err != nil {

		log.Default().Print("Error while adding URI to template")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

}
