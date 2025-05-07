package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	result := 4

	if sum != result {
		t.Errorf("the sum is = '%d', but it was expected = '%d'", sum, result)
	}
}

func ExampleAdd() {
	sum := Add(5, 5)
	fmt.Println(sum)
	// Output: 10
}
