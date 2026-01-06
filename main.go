package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// Record represents the win-loss record between two entities
type Record struct {
	W int `json:"W"`
	L int `json:"L"`
}

// H2H represents head-to-head records between two teams
// It is a nested map where the first key is the name of the primary team,
// the second key is the name of its opponent, and the value is the Record struct
type H2H map[string]map[string]Record

// Usage: go run main.go example.json
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <example.json>")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var h2h H2H
	err = json.Unmarshal(data, &h2h)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Printf("%#v\n", h2h)

	teams := getTeams(h2h)
	printMatrix(h2h, teams)
}

func getTeams(h2h H2H) []string {
	set := make(map[string]struct{})

	// Make sure all teams are collected,
	for team, opponents := range h2h {
		set[team] = struct{}{}
		for opponent := range opponents {
			set[opponent] = struct{}{}
		}
	}

	// Convert the set into a sorted slice
	teams := make([]string, 0, len(set))
	for team := range set {
		teams = append(teams, team)
	}
	sort.Strings(teams)

	return teams
}

// printMatrix prints the head-to-head matrix in a formatted way
// The one and only important part is to determine the width of each column
func printMatrix(h2h H2H, teams []string) {
	// Header
	fmt.Printf("%-5s", "Tm")
	for _, t := range teams {
		fmt.Printf("%8s", t)
	}
	fmt.Println()

	// Separator
	fmt.Printf("%-5s", "-----")
	for range teams {
		fmt.Printf("%8s", "--------")
	}
	fmt.Println()

	// Rows
	for _, rowTeam := range teams {
		fmt.Printf("%-5s", rowTeam)

		for _, colTeam := range teams {
			if rowTeam == colTeam {
				fmt.Printf("%8s", "--")
				continue
			}

			if rec, ok := h2h[rowTeam][colTeam]; ok {
				fmt.Printf("%8d", rec.W)
			} else {
				fmt.Printf("%8s", "")
			}
		}
		fmt.Println()
	}

	fmt.Printf("%-5s", "-----")
	for range teams {
		fmt.Printf("%8s", "--------")
	}
	fmt.Println()

	fmt.Printf("%-5s", "Tm")
	for _, t := range teams {
		fmt.Printf("%8s", t)
	}
	fmt.Println()
}
