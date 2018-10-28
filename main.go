package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
	spaceship    = "spaceship.bmp"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL: ", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Gaming in go",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)

	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}

	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}

	defer renderer.Destroy()

	plr, err := newPlayer(renderer)
	if err != nil {
		fmt.Println("creating player: ", err)
		return
	}

	enemy, err := newBasicEnemy(renderer, screenWidth/2.0, screenHeight/2.0)
	if err != nil {
		fmt.Println("creating basic enemy:", err)
		return
	}
	// add an infinite loop just to see the screen up forever
	for {
		// consume every event that sdl has detected in its QUEUE
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			//switch on the type of the event object
			switch event.(type) {
			// if the user quits, it closes
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 0, 0, 255)

		// fills screen with last draw color
		renderer.Clear()

		plr.draw(renderer)
		plr.update()

		enemy.draw(renderer)

		// places everything on the screen
		renderer.Present()
	}

}
