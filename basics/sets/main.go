package main

import (
	"fmt"
)

func main() {
	fruits := map[string]struct{}{
		"avocado": {},
		"tomato":  {},
		"banana":  {},
	}
	vegetables := map[string]struct{}{
		"beets":   {},
		"carrots": {},
		"tomato":  {},
	}

	union := make(map[string]struct{})
	intersection := make(map[string]struct{})
	difference1 := make(map[string]struct{})
	difference2 := make(map[string]struct{})

	// Union
	for k := range fruits {
		union[k] = struct{}{}
	}
	for k := range vegetables {
		union[k] = struct{}{}
	}

	// Intersection
	for k := range fruits {
		if _, found := vegetables[k]; found {
			intersection[k] = struct{}{}
		}
	}

	// Difference1
	for k := range fruits {
		if _, found := vegetables[k]; !found {
			difference1[k] = struct{}{}
		}
	}

	// Difference2
	for k := range vegetables {
		if _, found := fruits[k]; !found {
			difference2[k] = struct{}{}
		}
	}

	fmt.Println(union)        // Output: map[avocado:{} banana:{} beets:{} carrots:{} tomato:{}]
	fmt.Println(intersection) // Output: map[tomato:{}]
	fmt.Println(difference1)  // Output: map[avocado:{} banana:{}]
	fmt.Println(difference2)  // Output: map[beets:{} carrots:{}]
}
