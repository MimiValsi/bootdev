package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
)

type Issue struct {
	Title    string `json:"title"`
	Estimate int    `json:"estimate"`
}

func getIssuess(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	issues := []Issue{}
	if err = json.Unmarshal(data, &issues); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return issues, nil
}

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	issues := []Issue{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&issues); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return issues, nil
}

func getResources(url string) ([]map[string]any, error) {
	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&resources); err != nil {
		return resources, err
	}
	return resources, nil
}

func logResources(resources []map[string]any) {
	var formattedStrings []string

	for _, resource := range resources {
		for k, v := range resource {
			str := fmt.Sprintf("Key: %s - Value: %v", k, v)
			formattedStrings = append(formattedStrings, str)
		}
	}

	sort.Strings(formattedStrings)

	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}
