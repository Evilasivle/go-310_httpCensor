package main

import "igaleksus.com/spamMasker/internal/masker"

func main() {
	service := masker.NewService()
	service.Run()
}
