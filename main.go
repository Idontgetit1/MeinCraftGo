package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
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

	var window_width, window_height int = 800, 800
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

func CompileShader(shaderType int, source *string) uint32{
	id := gl.CreateShader(gl.VERTEX_SHADER)
	csource := gl.Str(*source)
	gl.ShaderSource(id, 1, &csource, nil)
	gl.CompileShader(id)
	
	return id
}

func CreateShader(vertexShader *string, fragmentShader *string) uint32 {
	program := gl.CreateProgram()
	vs := CompileShader(gl.VERTEX_SHADER, vertexShader)
	fs := CompileShader(gl.FRAGMENT_SHADER, fragmentShader)

	gl.AttachShader(program, vs)
	gl.AttachShader(program, fs)
	gl.LinkProgram(program)
	gl.ValidateProgram(program)

	return program
}

//main function
func main(){
	runtime.LockOSThread()

	window := initGLFW()
	defer glfw.Terminate()

	program := initOpenGL()

	window.SetKeyCallback(key_callback)
	

	var positions = []float32{
		-0.5, -0.5,
		0.0, 0.5,
		0.5, -0.5,
	}

	var buffer uint32
	gl.GenBuffers(1, &buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(positions), gl.Ptr(positions), gl.STATIC_DRAW)

	
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))
	

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.UseProgram(program)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		window.SwapBuffers()
		glfw.PollEvents()
	}

}