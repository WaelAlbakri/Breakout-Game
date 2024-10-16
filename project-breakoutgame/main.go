package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth   = 800
	screenHeight  = 600
	paddleWidth   = 100
	paddleHeight  = 20
	ballRadius    = 10
	blockRows     = 3     // 3 rows
	blockCols     = 10    // 10 columns
	blockSize     = 60    // Square block width and height
	blockPadding  = 10    // Padding between blocks
	initialSpeed  = 150.0 // Slower initial speed for smoother gameplay
	speedIncrease = 0.02  // Slower speed increase
	paddleSpeed   = 600.0 // Increased paddle speed for better control
)

type Block struct {
	x, y      float32
	width     float32
	height    float32
	destroyed bool
}

// Function to reset the game
func resetGame(paddle *rl.Rectangle, ball *rl.Vector2, ballSpeed *rl.Vector2, blocks *[blockRows][blockCols]Block) {
	paddle.X = screenWidth/2 - paddleWidth/2

	// Reset ball to be locked to the paddle
	ball.X = paddle.X + paddle.Width/2
	ball.Y = paddle.Y - ballRadius
	ballSpeed.X = 0
	ballSpeed.Y = 0

	for row := 0; row < blockRows; row++ {
		for col := 0; col < blockCols; col++ {
			blocks[row][col].destroyed = false
		}
	}
}

func Lerp(a, b, t float32) float32 {
	return a + t*(b-a)
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Breakout")
	rl.SetTargetFPS(60)

	paddle := rl.Rectangle{screenWidth/2 - paddleWidth/2, screenHeight - 40, paddleWidth, paddleHeight}

	ball := rl.Vector2{paddle.X + paddle.Width/2, paddle.Y - ballRadius}
	ballSpeed := rl.Vector2{0, 0}
	ballSpeedMultiplier := 1.0

	var blocks [blockRows][blockCols]Block
	for row := 0; row < blockRows; row++ {
		for col := 0; col < blockCols; col++ {
			blocks[row][col] = Block{
				x:      float32(col)*(blockSize+blockPadding) + (screenWidth-(blockCols*(blockSize+blockPadding)))/2, // Centering the blocks horizontally
				y:      float32(row)*(blockSize+blockPadding) + 50,                                                   // Vertical spacing
				width:  blockSize,
				height: blockSize,
			}
		}
	}

	gameActive := false

	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyRight) && paddle.X+paddle.Width < screenWidth {
			paddle.X += paddleSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyLeft) && paddle.X > 0 {
			paddle.X -= paddleSpeed * rl.GetFrameTime()
		}

		if rl.IsKeyPressed(rl.KeySpace) && ballSpeed.X == 0 && ballSpeed.Y == 0 {
			// Randomize the horizontal launch direction but ensure the vertical speed is always upward
			ballSpeed.X = initialSpeed * float32(rl.GetRandomValue(-100, 100)) / 100.0
			ballSpeed.Y = -initialSpeed
			gameActive = true
		}

		// Restart game with "R"
		if rl.IsKeyPressed(rl.KeyR) {
			resetGame(&paddle, &ball, &ballSpeed, &blocks)
			gameActive = false
		}

		if gameActive {
			ball.X += ballSpeed.X * rl.GetFrameTime() * float32(ballSpeedMultiplier)
			ball.Y += ballSpeed.Y * rl.GetFrameTime() * float32(ballSpeedMultiplier)
		}
		if ball.X-ballRadius <= 0 {
			ballSpeed.X *= -1
			ball.X = ballRadius
		} else if ball.X+ballRadius >= screenWidth {
			ballSpeed.X *= -1
			ball.X = screenWidth - ballRadius
		}

		if ball.Y-ballRadius <= 0 {
			ballSpeed.Y *= -1
			ball.Y = ballRadius
		}

		if rl.CheckCollisionCircleRec(ball, ballRadius, paddle) {
			ballSpeed.Y *= -1

			// Calculate the position where the ball hits the paddle
			hitPosition := (ball.X - (paddle.X + paddle.Width/2)) / (paddle.Width / 2)

			ballSpeed.X = Lerp(-initialSpeed, initialSpeed, (hitPosition+1)/2)
		}

		for row := 0; row < blockRows; row++ {
			for col := 0; col < blockCols; col++ {
				block := &blocks[row][col]
				if !block.destroyed && rl.CheckCollisionCircleRec(ball, ballRadius, rl.NewRectangle(block.x, block.y, block.width, block.height)) {
					block.destroyed = true
					ballSpeed.Y *= -1
					ballSpeedMultiplier += speedIncrease
				}
			}
		}

		if ball.Y+ballRadius >= screenHeight {
			resetGame(&paddle, &ball, &ballSpeed, &blocks)
			gameActive = false
		}

		// Check if all blocks are destroyed
		allDestroyed := true
		for row := 0; row < blockRows; row++ {
			for col := 0; col < blockCols; col++ {
				if !blocks[row][col].destroyed {
					allDestroyed = false
					break
				}
			}
		}
		if allDestroyed {
			resetGame(&paddle, &ball, &ballSpeed, &blocks)
			gameActive = false // Game restarts and i press space
		}

		// Drawing
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawRectangleRec(paddle, rl.Blue)

		rl.DrawCircleV(ball, ballRadius, rl.Red)

		// Draw blocks
		for row := 0; row < blockRows; row++ {
			for col := 0; col < blockCols; col++ {
				if !blocks[row][col].destroyed {
					rl.DrawRectangle(int32(blocks[row][col].x), int32(blocks[row][col].y), int32(blocks[row][col].width), int32(blocks[row][col].height), rl.Green)
				}
			}
		}

		rl.DrawText("Press SPACE to launch the ball", 10, 10, 16, rl.Black)
		rl.DrawText("Press Left/Right arrows to move the paddle", 10, 30, 16, rl.Black)
		rl.DrawText("Press R to restart the game", 10, 50, 16, rl.Black)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
