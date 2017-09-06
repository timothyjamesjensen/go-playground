// Example use: go run stash-prevent-changes-wo-pr.go username password PIPE

package main

import (
	"fmt"
	"net/http"
	"os"
	"crypto/tls"
	"encoding/json"
	"bytes"
)

type Repository struct {
	Values []Project `json:"values"`
}

type Project struct {
	Slug string `json:"slug"`
}

func main() {

	tr := &http.Transport {
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}

	r := new (Repository)

	baseUrl := "https://stash.americas.nwea.pvt/rest"

	username := os.Args[1]
	password := os.Args[2]
	project  := os.Args[3]

	client := &http.Client{Transport: tr}

	req, _ := http.NewRequest("GET", baseUrl + "/api/1.0/projects/" + project + "/repos?limit=1000", nil)
	req.SetBasicAuth(username, password)

	res, err := client.Do(req);
	if err != nil {
		panic(err)
	}

	json.NewDecoder(res.Body).Decode(&r)

	config := []byte(`{
		"id": "1",
		"type": "pull-request-only",
		"matcher": {
			"id": "refs/heads/master",
			"displayId": "master",
			"type": {
				"id": "BRANCH",
				"name": "Branch"
			},
			"active": "true"
		}
	}`)

	for i:= 0; i < len(r.Values); i++ {
		post, _ := http.NewRequest("POST", baseUrl + "/branch-permissions/2.0/projects/" + project + "/repos/" + r.Values[i].Slug  + "/restrictions", bytes.NewBuffer(config))
		post.SetBasicAuth(username, password)
		post.Header.Set("Content-Type", "application/json")

		res, err := client.Do(post)
		if err != nil {
			panic(err);
		}

		fmt.Println(res)
		fmt.Println(r.Values[i])
	}
}
