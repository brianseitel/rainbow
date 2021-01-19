package main

import (
	"fmt"
	"time"

	"github.com/brianseitel/rainbow/color"
	colorful "github.com/lucasb-eyer/go-colorful"
)

func main() {
	start := time.Now()
	c, _ := colorful.Hex("#ff0808")
	name, distance := color.Match(c)
	fmt.Println(time.Since(start))
	fmt.Println(name, 100-(distance*100))

	hex, err := color.Find("pink")
	if err != nil {
		panic(err)
	}
	fmt.Println(hex)
}
