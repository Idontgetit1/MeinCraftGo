package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)


// initOpenGL initializes OpenGL -> returns program id
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
}

// create a window and initialize glfw -> returns window
func initGLFW() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	var window_width, window_height int = 1270, 720
	var window_name string = "MeinCraftGo"

	//create window on secondary monitor
	window, err := glfw.CreateWindow(window_width, window_height, window_name, nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}


//draw function
func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.UseProgram(program)
	//make screen red
	gl.ClearColor(250.0/255.0, 119.0/255.0, 110.0/255.0, 1.0)

	//swap buffers
	window.SwapBuffers()
	glfw.PollEvents()
}


// key_callback function
// get key inputs
func key_callback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyE && action == glfw.Press {
		//send message to console
		log.Println("E pressed")
	}
}

//main function
func main(){
	runtime.LockOSThread()

	window := initGLFW()
	defer glfw.Terminate()

	program := initOpenGL()

	window.SetKeyCallback(key_callback)

	
	for !window.ShouldClose() {

		draw(window, program)
	}

}