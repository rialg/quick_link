package main

import (
	"net/http"
	"strings"

	"simple_forwarding_proxy/utils"
)

const PORT string = "<PROXYPORT>"

// Handle redirects
func redirect_handler(w http.ResponseWriter, r *http.Request) {

	url, foundUrl := utils.GetUris()[r.RequestURI]
	new_uri_container := strings.Split(r.RequestURI, "/add")

	if len(new_uri_container) > 1 {

		if r.Method == "POST" {

			r.ParseForm()
			body := r.PostForm
			new_url := body["url"][0]
			utils.AddUri(new_uri_container[1], new_url)
			http.Redirect(w, r, new_url, http.StatusTemporaryRedirect)

		}

		utils.AddUriTemplate(w, new_uri_container[1])

	} else if foundUrl {

		http.Redirect(w, r, url, http.StatusTemporaryRedirect)

	} else {

		http.Redirect(w, r, "http://localhost:"+PORT+"/add"+r.RequestURI, http.StatusTemporaryRedirect)

	}

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", redirect_handler)

	http.ListenAndServe(":"+PORT, mux)
}
