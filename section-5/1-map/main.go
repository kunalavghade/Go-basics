package main

func main() {
	color := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	println(color["red"])
	println(color["green"])
	println(color["blue"])

	// var numbers map[string]int
	// numbers["one"] = 1 // This will cause a runtime panic because the map is not initialized

	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3

	println(numbers["one"])
	println(numbers["two"])
	println(numbers["three"])

	delete(color, "green")
	println(color["green"]) // This will print an empty string since the key "green" has been deleted
	printMap(color)
}

func printMap(m map[string]string) {
	for key, value := range m {
		println(key, ":", value)
	}
}
