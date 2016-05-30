package main

// https://github.com/go-gl/glfw

// go get -u github.com/go-gl/glfw/v3.1/glfw

import (
    "runtime"
    "github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
    // This is needed to arrange that main() runs on main thread.
    // See documentation for functions that are only allowed to be called from the main thread.
    runtime.LockOSThread()
}

func main() {
    err := glfw.Init()
    if err != nil {
        panic(err)
    }
    defer glfw.Terminate()

    window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
    if err != nil {
        panic(err)
    }

    window.MakeContextCurrent()




// fmt.Println(gl.GetString(gl.VERSION))
// fmt.Println(gl.GetString(gl.VENDOR))
// fmt.Println(gl.GetString(gl.RENDERER))




    for !window.ShouldClose() {
        // Do OpenGL stuff.
        window.SwapBuffers()
        glfw.PollEvents()
    }
}

