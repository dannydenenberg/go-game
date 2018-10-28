package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	basicEnemyWidth  = 840
	basicEnemyHeight = 1034
)

type basicEnemy struct {
	tex  *sdl.Texture
	X, Y float64
}

func newBasicEnemy(renderer *sdl.Renderer, x, y float64) (be basicEnemy, err error) {
	// IMAGE SURFACE
	img, err := sdl.LoadBMP("basic_enemy.bmp") // 840 × 1034
	if err != nil {
		return basicEnemy{}, fmt.Errorf("loading bsaic enemy sprite: %v", err)
	}
	defer img.Free() // after image is placed into texture, you can free the memory

	be.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return basicEnemy{}, fmt.Errorf("creating basic enemy texture: %v", err)
	}

	be.X = screenWidth / 2
	be.Y = screenHeight - playerHeight/2.0

	return be, nil
}

func (be *basicEnemy) draw(renderer *sdl.Renderer) {
	// converting basic enemy coordinates to top left of sprite
	x := be.X - basicEnemyWidth/2.0
	y := be.Y - basicEnemyHeight/2.0

	// image stuff --> src is how much of the image you are referencing you want, and destination is where on the screen
	// you want to fill the image with
	renderer.Copy(be.tex,
		&sdl.Rect{X: 0, Y: 0, W: basicEnemyWidth, H: basicEnemyHeight},
		&sdl.Rect{X: int32(x), Y: int32(y), W: basicEnemyWidth / 5, H: basicEnemyHeight / 5})

}
