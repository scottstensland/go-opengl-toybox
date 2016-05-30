


package main

import (
"fmt"
gl "github.com/chsc/gogl/gl33"
"github.com/veandco/go-sdl2/sdl"
// "math"
"runtime"
"time"
)

    // "github.com/go-gl/glfw/v3.1/glfw"


func createprogram() gl.Uint {
// VERTEX SHADER
vs := gl.CreateShader(gl.VERTEX_SHADER)
vs_source := gl.GLString(vertexShaderSource)
gl.ShaderSource(vs, 1, &vs_source, nil)
gl.CompileShader(vs)
var vs_status gl.Int
gl.GetShaderiv(vs, gl.COMPILE_STATUS, &vs_status)
fmt.Printf("Compiled Vertex Shader: %v\n", vs_status)

// FRAGMENT SHADER
fs := gl.CreateShader(gl.FRAGMENT_SHADER)
fs_source := gl.GLString(fragmentShaderSource)
gl.ShaderSource(fs, 1, &fs_source, nil)
gl.CompileShader(fs)
var fstatus gl.Int
gl.GetShaderiv(fs, gl.COMPILE_STATUS, &fstatus)
fmt.Printf("Compiled Fragment Shader: %v\n", fstatus)

// CREATE PROGRAM
program := gl.CreateProgram()
gl.AttachShader(program, vs)
gl.AttachShader(program, fs)
fragoutstring := gl.GLString("outColor")
defer gl.GLStringFree(fragoutstring)
gl.BindFragDataLocation(program, gl.Uint(0), fragoutstring)

gl.LinkProgram(program)
var linkstatus gl.Int
gl.GetProgramiv(program, gl.LINK_STATUS, &linkstatus)
fmt.Printf("Program Link: %v\n", linkstatus)

return program
}

func main() {
var window *sdl.Window
var context sdl.GLContext
var event sdl.Event
var running bool
var err error
runtime.LockOSThread()
if 0 != sdl.Init(sdl.INIT_EVERYTHING) {
panic(sdl.GetError())
}
window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED,
sdl.WINDOWPOS_UNDEFINED,
winWidth, winHeight, sdl.WINDOW_OPENGL)
if err != nil {
panic(err)
}
if window == nil {
panic(sdl.GetError())
}
context = sdl.GL_CreateContext(window)
if context == nil {
panic(sdl.GetError())
}

gl.Init()
gl.Viewport(0, 0, gl.Sizei(winWidth), gl.Sizei(winHeight))
// OPENGL FLAGS
gl.ClearColor(0.0, 0.1, 0.0, 1.0)
gl.Enable(gl.DEPTH_TEST)
gl.DepthFunc(gl.LESS)
gl.Enable(gl.BLEND)
gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

// VERTEX BUFFER
var vertexbuffer gl.Uint
gl.GenBuffers(1, &vertexbuffer)
gl.BindBuffer(gl.ARRAY_BUFFER, vertexbuffer)
gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(triangle_vertices)*4), gl.Pointer(&triangle_vertices[0]), gl.STATIC_DRAW)

// COLOUR BUFFER
var colourbuffer gl.Uint
gl.GenBuffers(1, &colourbuffer)
gl.BindBuffer(gl.ARRAY_BUFFER, colourbuffer)
gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(triangle_colours)*4), gl.Pointer(&triangle_colours[0]), gl.STATIC_DRAW)

// GUESS WHAT
program := createprogram()

// VERTEX ARRAY
var VertexArrayID gl.Uint
gl.GenVertexArrays(1, &VertexArrayID)
gl.BindVertexArray(VertexArrayID)
gl.EnableVertexAttribArray(0)
gl.BindBuffer(gl.ARRAY_BUFFER, vertexbuffer)
gl.VertexAttribPointer(0, 3, gl.FLOAT, gl.FALSE, 0, nil)

// VERTEX ARRAY HOOK COLOURS
gl.EnableVertexAttribArray(1)
gl.BindBuffer(gl.ARRAY_BUFFER, colourbuffer)
gl.VertexAttribPointer(1, 3, gl.FLOAT, gl.FALSE, 0, nil)

gl.UseProgram(program)

running = true
for running {
for event = sdl.PollEvent(); event != nil; event =
sdl.PollEvent() {
switch t := event.(type) {
case *sdl.QuitEvent:
running = false
case *sdl.MouseMotionEvent:

fmt.Printf(string(t.State))
}
drawgl()
sdl.GL_SwapWindow(window)
}
}
sdl.GL_DeleteContext(context)
window.Destroy()
sdl.Quit()
}

func drawgl() {

gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
gl.DrawArrays(gl.TRIANGLES, gl.Int(0), gl.Sizei(len(triangle_vertices)*4))

time.Sleep(50 * time.Millisecond)

}

const (
winTitle = "OpenGL Shader"
winWidth = 640
winHeight = 480
vertexShaderSource = `
#version 330 core

// Input vertex data, different for all executions of this shader.
layout(location = 0) in vec3 vertexPosition_modelspace;

void main(){

gl_Position.xyz = vertexPosition_modelspace;
gl_Position.w = 1.0;

}
`
fragmentShaderSource = `
#version 330 core

// Ouput data
out vec3 color;

void main()
{

// Output color = red 
color = vec3(1,0,0);

}
`
)

var triangle_vertices = []gl.Float{
-1.0, -1.0, -1.0,
-1.0, -1.0, 1.0,
-1.0, 1.0, 1.0,
1.0, 1.0, -1.0,
-1.0, -1.0, -1.0,
-1.0, 1.0, -1.0,
1.0, -1.0, 1.0,
-1.0, -1.0, -1.0,
1.0, -1.0, -1.0,
1.0, 1.0, -1.0,
1.0, -1.0, -1.0,
-1.0, -1.0, -1.0,
-1.0, -1.0, -1.0,
-1.0, 1.0, 1.0,
-1.0, 1.0, -1.0,
1.0, -1.0, 1.0,
-1.0, -1.0, 1.0,
-1.0, -1.0, -1.0,
-1.0, 1.0, 1.0,
-1.0, -1.0, 1.0,
1.0, -1.0, 1.0,
1.0, 1.0, 1.0,
1.0, -1.0, -1.0,
1.0, 1.0, -1.0,
1.0, -1.0, -1.0,
1.0, 1.0, 1.0,
1.0, -1.0, 1.0,
1.0, 1.0, 1.0,
1.0, 1.0, -1.0,
-1.0, 1.0, -1.0,
1.0, 1.0, 1.0,
-1.0, 1.0, -1.0,
-1.0, 1.0, 1.0,
1.0, 1.0, 1.0,
-1.0, 1.0, 1.0,
1.0, -1.0, 1.0}

var triangle_colours = []gl.Float{
0.583, 0.771, 0.014,
0.609, 0.115, 0.436,
0.327, 0.483, 0.844,
0.822, 0.569, 0.201,
0.435, 0.602, 0.223,
0.310, 0.747, 0.185,
0.597, 0.770, 0.761,
0.559, 0.436, 0.730,
0.359, 0.583, 0.152,
0.483, 0.596, 0.789,
0.559, 0.861, 0.639,
0.195, 0.548, 0.859,
0.014, 0.184, 0.576,
0.771, 0.328, 0.970,
0.406, 0.615, 0.116,
0.676, 0.977, 0.133,
0.971, 0.572, 0.833,
0.140, 0.616, 0.489,
0.997, 0.513, 0.064,
0.945, 0.719, 0.592,
0.543, 0.021, 0.978,
0.279, 0.317, 0.505,
0.167, 0.620, 0.077,
0.347, 0.857, 0.137,
0.055, 0.953, 0.042,
0.714, 0.505, 0.345,
0.783, 0.290, 0.734,
0.722, 0.645, 0.174,
0.302, 0.455, 0.848,
0.225, 0.587, 0.040,
0.517, 0.713, 0.338,
0.053, 0.959, 0.120,
0.393, 0.621, 0.362,
0.673, 0.211, 0.457,
0.820, 0.883, 0.371,
0.982, 0.099, 0.879}




Update:

package main

import (
"fmt"
// gl "github.com/chsc/gogl/gl33"
"github.com/veandco/go-sdl2/sdl"
// "math"
"github.com/Jragonmiris/mathgl"
"github.com/go-gl/gl"
"runtime"
"time"
)

// var program gl.Program = 0
// var buffer gl.Buffer = 0

func MakeProgram(vert, frag string) gl.Program {

vertShader, fragShader := gl.CreateShader(gl.VERTEX_SHADER), gl.CreateShader(gl.FRAGMENT_SHADER)
vertShader.Source(vert)
fragShader.Source(frag)

vertShader.Compile()
fragShader.Compile()

prog := gl.CreateProgram()

prog.AttachShader(vertShader)
prog.AttachShader(fragShader)
prog.Link()
prog.Validate()
fmt.Println(prog.GetInfoLog())

return prog
}

func main() {
var window *sdl.Window
var context sdl.GLContext
var event sdl.Event
var running bool
var err error

runtime.LockOSThread()

if 0 != sdl.Init(sdl.INIT_EVERYTHING) {
panic(sdl.GetError())
}
window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED,
sdl.WINDOWPOS_UNDEFINED,
winWidth, winHeight, sdl.WINDOW_OPENGL)
if err != nil {
panic(err)
}
if window == nil {
panic(sdl.GetError())
}
context = sdl.GL_CreateContext(window)
if context == nil {
panic(sdl.GetError())
}

if gl.Init() != 0 {
panic("gl error")
}

gl.ClearColor(1.0, 1.0, 1.0, .5)
gl.Viewport(0, 0, winWidth, winHeight)

program := MakeProgram(vertexShaderSource, fragmentShaderSource)
defer program.Delete()

matrixID := program.GetUniformLocation("MVP")
Projection := mathgl.Perspective(45.0, 4.0/3.0, 0.1, 100.0)
View := mathgl.LookAt(4.0, 3.0, 3.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)
Model := mathgl.Ident4f()
MVP := Projection.Mul4(View).Mul4(Model) 

gl.Enable(gl.DEPTH_TEST)
gl.DepthFunc(gl.LESS)
gl.Enable(gl.BLEND)
gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

vertexArray := gl.GenVertexArray()
defer vertexArray.Delete()
vertexArray.Bind()

buffer := gl.GenBuffer()
defer buffer.Delete()
buffer.Bind(gl.ARRAY_BUFFER)
gl.BufferData(gl.ARRAY_BUFFER, len(triangle_vertices)*4, &triangle_vertices, gl.STATIC_DRAW)

running = true
for running {
for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
switch t := event.(type) {
case *sdl.QuitEvent:
running = false
case *sdl.MouseMotionEvent:

fmt.Printf(string(t.Timestamp))
}
}

gl.Clear(gl.COLOR_BUFFER_BIT) // | gl.DEPTH_BUFFER_BIT)
program.Use()
matrixID.UniformMatrix4fv(false, MVP)
attribLoc := gl.AttribLocation(0)
attribLoc.EnableArray()
buffer.Bind(gl.ARRAY_BUFFER)
attribLoc.AttribPointer(3, gl.FLOAT, false, 0, nil)

gl.DrawArrays(gl.TRIANGLES, 0, 3)

attribLoc.DisableArray()

time.Sleep(50 * time.Millisecond)

sdl.GL_SwapWindow(window)
}

sdl.GL_DeleteContext(context)
window.Destroy()
sdl.Quit()
}



const (
winTitle = "OpenGL Shader"
winWidth = 640
winHeight = 480
vertexShaderSource = `
#version 330 core

// Input vertex data, different for all executions of this shader.
layout(location = 0) in vec3 vertexPosition_modelspace;
// Values that stay constant for the whole mesh.
uniform mat4 MVP;
void main(){

gl_Position = MVP * vec4 (vertexPosition_modelspace,1.0);

}
`
fragmentShaderSource = `
#version 330 core

// Ouput data
out vec3 color;

void main()
{

// Output color = red 
color = vec3(1,0,0);

}
`
)

var triangle_vertices = []float32{
-.5, -.5, -.5,
.5, -.5, -.5,
0.0, 0.5, -.5,
}

