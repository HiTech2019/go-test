package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	result := `{"points": A}`

	points := "100.102"
	result = strings.Replace(result, "A", points, -1)

	fmt.Printf("%s\n", result)

	var Points struct {
		Points float64 `json:"points"`
	}

	err := json.Unmarshal([]byte(result), &Points)
	if err != nil {
		fmt.Printf("Unmarshal err:%s", err.Error())
		return
	}

	fmt.Println(Points.Points)

	fmt.Printf("%.4f\n", Points.Points)
	fmt.Printf("%.6f\n", Points.Points)
	fmt.Printf("%.8f\n", Points.Points)
}
