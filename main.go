// author: Jacky Boen

package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "TurtleGo"
var winWidth, winHeight int32 = 640, 480

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer

	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	running := true

	image := map[sdl.Point]bool{}
	cursor := sdl.Point{winWidth / 2, winHeight / 2}
	is_drawing := false

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if t.State == sdl.PRESSED {
					switch t.Keysym.Sym {
					case sdl.K_LEFT:
						cursor.X--
					case sdl.K_RIGHT:
						cursor.X++
					case sdl.K_UP:
						cursor.Y--
					case sdl.K_DOWN:
						cursor.Y++
					case sdl.K_SPACE:
						is_drawing = !is_drawing
					case sdl.K_q:
						if t.Keysym.Mod == sdl.KMOD_LCTRL || t.Keysym.Mod == sdl.KMOD_RCTRL {
							running = false
						}
					}
					if is_drawing {
						image[cursor] = true
					}
				}
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()
		for k, v := range image {
			if v {
				renderer.SetDrawColor(0, 0, 0, 255)
				renderer.DrawPoint(k.X, k.Y)
			}
		}
		if is_drawing {
			renderer.SetDrawColor(255, 0, 0, 255)
		} else {
			renderer.SetDrawColor(0, 255, 0, 255)
		}
		renderer.DrawPoint(cursor.X, cursor.Y)

		renderer.Present()
		sdl.Delay(16)
	}

	return 0
}

func main() {
	os.Exit(run())
}
