package maputil

// Returns a new map with all elements from input map that satisfy the condition specified in the function.
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter_ReturnsNewMapWithSatisfyingElements(t *testing.T) {
	// Arrange
	inputMap := map[string]int{"a": 1, "b": 2, "c": 3}
	expectedMap := map[string]int{"b": 2, "c": 3}

	// Act
	result := Filter(inputMap, func(k string, v int) bool {
		return v > 1
	})

	// Assert
	assert.Equal(t, expectedMap, result)
}

func TestFilter_ReturnsNewMapWithOneElement(t *testing.T) {
	// Arrange
	inputMap := map[string]int{"a": 1}
	expectedMap := map[string]int{"a": 1}

	// Act
	result := Filter(inputMap, func(k string, v int) bool {
		return v == 1
	})

	// Assert
	assert.Equal(t, expectedMap, result)
}

func TestMapSameLength(t *testing.T) {
	input := map[int]string{1: "one", 2: "two", 3: "three"}
	expectedLength := len(input)

	result := Map(input, func(k int, v string) (int, string) {
		return k, v
	})

	if len(result) != expectedLength {
		t.Errorf("Expected map length %d, but got %d", expectedLength, len(result))
	}
}

func TestMapEmptyInput(t *testing.T) {
	input := map[int]string{}

	result := Map(input, func(k int, v string) (int, string) {
		return k, v
	})

	if len(result) != 0 {
		t.Errorf("Expected empty map, but got map with length %d", len(result))
	}
}

func TestReduceEmptyMap(t *testing.T) {
	// Arrange
	m := make(map[int]string)
	init := "initial"
	f := func(k int, v string, acc string) string {
		return acc + v
	}

	// Act
	result := Reduce(m, init, f)

	// Assert
	assert.Equal(t, init, result)
}

func TestReduceNilFunction(t *testing.T) {
	// Arrange
	m := map[int]string{1: "one", 2: "two"}
	init := ""
	var f func(int, string, string) string

	// Act & Assert
	assert.Panics(t, func() { Reduce(m, init, f) })
}

func TestInvertEmptyMap(t *testing.T) {
	input := make(map[int]string)
	expected := make(map[string]int)

	result := Invert(input)

	assert.Equal(t, expected, result)
}

func TestInvertNilMap(t *testing.T) {
	var input map[int]string
	expected := make(map[string]int)

	result := Invert(input)

	assert.Equal(t, expected, result)
}
