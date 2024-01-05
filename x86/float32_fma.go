//go:build fma

package x86

import _ "unsafe"

//go:linkname Float32FMAdd8 github.com/octu0/go-intrin/x86.immFMAddFloat32_256
//go:noescape
func Float32FMAdd8(a, b, c [8]float32) [8]float32