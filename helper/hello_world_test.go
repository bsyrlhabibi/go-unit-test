package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Habibi",
			request: "Habibi",
		},
		{
			name:    "Basyarul",
			request: "Basyarul",
		},
	}
	for _, benchmarks := range benchmarks {
		b.Run(benchmarks.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmarks.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Habibi", func(b *testing.B) {
		result := HelloWorld("Habibi")
		require.Equal(b, "Hello Habibi", result, "Result must be 'Hello Habibi'")
	})
	b.Run("Basyarul", func(B *testing.B) {
		result := HelloWorld("Basyarul")
		require.Equal(b, "Hello Basyarul", result, "Result must be 'Hello Basyarul'")
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Habibi")
	}
}

func BenchmarkHelloWorldBaysarul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Basyarul")
	}
}

// table test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Muhammadun",
			request:  "Muhammadun",
			expected: "Hello Muhammadun",
		},
		{
			name:     "Basyarul",
			request:  "Basyarul",
			expected: "Hello Basyarul",
		},
		{
			name:     "Habibi",
			request:  "Habibi",
			expected: "Hello Habibi",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Habibi", func(t *testing.T) {
		result := HelloWorld("Habibi")
		require.Equal(t, "Hello Habibi", result, "Result must be 'Hello Habibi'")
	})
	t.Run("Basyarul", func(t *testing.T) {
		result := HelloWorld("Basyarul")
		require.Equal(t, "Hello Basyarul", result, "Result must be 'Hello Basyarul'")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on Windows OS")
	}
	result := HelloWorld("Assert")
	assert.Equal(t, "Hello Assert", result, "Result must be 'Hello Assert'")
}

// next nya lebih baik menggunakan Assertions tidak manual menggunakan if else
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Assert")
	assert.Equal(t, "Hello Assert", result, "Result must be 'Hello Assert'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Require")
	require.Equal(t, "Hello Require", result, "Result must be 'Hello Require'")
	fmt.Println("TestHelloWorld with Require Done")
}

// manual if-else
func TestHelloWorldHabibi(t *testing.T) {
	result := HelloWorld("Habibi")
	if result != "Hello Habibi" {
		//unit test failed
		t.Error("Result must be 'Hello Habibi'")
	}
	fmt.Println("TestHelloWorldHabibi Done")
}

func TestHelloWorldBibi(t *testing.T) {
	result := HelloWorld("Bibi")
	if result != "Hello Bibi" {
		//unit test failed
		// panic("Result is not 'Hello Bibi'") // jangan menggunakan panic untuk unit test

		// disarankan lebih baik menggunakan t.Error() atau t.Fatal() supaya bisa mendapatkan informasi Unit Test gagal
		// t.Fail()
		// t.FailNow()

		t.Fatal("Result must be 'Hello Bibi'")
	}
	fmt.Println("TestHelloWorldHabibi Done")
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Sebelum Unit Test")

	m.Run()

	// after
	fmt.Println("Sesudah Unit Test")
}
