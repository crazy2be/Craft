package main

import (
	"log"
	"math"
	"io/ioutil"
	"github.com/go-gl/gl"
)

func load_file(name string) []byte {
	contents, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal("read file: ", err)
	}
	return contents
}

func make_buffer(target gl.GLenum, size int, data interface{}) gl.Buffer {
	buf := gl.GenBuffer()
	buf.Bind(target)
	gl.BufferData(target, size, data, gl.STATIC_DRAW)
	buf.Unbind(target)
	return buf
}

func make_shader(type_ gl.GLenum, src []byte) gl.Shader {
	shader := gl.CreateShader(type_)
	shader.Source(string(src))
	shader.Compile()
	status := shader.Get(gl.COMPILE_STATUS)
	if status == gl.FALSE {
		log.Fatal("shader:", shader.GetInfoLog())
	}
	return shader
}

func load_shader(type_ gl.GLenum, path string) gl.Shader {
	return make_shader(type_, load_file(path))
}

func make_program(shader1, shader2 gl.Shader) gl.Program {
	program := gl.CreateProgram()
	program.AttachShader(shader1)
	program.AttachShader(shader2)
	program.Link()
	status := program.Get(gl.LINK_STATUS)
	if status == gl.FALSE {
		log.Fatal("link:", program.GetInfoLog())
	}
	return program
}

func load_program(type1, type2 gl.GLenum, path1, path2 string) gl.Program {
	return make_program(load_shader(type1, path1), load_shader(type2, path2))
}

type Matrix [16]float32

func (m *Matrix) Identity() {
	m[0] = 1;
    m[1] = 0;
    m[2] = 0;
    m[3] = 0;
    m[4] = 0;
    m[5] = 1;
    m[6] = 0;
    m[7] = 0;
    m[8] = 0;
    m[9] = 0;
    m[10] = 1;
    m[11] = 0;
    m[12] = 0;
    m[13] = 0;
    m[14] = 0;
    m[15] = 1;
}

func (m *Matrix) Frustum(left, right, bottom, top, znear, zfar float32) {
    temp := 2.0 * znear;
    temp2 := right - left;
    temp3 := top - bottom;
    temp4 := zfar - znear;
    m[0] = temp / temp2;
    m[1] = 0.0;
    m[2] = 0.0;
    m[3] = 0.0;
    m[4] = 0.0;
    m[5] = temp / temp3;
    m[6] = 0.0;
    m[7] = 0.0;
    m[8] = (right + left) / temp2;
    m[9] = (top + bottom) / temp3;
    m[10] = (-zfar - znear) / temp4;
    m[11] = -1.0;
    m[12] = 0.0;
    m[13] = 0.0;
    m[14] = (-temp * zfar) / temp4;
    m[15] = 0.0;
}

func (m *Matrix) Perspective(fov, aspect, znear, zfar float32) {
    ymax := znear * float32(math.Tan(float64(fov * math.Pi / 360.0)))
    xmax := ymax * aspect
    m.Frustum(-xmax, xmax, -ymax, ymax, znear, zfar)
}

func make_cube(output []float32, x, y, z, n float32) {
	i := 0
	i++; output[i-1] =  x - n;
	i++; output[i-1] =  y - n;
	i++; output[i-1] =  z - n;
	i++; output[i-1] =  x - n;
	i++; output[i-1] =  y - n;
	i++; output[i-1] =  z + n;
	i++; output[i-1] =  x - n;
	i++; output[i-1] =  y + n;
	i++; output[i-1] =  z - n;
	i++; output[i-1] =  x - n;
	i++; output[i-1] =  y + n;
	i++; output[i-1] =  z + n;
	i++; output[i-1] =  x + n;
	i++; output[i-1] =  y - n;
	i++; output[i-1] =  z - n;
	i++; output[i-1] =  x + n;
	i++; output[i-1] =  y - n;
	i++; output[i-1] =  z + n;
	i++; output[i-1] =  x + n;
	i++; output[i-1] =  y + n;
	i++; output[i-1] =  z - n;
	i++; output[i-1] =  x + n;
	i++; output[i-1] =  y + n;
	i++; output[i-1] =  z + n;
}
