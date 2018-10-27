package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed  = 0.2
	playerWidth  = 95
	playerHeight = 180
)

type player struct {
	tex  *sdl.Texture
	X, Y float64
}

/*
 * Creates a new player
 * By using named return values, it is as if you have already initialized variables and can use them throughout the function
 */
func newPlayer(renderer *sdl.Renderer) (p player, err error) {
	// IMAGE SURFACE
	img, err := sdl.LoadBMP("rocketship1.bmp")
	if err != nil {
		return player{}, fmt.Errorf("loading player sprite: %v", err)
	}
	defer img.Free() // after image is placed into texture, you can free the memory

	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("creating player texture: %v", err)
	}

	p.X = screenWidth / 2
	p.Y = screenHeight - playerHeight/2.0

	return p, nil
}

/*
 * Draws the player image to the screen
 */
func (p *player) draw(renderer *sdl.Renderer) {

	x := p.X - playerWidth/2.0
	y := p.Y - playerHeight/2.0

	// image stuff --> src is how much of the image you are referencing you want, and destination is where on the screen
	// you want to fill the image with
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 190, H: 360},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 95, H: 180})

}

/*
 * Update the player status:
 * 		- Keyboard events
 */
func (p *player) update() {
	keys := sdl.GetKeyboardState() // returns []uint8 => state of every key on keyboard

	// 0 is not pressed, 1 = is pressed
	if keys[sdl.SCANCODE_LEFT] == 1 && p.X > 0+playerWidth/2.0 {
		// move player left
		p.X -= playerSpeed

	} else if keys[sdl.SCANCODE_RIGHT] == 1 && p.X < screenWidth-playerWidth/2.0 {
		// move player right
		p.X += playerSpeed

	}
}
