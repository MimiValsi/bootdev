package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Role       string `json:"role"`
	ID         string `json:"id"`
	Experience int    `json:"experience"`
	Remote     bool   `json:"remote"`
	User       struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Age      int    `json:"age"`
	} `json:"user"`
}


func createUser(url, apiKey string, data User) (User, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	user := User{}
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&user); err != nil {
		return User{}, err
	}

	return user, nil
	
}

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id

	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}

	user := User{}
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&user); err != nil {
		return User{}, err
	}

	return user, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, err
	}
	req.Header.Set("X-API-Key", apiKey)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	user := User{}
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&user); err != nil {
		return User{}, err
	}
	return user, nil
}

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("failed to delete user with id: %s", id)
	}

	return nil
}

func fetchData(url string) (int, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("network error: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		return res.StatusCode, fmt.Errorf("non-OK HTTP status: %s", res.Status)
	}

	return res.StatusCode, nil
}
