package f32i

import (
	"image"
	"image/color"
)

type Image32 struct {
	Pix       []float32
	Stride    int
	PixLayout PixLayout
	Rect      image.Rectangle
}

func NewImage(r image.Rectangle, ColorType PixLayout) *Image32 {
	w, h := r.Dx(), r.Dy()

	return &Image32{
		Pix:       make([]float32, ColorType.Count()*w*h),
		Stride:    ColorType.Count() * w,
		PixLayout: ColorType,
		Rect:      r,
	}
}
func (s *Image32) ColorModel() color.Model {
	return F32ColorModel
}
func (s *Image32) Bounds() image.Rectangle {
	return s.Rect
}
func (s *Image32) PixOffset(x, y int) int {
	return (y-s.Rect.Min.Y)*s.Stride + (x-s.Rect.Min.X)*s.PixLayout.Count()
}
func (s *Image32) At(x, y int) color.Color {
	return s.F32ColorAt(x, y)
}
func (s *Image32) F32ColorAt(x, y int) F32Color {
	offset := s.PixOffset(x, y)
	switch s.PixLayout.Count() {
	case 1:
		return F32Color1{
			s.Pix[offset] + 0,
		}
	case 2:
		return F32Color2{
			s.Pix[offset+0],
			s.Pix[offset+1],
		}
	case 3:
		return F32Color3{
			s.Pix[offset+0],
			s.Pix[offset+1],
			s.Pix[offset+2],
		}
	case 4:
		return F32Color4{
			s.Pix[offset+0],
			s.Pix[offset+1],
			s.Pix[offset+2],
			s.Pix[offset+3],
		}
	default:
		panic("Image32's Count must be one of [1,2,3,4]")
	}
}
func (s *Image32) Set(x, y int, c color.Color) {
	offset := s.PixOffset(x, y)
	f32c := F32ColorModel.Convert(c).(F32Color)
	switch s.PixLayout.Count() {
	case 4:
		s.Pix[offset+3] = f32c.F32(P3)
		fallthrough
	case 3:
		s.Pix[offset+2] = f32c.F32(P2)
		fallthrough
	case 2:
		s.Pix[offset+1] = f32c.F32(P1)
		fallthrough
	case 1:
		s.Pix[offset+0] = f32c.F32(P0)
		return
	default:
		panic("Image32's Count must be one of [1,2,3,4]")
	}
}

//func (s *Image32) WithPicker(pixLayout PixLayout) Picker {
//	if s.CanPicker(pixLayout) {
//		return &Image32Picker{
//			ref:    s,
//			picker: pixLayout,
//		}
//	}
//	return nil
//}
//func (s *Image32) CanPicker(pixLayout PixLayout) (available bool) {
//	return s.PixLayout&pixLayout == pixLayout
//}
//
//type Image32Picker struct {
//	ref    *Image32
//	picker PixLayout
//}
//
//func (s *Image32Picker) Layout() PixLayout {
//	return s.picker
//}
//func (s *Image32Picker) Bound() image.Rectangle {
//	return s.ref.Rect
//}
//func (s *Image32Picker) Pick(x, y int) (res []float32) {
//	offset := s.ref.PixOffset(x, y)
//	if s.picker&P0 == P0 {
//		res = append(res, s.ref.Pix[offset+0])
//	}
//	if s.picker&P1 == P1 {
//		res = append(res, s.ref.Pix[offset+1])
//	}
//	if s.picker&P2 == P2 {
//		res = append(res, s.ref.Pix[offset+2])
//	}
//	if s.picker&P3 == P3 {
//		res = append(res, s.ref.Pix[offset+3])
//	}
//	return
//}
