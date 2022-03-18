// Package main
package main

import (
	"github.com/joho/godotenv"
	"github.com/orn-id/wedding-api/src/cmd"
)

func main() {
	godotenv.Load()
	cmd.Start()
}
