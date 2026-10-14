// MAPS
//
// A map is a key -> value collection (like an object/Map in JS, a dict in Python).
// Type: map[KeyType]ValueType
//
//   - The key must be a comparable type (string, int, simple struct...).
//   - ITERATION ORDER IS RANDOM — never depend on it.
//   - Reading a missing key returns the zero value (no error).
//   - Use "value, ok := m[key]" to check whether the key exists.
//   - A nil map can be read, but writing to it causes a panic: initialize it
//     with make() or a literal.
//   - Maps are references: passing one to a function does NOT copy the data.
//   - They are not safe for concurrent use without synchronization (sync.Mutex/sync.Map).
package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	// Creating with a literal
	ages := map[string]int{
		"Ana":   25,
		"Bruno": 30,
	}

	// Creating with make
	capitals := make(map[string]string)
	capitals["RS"] = "Porto Alegre"
	capitals["SP"] = "São Paulo"

	fmt.Println(ages, capitals, "size:", len(capitals))

	// Reading
	fmt.Println("Ana's age:", ages["Ana"])
	fmt.Println("age of someone who doesn't exist:", ages["Zé"]) // 0 (zero value)

	// Checking existence (comma ok idiom)
	if age, ok := ages["Zé"]; ok {
		fmt.Println("Zé is", age)
	} else {
		fmt.Println("Zé is not in the map")
	}

	// Updating and deleting
	ages["Ana"] = 26
	delete(ages, "Bruno")
	fmt.Println("after update/delete:", ages)

	// Iterating (random order!)
	for state, city := range capitals {
		fmt.Println(state, "->", city)
	}

	// To iterate in order: sort the keys
	keys := slices.Sorted(maps.Keys(capitals))
	fmt.Println("sorted keys:", keys)

	// Map with a slice as the value
	tags := map[string][]string{
		"go": {"backend", "concurrency"},
	}
	tags["go"] = append(tags["go"], "compiled")
	fmt.Println("tags:", tags)

	// Map as a "set"
	seen := map[string]bool{}
	for _, p := range []string{"a", "b", "a"} {
		if seen[p] {
			fmt.Println("duplicate:", p)
		}
		seen[p] = true
	}
}
