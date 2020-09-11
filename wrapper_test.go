package textwrapper_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/emersion/go-textwrapper"
)

func TestNew(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{
			input:    []string{"helloworldhelloworldhelloworld"},
			expected: "hello/world/hello/world/hello/world",
		},
		{
			input:    []string{"helloworldhelloworldhe"},
			expected: "hello/world/hello/world/he",
		},
		{
			input:    []string{"helloworldhelloworldhe", "ll"},
			expected: "hello/world/hello/world/hell",
		},
		{
			input:    []string{"helloworldhelloworldhe", "llo"},
			expected: "hello/world/hello/world/hello",
		},
		{
			input:    []string{"helloworldhelloworldhe", "lloworld"},
			expected: "hello/world/hello/world/hello/world",
		},
		{
			input:    []string{"helloworldhelloworldhe", "llo", "wor", "ld"},
			expected: "hello/world/hello/world/hello/world",
		},
	}

	for _, test := range tests {
		var b bytes.Buffer
		w := textwrapper.New(&b, "/", 5)

		for _, i := range test.input {
			w.Write([]byte(i))
		}

		output := b.String()
		if output != test.expected {
			t.Error("Got " + output + " instead of " + test.expected)
		}
	}
}

func BenchmarkWriteWith10K(b *testing.B) {
	b.ReportAllocs()

	input := make([]byte, 10000)
	w := textwrapper.New(ioutil.Discard, "/", 3)
	for i := 0; i < b.N; i++ {
		_, err := w.Write(input)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkWriteWith100K(b *testing.B) {
	b.ReportAllocs()

	input := make([]byte, 100000)
	w := textwrapper.New(ioutil.Discard, "/", 3)
	for i := 0; i < b.N; i++ {
		_, err := w.Write(input)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkWriteWith1M(b *testing.B) {
	b.ReportAllocs()

	input := make([]byte, 1000000)
	w := textwrapper.New(ioutil.Discard, "/", 3)
	for i := 0; i < b.N; i++ {
		_, err := w.Write(input)
		if err != nil {
			b.Error(err)
		}
	}
}
