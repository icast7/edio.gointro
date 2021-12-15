package main

import (
	"fmt"
	"reflect"
)

func main() {
	p := []int{1, 2, 3, 5, 7, 11, 13}
	theType := reflect.TypeOf(p)
	fmt.Println(theType.Kind())
	fmt.Printf("%+v\n", p)
	fmt.Println(p[1:4])
	fmt.Println(p[:3]) // missing low = 0
	fmt.Println(p[4:]) // missing high = len(x)
	fmt.Println("- - - - - - - - - - - - - - - - - - - -")

	names := [4]string{"Lennon", "McCartney", "Harrison", "Starr"}
	fmt.Println(reflect.TypeOf(names).Kind())
	fmt.Println(names)

	a := names[0:2]
	fmt.Println(reflect.TypeOf(a).Kind())
	fmt.Println(a)

	b := names[1:3]
	fmt.Println(reflect.TypeOf(b).Kind())
	fmt.Println(b)
	fmt.Println("a:", a, "b:", b)
	// Demonstrates value has been changed by pointers
	b[0] = "XXX"
	fmt.Println("a:", a, "b:", b)
	fmt.Println("original array:", names)

	// Making slices
	cities := make([]string, 3)
	cities[0] = "Santa Mónica"
	cities[1] = "Venice"
	cities[2] = "Los Ángeles"
	fmt.Printf("%q\n", cities)

	// Empty slice
	moreCities := []string{}
	// moreCities[0] = "Calabasas" // can't do
	moreCities = append(moreCities, "Calabasas")
	fmt.Printf("%q\n", moreCities)

	moreCities = append(moreCities, "EastLos", "Pasadena")
	fmt.Printf("%q\n", moreCities)

	citiez := []string{"San Diego", "Mountain View"}
	moreCitiez := []string{"Escondido", "Encinitas"}
	citiez = append(citiez, moreCitiez...)
	fmt.Printf("%q\n", citiez)
}
