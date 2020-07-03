package main

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	dxSlice := make([]uint8, dx)
	for i := range result {
		for j := range dxSlice {
			dxSlice[j] = uint8(i ^ j)
		}
		result[i] = dxSlice
	}
	return result
}

// func main() {
// 	pic.Show(Pic)
// }
