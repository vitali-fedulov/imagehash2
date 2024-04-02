package imagehash2

import (

	// N-dimensional space discretization and hashing package.
	"github.com/vitali-fedulov/hyper"

	// Icon generation package.
	"github.com/vitali-fedulov/images4"
)

// CentralHash9 generates a central hash for an icon.
// This hash can then be used for a record or a query.
// When used for a record, you will need a hash set made with func HashSet
// for a query. And vice versa. They are interchangeable for optimization.
// The hash is generated for 9 average luma values of nine 3x3 pixel blocks
// in the central 9x9 area of the icon.
func CentralHash9(icon images4.IconT, epsilon float64, numBuckets int) uint64 {

	vector := lumaVector9(icon)

	// Central hypercube made with package hyper.
	cube := hyper.CentralCube(vector,
		hyper.Params{
			Min:        0,
			Max:        255,
			EpsPercent: epsilon,
			NumBuckets: numBuckets})

	// Preventing hash overflow.
	if numBuckets > 10 { //  or len(vector) > 19, which is not.
		return cube.FNV1aHash()
	}

	return cube.DecimalHash()
}

// HashSet9 generates a hash set for an icon.
// This hash set can then be used for a record or a query.
// When used for a query, you will need a hash made with
// func CentralHash as a record. And vice versa.
func HashSet9(icon images4.IconT, epsilon float64, numBuckets int) []uint64 {

	vector := lumaVector9(icon)

	// Hypercube set made with package hyper.
	cubeSet := hyper.CubeSet(vector,
		hyper.Params{
			Min:        0,
			Max:        255,
			EpsPercent: epsilon,
			NumBuckets: numBuckets})

	// Preventing hash overflow.
	if numBuckets > 10 { //  or len(vector) > 19, which is not.
		return cubeSet.HashSet((hyper.Cube).FNV1aHash)
	}

	return cubeSet.HashSet((hyper.Cube).DecimalHash)
}
