package main

import "testing"

func BenchmarkSlugify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = slugify(`Go Talks: "Cuddle: an App Engine Demo":`)
	}
}
