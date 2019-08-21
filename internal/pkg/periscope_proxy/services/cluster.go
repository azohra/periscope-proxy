package services

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// Negotiate negotiate with cluster
func Negotiate(header string, endpoint string) {
	apiURL := fmt.Sprintf("%s/allocate", endpoint)

	template := `{
		"stateID": "%s",
		"svcPort": 7950,
		"images": [
			{ "image":"gcr.io/williams-playground/bigwill/periscope_mockserver_opensource", "port":7950 },
			{ "image":"gcr.io/williams-playground/bigwill/periscope_mockdb_opensource", "port":7990 }
		]
	}`

	formData := url.Values{
		"payload": {fmt.Sprintf(template, header)},
	}
	resp, err := http.PostForm(apiURL, formData)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
		log.Fatalln(resp)
	}
	fmt.Println(resp)
}
