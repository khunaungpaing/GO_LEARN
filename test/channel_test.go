package main

import (
	"testing"

	lvl2 "github.com/khunaungpaing/GO_LEARN/level-two"
)

func BenchmarkChannelInGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lvl2.ChannelInGo()
	}
}
