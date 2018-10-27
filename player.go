package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	tex *sdl.Texture
}

/*
 * Creates a new player
 * By using named return values, it is as if you have already initialized variables and can use them throughout the function
 */
func newPlayer(renderer *sdl.Renderer) (p player, err error) {
	// IMAGE SURFACE
	img, err := sdl.LoadBMP("spaceship.bmp")
	if err != nil {
		return player{}, fmt.Errorf("loading player sprite: %v", err)
	}
	defer img.Free() // after image is placed into texture, you can free the memory

	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("creating player texture: %v", err)
	}

	return p, nil
}

/*
 * Draws the player image to the screen
 */
func (p *player) draw(renderer *sdl.Renderer) {
	// image stuff --> src is how much of the image you are referencing you want, and destination is where on the screen
	// you want to fill the image with
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 4000, H: 3000},
		&sdl.Rect{X: 40, Y: 20, W: 240, H: 180})

}
