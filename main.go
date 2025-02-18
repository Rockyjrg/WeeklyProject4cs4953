package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	//player variables
	var playerX float32 = 100
	var playerY float32 = 100
	var playerSize float32 = 40
	var playerSpeed float32 = 100

	//pipe variables
	var pipeX float32 = 700
	var topPipeY float32 = 0
	var bottomPipeY float32 = 300
	var pipeHeight float32 = 100
	var pipeWidth float32 = 40

	//point counter variables
	//var pointCounter int = 0

	//size of the window
	rl.InitWindow(800, 400, "Flappy Bird - Weekly Project 4")

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Current Points: 0", 0, 0, 10, rl.Black)

		//player rectangle
		rl.DrawRectangle(int32(playerX), int32(playerY), int32(playerSize), int32(playerSize), rl.Blue)

		//bottom pipe rectangle
		rl.DrawRectangle(int32(pipeX), int32(bottomPipeY), int32(pipeWidth), int32(pipeHeight), rl.Green)
		//top pipe rectangle
		rl.DrawRectangle(int32(pipeX), int32(topPipeY), int32(pipeWidth), int32(pipeHeight), rl.Green)

		//character movement + blocking going outside window
		if rl.IsKeyDown(rl.KeyW) && playerY > 0 {
			playerY -= playerSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyS) && playerY < 400-playerSize {
			playerY += playerSpeed * rl.GetFrameTime()
		}

		rl.EndDrawing()
	}
}
