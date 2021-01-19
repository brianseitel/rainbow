package color_test

import (
	"testing"

	"github.com/brianseitel/rainbow/color"
	colorful "github.com/lucasb-eyer/go-colorful"
)

func BenchmarkMatch(b *testing.B) {

	for i := 0; i < b.N; i++ {
		hex := colorful.FastHappyColor()
		c, d := color.Match(hex)
		_ = c
		_ = d
	}
}

func BenchmarkCluster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hex := colorful.FastHappyColor()
		c, d := color.Cluster(hex)
		_ = c
		_ = d
	}
}
