package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	vi := viper.New()
	vi.SetConfigFile("test.yaml")
	vi.ReadInConfig()

	// Using https://github.com/spf13/viper#getting-values-from-viper
	fmt.Println("Is name set? ", vi.IsSet("name"))
	fmt.Println("Is id set? ", vi.IsSet("id"))
	fmt.Println("Is inputs set? ", vi.IsSet("inputs"))
	fmt.Println("Is api set? ", vi.IsSet("api"))
	fmt.Println("----- First Level -----")
	fmt.Println("Name Value:", vi.GetString("name"))
	fmt.Println("Id Value:", vi.GetInt("id"))
	fmt.Println("Inputs Map:", vi.Get("inputs"))
	fmt.Println("----- Second Level -----")
	fmt.Println("Inputs-Request1:", vi.Get("inputs.request1"))
	fmt.Println("----- Third Level -----")
	fmt.Println("Inputs-Request1-Method Value:", vi.Get("inputs.request1.method"))
	fmt.Println("Inputs-Request1-Url Value:", vi.Get("inputs.request1.url"))
	fmt.Println("Inputs-Request1-Headers Map:", vi.Get("inputs.request1.headers"))
	fmt.Println("Inputs-Request1-Body Map:", vi.Get("inputs.request1.body"))
	fmt.Println("----- Second Level -----")
	fmt.Println("Request2:", vi.Get("inputs.request2"))
	fmt.Println("----- Third Level -----")
	fmt.Println("Inputs-Request2-Method Value:", vi.Get("inputs.request2.method"))
	fmt.Println("Inputs-Request2-Url Value:", vi.Get("inputs.request2.url"))
	fmt.Println("Inputs-Request2-Headers Map:", vi.Get("inputs.request2.headers"))
	fmt.Println("Inputs-Request2-Body Map:", vi.Get("inputs.request2.body"))
	fmt.Println("-----")
	fmt.Println("All:", vi.AllSettings())
}
