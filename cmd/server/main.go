package main

import (
	"market-brief-be/internal/router"
)

func main() {
	r := router.New()
	r.Run(":8080")
}
