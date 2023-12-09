package dog

import (
	"testing"

	"github.com/nilavanrajamani/golang-gettingstarted/254_bet/dog"
)

func BenchmarkYears(b *testing.B) {
	// TODO: Initialize
	for i := 0; i < b.N; i++ {
		dog.Years(3)
	}
}
