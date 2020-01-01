package main

import "fmt"

func main() {
	// var a map[string]int = make(map[string]int)
	// var b = make(map[string]int)
	// c := make(map[string]int)

	d := map[string]int{"Hello": 10, "world": 20}

	e := map[string]int{
		"Hello": 10,
		"world": 20,
	}

	fmt.Println(d["Hello"])
	fmt.Println(e["Hello"])
	fmt.Println()

	solarSystem := make(map[string]float32)

	solarSystem["Mercury"] = 87.969
	solarSystem["Venus"] = 224.70069
	solarSystem["Earth"] = 365.25641
	solarSystem["Mars"] = 686.9600
	solarSystem["Jupiter"] = 4333.2867
	solarSystem["Saturn"] = 10756.1995
	solarSystem["Uranus"] = 30707.4896
	solarSystem["Neptune"] = 60223.3528

	fmt.Println(solarSystem["Earth"])
	fmt.Println(solarSystem["Neptune"])
	fmt.Println(solarSystem["Pluto"]) // 0: 존재하지 않는 키를 조회
	fmt.Println()

	value, ok := solarSystem["Pluto"]
	fmt.Println(value, ok)

	if value, ok := solarSystem["Neptune"]; ok {
		fmt.Println(value)
	}

	for key, value := range solarSystem {
		fmt.Println(key, value)
	}

	for _, value := range solarSystem {
		fmt.Println(value)
	}

	z := map[string]int{"hello": 10, "world": 20}
	delete(z, "world")
	fmt.Println(z)

	terrestrialPlanet := map[string]map[string]float32{
		"Mercury": map[string]float32{
			"meanRadius":    2439.7,
			"mass":          3.3022E+23,
			"orbitalPeriod": 87.969,
		},
		"Venus": map[string]float32{
			"meanRadius":    6051.8,
			"mass":          4.8676E+24,
			"orbitalPeriod": 224.70069,
		},
		"Earth": map[string]float32{
			"meanRadius":    6371.0,
			"mass":          5.97219E+24,
			"orbitalPeriod": 365.25641,
		},
		"Mars": map[string]float32{
			"meanRadius":    3389.5,
			"mass":          6.4185E+23,
			"orbitalPeriod": 686.9600,
		},
	}

	fmt.Println(terrestrialPlanet["Mars"]["mass"])
}
