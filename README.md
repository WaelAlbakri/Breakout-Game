# Breakout Game

A classic **Breakout** game where the player controls a paddle to bounce a ball and break all the blocks on the screen. The game was created using **Go** and **Raylib**.

[Watch the gameplay on YouTube](https://www.youtube.com/watch?v=r3oxYARird4)

## Gameplay
- The player’s goal is to clear a grid of blocks using a paddle and a bouncing ball.
- The ball is locked to the paddle at the start of the game. Press **space** to launch the ball.
- The ball's movement direction depends on where the paddle is moving at launch.
- If the ball hits a block, the block will disappear.
- The ball speeds up with each block destroyed.
- The game resets if the player clears all the blocks or the ball falls off the bottom of the screen.

## Features
- Player paddle moves left and right, locked to the screen.
- The ball begins locked to the player paddle and launches with the space bar.
- The ball bounces off the walls, paddle, and blocks.
- A grid of blocks appears with spacing between them.
- The ball destroys blocks upon collision.
- The ball speeds up after hitting blocks.
- Game resets when all blocks are destroyed or the ball goes off the screen.

## Controls
- **Left Arrow**: Move paddle left
- **Right Arrow**: Move paddle right
- **Space Bar**: Launch the ball

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/WaelAlbakri/Breakout-Game.git
2. Ensure you have Raylib installed for Go.
3. Navigate to the project directory
4. Run the game: go run main.go
## Technologies Used
- **Go**: The programming language used for the game logic.
- **Raylib**: A simple and easy-to-use library for game development.
## Possible Improvements
- Adding multiple levels with increasing difficulty.
- Power-ups (such as expanding the paddle or slowing down the ball).
- Adding sound effects.
## Gameplay Video
Watch the gameplay on YouTube: [Breakout Game Play](https://www.youtube.com/watch?v=r3oxYARird4)
