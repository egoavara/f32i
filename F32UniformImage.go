package f32i

import (
	"image"
	"image/color"
)

type Uniform32 struct {
	Pix  F32Color
	Rect image.Rectangle
}

func NewUniform(r image.Rectangle, p F32Color) *Uniform32 {
	p.Layout()
	return &Uniform32{
		Rect: r,
		Pix:  p,
	}
}
func (s *Uniform32) ColorModel() color.Model {
	return F32ColorModel
}
func (s *Uniform32) Bounds() image.Rectangle {
	return s.Rect
}
func (s *Uniform32) At(x, y int) color.Color {
	return s.Pix
}
func (s *Uniform32) Set(x, y int, c color.Color) {
	s.Pix = F32ColorModel.Convert(c).(F32Color)
}

//
//func (s *Uniform32) WithPicker(pixLayout PixLayout) Picker {
//	if s.CanPicker(pixLayout) {
//		return &Uniform32Picker{
//			ref:    s,
//			picker: pixLayout,
//		}
//	}
//	return nil
//}
//func (s *Uniform32) CanPicker(pixLayout PixLayout) (available bool) {
//
//	return s.Pix.Layout()&pixLayout == pixLayout
//}
//
//type Uniform32Picker struct {
//	ref    *Uniform32
//	picker PixLayout
//}
//
//func (s *Uniform32Picker) Layout() PixLayout {
//	return s.picker
//}
//func (s *Uniform32Picker) Bound() image.Rectangle {
//	return s.ref.Rect
//}
//func (s *Uniform32Picker) Pick(x, y int) (res []float32) {
//	if s.picker&P0 == P0 {
//		res = append(res, s.ref.Pix.F32(P0))
//	}
//	if s.picker&P1 == P1 {
//		res = append(res, s.ref.Pix.F32(P1))
//	}
//	if s.picker&P2 == P2 {
//		res = append(res, s.ref.Pix.F32(P2))
//	}
//	if s.picker&P3 == P3 {
//		res = append(res, s.ref.Pix.F32(P3))
//	}
//	return
//}
