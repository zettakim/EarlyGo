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
}
