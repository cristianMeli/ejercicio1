package main

import(
	"net/http"
	"testing"

)

const baseUrl = "http://localhost:8080/"

func BenchmarkResponse(b *testing.B) {

	for n := 0; n < b.N; n++ {
		http.Get(baseUrl + "response/2/mock")
	}
}

func BenchmarkResponseWg(b *testing.B) {

	for n := 0; n < b.N; n++ {
		http.Get(baseUrl + "responseWg/2/mock")
	}
}

func BenchmarlResponseCh(b *testing.B) {

	for n := 0; n < b.N; n++ {
		http.Get(baseUrl + "responseCh/2/mock")
	}
}
