package imagehash2

import (
	"image"

	"github.com/vitali-fedulov/images4"
)

const one9th float64 = 1.0 / 9.0

// lumaVector9 returns a 9D vector of everage luma values for 3x3 pixel blocks
// of the icon inner 9x9 block. The outer 1 pixel icon frame pixels are unused,
// but could be used in future versions as average value over all those frame
// pixels, making the 10th luma value. Excluding the 1px frame could be beneficial
// because it sometimes contains noise, like watermarks and white space.
func lumaVector9(icon images4.IconT) (v []float64) {

	v = make([]float64, 0)
	var sum float64

	// For each 3x3 pixel block in the center of the 11x11 icon.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			// Initialize sum for each block.
			sum = 0

			// Iterate over the 3x3 pixel block.
			for m := 0; m < 3; m++ {
				for n := 0; n < 3; n++ {

					// Accumulate the sum in the 3x3 block.
					c1, _, _ := images4.Get(
						icon,
						images4.IconSize,
						image.Point{
							// 1 for ignored border pixel.
							1 + 3*i + m,
							1 + 3*j + n,
							// For min i and m: x = 1 + 3*0 + 0 = 1.
							// For max i and m: x = 1 + 3*2 + 2 = 9.
							// For mid i and m (icon center pixel):
							// x = 1 + 3*1 + 1 = 5. Correct.
						},
					)

					sum += c1
				}
			}

			v = append(v, sum*one9th)
		}
	}

	return v
}
