package main

// https://gist.github.com/mbertschler/8672365
// http://stackoverflow.com/questions/21412844/why-doesnt-this-opengl-program-draw-a-triangle

import (

	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"

	"time"
	"runtime"
	"os"
)

	// gl "github.com/go-gl/gl/v3.1/gles2"

	// "github.com/go-gl/gl"
	// glfw "github.com/go-gl/glfw3"

func init() {
	runtime.LockOSThread()
}

	// vertexBuffer	gl.Buffer

var (
	vertexBuffer	gl.GenBuffer()
	program 		gl.Program
	vertexAttrib 	gl.AttribLocation
)

func setup() {
	gl.ClearColor(0.0, 0.0, 0.4, 0.0)

	makeProgram(vertexShaderSource,fragmentShaderSource)

	vertexBufferData := []float32{
		-1,-1,0,
		 1,-1,0,
		 0, 1,0,
	}

	vertexBuffer = gl.GenBuffer()
	vertexBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexBufferData)*4, vertexBufferData, gl.STATIC_DRAW)
}

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	program.Use()
    // first attribute buffer: vertices
    var vertexAttrib = program.GetAttribLocation("vertexPosition_modelspace")
    vertexAttrib.EnableArray()
    vertexBuffer.Bind(gl.ARRAY_BUFFER)
    var f float32 = 0.0
    
    vertexAttrib.AttribPointer(
        3,     // size
        gl.FLOAT, // type
        false, // normalized
        0,     // stride
        nil) // array buffer offset
        // &f) // array buffer offset

    // draw the triangle
    gl.DrawArrays(gl.TRIANGLES, 0, 3)

    vertexAttrib.DisableArray()
	
}

func main() {
	glfw.SetErrorCallback(errorCallback)

	if !glfw.Init() {
		panic("Failed to initialize GLFW")
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
  	glfw.WindowHint(glfw.ContextVersionMinor, 0)

	window, err := glfw.CreateWindow(400, 300, "Last Chance", nil, nil)
	if err != nil {
		panic(err)
	}

	window.SetFramebufferSizeCallback(resizeCallback)
	window.SetKeyCallback(keyCallback)

	window.MakeContextCurrent()

	glfw.SwapInterval(1)

	width, height := window.GetFramebufferSize()
	resizeCallback(window, width, height)

	gl.Init()
	fmt.Println(gl.GetString(gl.VERSION))
	
	setup()

	render := time.Tick(16 * time.Millisecond)
	for !window.ShouldClose() {
		glfw.PollEvents()
		draw()
		window.SwapBuffers()
		<- render
	}
}

// ====================== GLFW CALLBACKS =============================
func keyCallback(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}
	switch glfw.Key(k) {
	case glfw.KeyEscape:
		window.SetShouldClose(true)
	default:
		return
	}
}

func resizeCallback(window *glfw.Window, width, height int) {
	wX,wY := window.GetSize()
	fX,fY := window.GetFramebufferSize()

	fmt.Println("resize: window",wX,wY,"framebuffer",fX,fY)
	gl.Viewport(0, 0, width, height)
	
	window.SwapBuffers()
}

func errorCallback(err glfw.ErrorCode, desc string) {
	fmt.Printf("%v: %v\n", err, desc)
}


// ====================== SHADER SOURCES =============================
var vertexShaderSource = `#version 120
// Input vertex data, different for all executions of this shader.
attribute vec3 vertexPosition_modelspace;
void main(){
	gl_Position = vec4(vertexPosition_modelspace, 1.0);
}`

var fragmentShaderSource = `#version 120
void main()
{
	// Output color = red 
	gl_FragColor = vec4(1,0,0,1);
}`


// ====================== MAKE PROGRAM =============================
func makeProgram(vertexShaderSource, fragmentShaderSource string) {

	// Create the shaders
	VertexShader := gl.CreateShader(gl.VERTEX_SHADER) 
	FragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)

	// Compile Vertex Shader
	VertexShader.Source(vertexShaderSource)
	VertexShader.Compile()

	// Check Vertex Shader
	status := VertexShader.Get(gl.COMPILE_STATUS)
	infoLog := VertexShader.GetInfoLog()
	if status != 1 {
		fmt.Println("Vertex shader:",status,infoLog)
		os.Exit(1)
	}
	// Compile Fragment Shader
	FragmentShader.Source(fragmentShaderSource)
	FragmentShader.Compile()

	// Check Fragment Shader
	status = FragmentShader.Get(gl.COMPILE_STATUS)
	infoLog = FragmentShader.GetInfoLog()
	if status != 1 {
		fmt.Println("Fragment shader:",status,infoLog)
		os.Exit(1)
	}

	// Link the program
	program = gl.CreateProgram()
	program.AttachShader(VertexShader)
	program.AttachShader(FragmentShader)
	program.Link()

	// Check program
	status = program.Get(gl.LINK_STATUS)
	infoLog = program.GetInfoLog()
	if status != 1 {
		fmt.Println("Program:",status,infoLog)
		os.Exit(1)
	}
}