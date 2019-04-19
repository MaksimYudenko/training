package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	root = "/v1/training"
)

func main() {
	// get configuration
	address := flag.String("server", "http://localhost:8080",
		"HTTP gateway url, e.g. http://localhost:8080")
	flag.Parse()
	var body string
	// Create
	resp, err := http.Post(*address+root, "application/json",
		strings.NewReader(fmt.Sprintf(`
		{
			"api":"v1",
			"trainee":
				{
					"name":"goName",
					"surname":"goSurname",
					"email":"goEmail",
					"hasAttend":false,
					"score":50
				}
		}`)))
	if err != nil {
		log.Fatalf("failed to call Create method: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Create response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Create response: Code=%d, Body=%s\n", resp.StatusCode, body)
	// parse ID of created trainee
	var created struct {
		API string `json:"api"`
		ID  string `json:"id"`
	}
	err = json.Unmarshal(bodyBytes, &created)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON response of Create method: %v\n", err)
	}
	// Read
	resp, err = http.Get(fmt.Sprintf("%s%s/%s", *address, root, created.ID))
	if err != nil {
		log.Fatalf("failed to call Read method: %v", err)
	}
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Read response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Read response: Code=%d, Body=%s\n", resp.StatusCode, body)
	// Update
	req, err := http.NewRequest("PUT",
		fmt.Sprintf("%s%s/%s", *address, root, created.ID),
		strings.NewReader(fmt.Sprintf(`
		{
			"api":"v1",
			"trainee": 
				{
					"name":"goName + updated",
					"surname":"goSurname + updated",
					"email":"goEmail + updated",
					"hasAttend": true,
					"score":100
				}
		}`)))
	req.Header.Set("Content-Type", "application/json")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to call Update method: %v", err)
	}
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Update response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Update response: Code=%d, Body=%s\n", resp.StatusCode, body)
	// ReadAll
	resp, err = http.Get(*address + root)
	if err != nil {
		log.Fatalf("failed to call ReadAll method: %v", err)
	}
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read ReadAll response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("ReadAll response: Code=%d, Body=%s\n", resp.StatusCode, body)
	// Delete
	req, err = http.NewRequest("DELETE",
		fmt.Sprintf("%s%s/%s", *address, root, created.ID), nil)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to call Delete method: %v", err)
	}
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Delete response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Delete response: Code=%d, Body=%s\n", resp.StatusCode, body)
}
