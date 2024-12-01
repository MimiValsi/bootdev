package main

import (
	"errors"
	"strings"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	users := make(map[string]user)
	
	for i, name := range names {
		users[name] = user{
			name: name,
			phoneNumber: phoneNumbers[i],
			
		}
	}

	return users, nil
}

type user struct {
	name        string
	phoneNumber int
}

func getNameCounts(names []string) map[rune]map[string]int {
	parent := make(map[rune]map[string]int)
	for _, name := range names {
		runes := []rune(name)
		r := runes[0]
		if _, ok := parent[r]; !ok {
			parent[r] = make(map[string]int)
		}
		parent[r][name]++
		
		
	}
	return parent
}

func countDistinctWords(messages []string) int {
	unique := make(map[string]int)
	for _, msg := range messages {
		words := strings.Fields(strings.ToLower(msg))
		for _, word := range words {
			if _, ok := unique[word]; !ok {
				unique[word] = 0
			}
			unique[word]++
		}
	}

	return len(unique)
}
