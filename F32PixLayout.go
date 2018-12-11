package f32i

// bitcount
var _WeightTable = [16]int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4}

const _Under4Mask = 0x0F
const _Upper4Mask = 0xF0

// 0x 0000 0000
//    ^^^^ : Type
// 0x 0000 0000
//         ^^^^ : Pixel Exist
type PixLayout uint8

func (s PixLayout) String() string {
	var upper = s & _Upper4Mask
	var under = s & _Under4Mask
	//
	var res = "["
	switch upper {
	case TypeUndefined:
		if under&P0 == P0 {
			res += "P0, "
		}
		if under&P1 == P1 {
			res += "P1, "
		}
		if under&P2 == P2 {
			res += "P2, "
		}
		if under&P3 == P3 {
			res += "P3, "
		}
	case TypeRGBA:

		if under&P0 == P0 {
			res += "R, "
		}
		if under&P1 == P1 {
			res += "G, "
		}
		if under&P2 == P2 {
			res += "B, "
		}
		if under&P3 == P3 {
			res += "A, "
		}
	case TypePBR:
		if under&P0 == P0 {
			res += "Metallic, "
		}
		if under&P1 == P1 {
			res += "Roughness, "
		}
		if under&P2 == P2 {
			res += "_, "
		}
		if under&P3 == P3 {
			res += "_, "
		}
	}

	if len(res) > 2 {
		res = res[:len(res)-2]
	}
	res += "]"
	return res
}

func (s PixLayout) Count() int {
	return _WeightTable[s&_Under4Mask]
}

// Upper 4 bit, Type Notation
const (
	TypeUndefined PixLayout = 0x0 << 4
	TypeRGBA      PixLayout = 0x1 << 4
	TypePBR       PixLayout = 0x2 << 4
	// Type... PixLayout = 0x11 << 4 : Reserved
)

// Under 4 bit, Pixel Exist
const (
	P0 PixLayout = 1 << 0
	P1 PixLayout = 1 << 1
	P2 PixLayout = 1 << 2
	P3 PixLayout = 1 << 3
)

// RGBA Alias
const (
	R    = TypeRGBA | P0
	G    = TypeRGBA | P1
	B    = TypeRGBA | P2
	A    = TypeRGBA | P3
	RG   = R | G
	RGB  = R | G | B
	RGBA = R | G | B | A
)

// PBR Alias
const (
	Metallic  = TypePBR | P0
	Roughness = TypePBR | P1
)
