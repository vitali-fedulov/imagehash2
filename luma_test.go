package imagehash2

import (
	"image"
	"reflect"
	"testing"

	"github.com/vitali-fedulov/images4"
)

func TestLumaVector(t *testing.T) {

	// Preparing test icon.
	icon := images4.EmptyIcon()
	icon.Pixels = make([]uint16, images4.IconSize*images4.IconSize*3)
	// Random values.
	lumaValues := []float64{
		47, 14, 44, 4, 41, 35, 28, 21, 35, 14, 30,
		27, 6, 20, 26, 9, 45, 47, 43, 43, 0, 18,
		3, 47, 44, 5, 28, 2, 21, 4, 18, 8, 8,
		13, 20, 25, 47, 7, 22, 40, 50, 32, 35, 16,
		13, 13, 15, 36, 29, 37, 42, 9, 3, 5, 45,
		22, 26, 14, 50, 49, 10, 48, 4, 21, 7, 45,
		41, 48, 37, 2, 49, 3, 48, 47, 12, 46, 37,
		31, 42, 46, 42, 21, 28, 19, 29, 19, 17, 18,
		0, 20, 6, 14, 48, 21, 21, 13, 33, 25, 50,
		3, 11, 27, 1, 6, 32, 31, 25, 0, 12, 25,
		50, 35, 14, 3, 11, 19, 37, 27, 42, 19, 30,
	}
	// Filling in the test icon.
	for i := 0; i < images4.IconSize; i++ {
		for j := 0; j < images4.IconSize; j++ {
			images4.Set(icon, images4.IconSize,
				image.Point{i, j}, lumaValues[i*images4.IconSize+j], 0, 0)
		}
	}

	// Luma vector.
	got := lumaVector9(icon)

	// Expected luma vector.
	expected := []float64{
		26.666666666666664, 24.555555555555554, 25.888888888888886,
		26.777777777777775, 35, 17.11111111111111,
		23.22222222222222, 25.22222222222222, 19.22222222222222,
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf(
			`Luma vector not matching expected.
			 Got: %v, expected %v.`, got, expected,
		)
	}
}
