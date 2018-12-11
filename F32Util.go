package f32i

import "math"

// F32 NaN feature
var nan = float32(math.NaN())

func IsNaN(f32 float32) bool {
	return math.IsNaN(float64(f32))
}
func _fbound(f32 float32, min, max uint32) (res uint32) {
	res = uint32(f32 * math.MaxUint16)
	if res < min {
		res = min
	}
	if res > max {
		res = max
	}
	return res
}
func _ToFbound(u32 uint32) (res float32) {
	return float32(u32) / float32(math.MaxUint16)
}
