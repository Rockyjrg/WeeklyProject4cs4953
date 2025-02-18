package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 400, "Flappy Bird - Weekly Project 4")

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

}
