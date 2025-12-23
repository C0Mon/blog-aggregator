package main

import (
	"fmt"

	"github.com/C0Mon/blog-aggregator/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.SetUser("Colm")
	cfg = config.Read()
	fmt.Printf("%+v\n", cfg)
}
