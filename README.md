# Fast similar image search with Go (LATEST version)

Resized and near-duplicate image search **for large image collections** (thousands, millions, and more). The package generates 'real' hashes to be used in hash-tables, and consumes very little memory. It is recommended to cross-check the similarity result with more precise (but slow) [images4](https://github.com/vitali-fedulov/images4) package.

[Demo](https://vitali-fedulov.github.io/similar.pictures/) (a usage scenario for image similarity search).

[Algorithm](https://vitali-fedulov.github.io/similar.pictures/algorithm-for-hashing-high-dimensional-float-vectors.html) for nearest neighbour vector search by vector quantization.

[Go doc](https://pkg.go.dev/github.com/vitali-fedulov/imagehash2).

Major (semantic) versions have their own repositories and are mutually incompatible:
| Major version | Repository | Comment |
| ----------- | ---------- | ----------|
| 2 | imagehash2 (this) | recommended, with improved precision |
| 1 | [imagehash](https://github.com/vitali-fedulov/imagehash) | as fast, but has a minor generalization issue |

## Parameters

The most important parameter is `numBuckets`. It defines granularity of hyper-space quantization. The higher the value, the more restrictive the comparison is. And, when used together with images4 package, higher `numBuckets` considerably accelerates the search process, because fewer image ids fall into a single quantization cell.

The second parameter is `epsilon`, which can be safely set to 0.25.

## Example of comparison for 2 photos

The demo shows only the hash-based similarity comparison (without making actual hash table). But the hash table, typically a Golang map, is implied in full implementation.

```go
package main

import (
	"fmt"
	"github.com/vitali-fedulov/imagehash2"
	"github.com/vitali-fedulov/images4"
)

const (
	// Recommended initial parameters.

	// Increase this value to get higher precision.
	numBuckets = 4

	// No need to change epsilon value.
	epsilon = 0.25
)

func main() {

	// Open and decode photos (skipping error handling for clarity).
	img1, _ := images4.Open("1.jpg")
	img2, _ := images4.Open("2.jpg")

	// Icons are compact image representations needed for comparison.
	icon1 := images4.Icon(img1)
	icon2 := images4.Icon(img2)

	// Hash table values.

	// Value to save to the hash table as a KEY with corresponding
	// image ids. Table structure: map[centralHash][]imageId.
	// imageId is simply an image number in a directory tree.
	// centralHash is uint64.
	centralHash := imagehash2.CentralHash9(icon1, epsilon, numBuckets)

	// Hash set to be used as a QUERY to the hash table.
	// Each hash from the hashSet must be checked against the hash table.
	// The length of hashSet is different for each image.
	// The most frequent length is 1.
	hashSet := imagehash2.HashSet9(icon2, epsilon, numBuckets)

	// As if we are searching in the table.

	foundSimilarImage := false

	// Checking hash matches. In full implementation this will be done
	// on the map mentioned above.
	for _, hash := range hashSet { // Query. Check full hashSet.

		if centralHash == hash { // Sub-query hash found in the table.
			foundSimilarImage = true
			break
		}

	}

	// Comparison result.
	if foundSimilarImage {

		fmt.Println("Images are *approximately* similar.")

		// It is recommended to cross-check the result with
	        // the higher-precision func Similar from package images4.
		if images4.Similar(icon1, icon2) == true {
			fmt.Println("Images are similar")
		}

	} else {
		fmt.Println("Images are distinct.")
	}

}
```
