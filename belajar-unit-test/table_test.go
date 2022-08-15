package belajarunittest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("======Before")

	m.Run()

	fmt.Println("======After")
}

func TestTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Halo(Eko)",
			request:  "Eko",
			expected: "Halo Eko",
		},
		{
			name:     "Halo(Jhon)",
			request:  "Jhon",
			expected: "Halo Jhon",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Halo(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func Halo(param string) string {
	return "Halo " + param
}
