package main

import (
	"fmt"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	//player variables
	var playerX float32 = 100
	var playerY float32 = 100
	var playerSize float32 = 40
	var playerSpeed float32 = 170

	//pipe variables
	var pipeX float32 = 700
	var pipeWidth float32 = 40
	var pipeSpeed float32 = 300
	var gapHeight float32 = 70 //space between top and bottom pipes

	//initial pipe heights
	var topPipeHeight float32 = float32(rand.IntN(200) + 50)
	var bottomPipeY float32 = topPipeHeight + gapHeight
	var bottomPipeHeight float32 = 400 - bottomPipeY

	//score variables
	var score int = 0
	var passedPipe bool = false

	//game state variable
	var gameOver bool = false

	//size of the window
	rl.InitWindow(800, 400, "Flappy Bird - Weekly Project 4")

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Points: "+fmt.Sprintf("%d", score), 10, 10, 25, rl.Black)

		if gameOver {
			rl.DrawText("Game over! Press R to Restart", 250, 150, 20, rl.Red)

			if rl.IsKeyPressed(rl.KeyR) {
				//reset game
				playerY = 100
				pipeX = 700
				score = 0
				gameOver = false
				passedPipe = false

				//reset random pipe heights
				topPipeHeight = float32(rand.IntN(200) + 50)
				bottomPipeY = topPipeHeight + gapHeight
				bottomPipeHeight = 400 - bottomPipeY
			}
		} else {

			if rl.IsKeyPressed(rl.KeyR) {
				//reset game
				playerY = 100
				pipeX = 700
				score = 0
				gameOver = false
				passedPipe = false

				//reset random pipe heights
				topPipeHeight = float32(rand.IntN(200) + 50)
				bottomPipeY = topPipeHeight + gapHeight
				bottomPipeHeight = 400 - bottomPipeY
			}

			//character movement + blocking going outside window
			if rl.IsKeyDown(rl.KeyW) && playerY > 0 {
				playerY -= playerSpeed * rl.GetFrameTime()
			}
			if rl.IsKeyDown(rl.KeyS) && playerY < 400-playerSize {
				playerY += playerSpeed * rl.GetFrameTime()
			}

			//player rectangle
			rl.DrawRectangle(int32(playerX), int32(playerY), int32(playerSize), int32(playerSize), rl.Blue)

			//top pipe rectangle
			rl.DrawRectangle(int32(pipeX), 0, int32(pipeWidth), int32(topPipeHeight), rl.Green)
			//bottom pipe rectangle
			rl.DrawRectangle(int32(pipeX), int32(bottomPipeY), int32(pipeWidth), int32(bottomPipeHeight), rl.Green)

			//if pipe goes past left border of the screen respawn on the right
			pipeX -= pipeSpeed * rl.GetFrameTime()

			//set up a way to randomize the heights of each pipe on each turn
			if pipeX < -pipeWidth {
				pipeX = 800
				topPipeHeight = float32(rand.IntN(200) + 50)
				bottomPipeY = topPipeHeight + gapHeight
				bottomPipeHeight = 400 - bottomPipeY
				passedPipe = false
			}

			//collision detection
			playerRect := rl.NewRectangle(playerX, playerY, playerSize, playerSize)
			topPipeRect := rl.NewRectangle(pipeX, 0, pipeWidth, topPipeHeight)
			bottomPipeRect := rl.NewRectangle(pipeX, bottomPipeY, pipeWidth, bottomPipeHeight)

			if rl.CheckCollisionRecs(playerRect, topPipeRect) || rl.CheckCollisionRecs(playerRect, bottomPipeRect) {
				gameOver = true
			}

			if !passedPipe && pipeX+pipeWidth < playerX {
				score++
				passedPipe = true
			}
		}

		rl.EndDrawing()
	}
}
