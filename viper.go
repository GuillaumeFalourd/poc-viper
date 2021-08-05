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
	fmt.Println(vi.GetString("name"))
	fmt.Println(vi.GetInt("id"))
	fmt.Println(vi.Get("inputs"))
	fmt.Println(vi.AllSettings())
}
