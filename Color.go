package f32i

import (
	"image/color"
	"math"
)

var F32ColorModel = color.ModelFunc(func(c color.Color) color.Color {
	if _, ok := c.(F32Color); ok {
		return c
	}
	r, g, b, a := c.RGBA()

	return F32Color4{_ToFbound(r), _ToFbound(g), _ToFbound(b), _ToFbound(a)}
})

type F32Color interface {
	color.Color
	Layout() PixLayout
	F32(ct PixLayout) float32
}

func NewColor(f32 ...float32) F32Color {
	switch len(f32) {
	case 1:
		return F32Color1{f32[0]}
	case 2:
		return F32Color2{f32[0], f32[1]}
	case 3:
		return F32Color3{f32[0], f32[1], f32[2]}
	case 4:
		return F32Color4{f32[0], f32[1], f32[2], f32[3]}
	default:
		return nil
	}
}

type F32Color1 [1]float32

func (s F32Color1) RGBA() (r, g, b, a uint32) {
	r = _fbound(s[0], 0, math.MaxUint16)
	g = r
	b = r
	a = math.MaxUint16
	return
}
func (s F32Color1) Layout() PixLayout {
	return P0
}
func (s F32Color1) F32(ct PixLayout) float32 {
	switch ct & _Under4Mask {
	case P0:
		return s[0]
	}
	return nan
}

type F32Color2 [2]float32

func (s F32Color2) RGBA() (r, g, b, a uint32) {
	r = _fbound(s[0], 0, math.MaxUint16)
	g = _fbound(s[1], 0, math.MaxUint16)
	b = 0
	a = math.MaxUint16
	return
}
func (s F32Color2) Layout() PixLayout {
	return P0 | P1
}
func (s F32Color2) F32(ct PixLayout) float32 {
	switch ct & _Under4Mask {
	case P0:
		return s[0]
	case P1:
		return s[1]
	}
	return nan
}

type F32Color3 [3]float32

func (s F32Color3) RGBA() (r, g, b, a uint32) {
	r = _fbound(s[0], 0, math.MaxUint16)
	g = _fbound(s[1], 0, math.MaxUint16)
	b = _fbound(s[2], 0, math.MaxUint16)
	a = math.MaxUint16
	return
}
func (s F32Color3) Layout() PixLayout {
	return P0 | P1 | P2
}
func (s F32Color3) F32(ct PixLayout) float32 {
	switch ct & _Under4Mask {
	case P0:
		return s[0]
	case P1:
		return s[1]
	case P2:
		return s[2]
	}
	return nan
}

type F32Color4 [4]float32

func (s F32Color4) RGBA() (r, g, b, a uint32) {
	r = _fbound(s[0], 0, math.MaxUint16)
	g = _fbound(s[1], 0, math.MaxUint16)
	b = _fbound(s[2], 0, math.MaxUint16)
	a = _fbound(s[3], 0, math.MaxUint16)
	return
}
func (s F32Color4) Layout() PixLayout {
	return P0 | P1 | P2 | P3
}
func (s F32Color4) F32(ct PixLayout) float32 {
	switch ct & _Under4Mask {
	case P0:
		return s[0]
	case P1:
		return s[1]
	case P2:
		return s[2]
	case P3:
		return s[3]
	}
	return nan
}
