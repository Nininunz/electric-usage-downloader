package main

import (
	"fmt"
	"usage-scraper/internal/app"
)

func main() {
	err := app.Main()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
