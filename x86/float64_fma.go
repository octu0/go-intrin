//go:build fma

package x86

import _ "unsafe"

//go:linkname Float64FMAdd github.com/octu0/go-intrin/x86.immFMAddFloat64_128
//go:noescape
func Float64FMAdd(a, b, c [2]float64) [2]float64

//go:linkname Float64FMAdd4 github.com/octu0/go-intrin/x86.immFMAddFloat64_256
//go:noescape
func Float64FMAdd4(a, b, c [4]float64) [4]float64
