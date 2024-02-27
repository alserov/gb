package main

import (
	"github.com/alserov/gb/sm_2_4/internal/app"
	"github.com/alserov/gb/sm_2_4/internal/config"
)

func main() {
	cfg := config.MustLoad()
	app.MustStart(cfg)
}
