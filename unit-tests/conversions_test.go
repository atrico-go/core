package unit_tests

import (
	"fmt"
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

	"github.com/atrico-go/testing/random"

	"github.com/atrico-go/core"
)

var rg = random.NewValueGenerator()

func Test_Conversion_ConvertSlice_StringToInterface(t *testing.T) {
	// Arrange
	var output []interface{}
	var input []string
	rg.Value(&input)

	// Act
	err := core.ConvertSlice(input, &output)

	// Assert
	Assert(t).That(err, is.Nil, "No error")
	for i, v := range input {
		Assert(t).That(fmt.Sprintf("%v", output[i]), is.EqualTo(v), fmt.Sprintf("Item %d", i))
	}
}

func Test_Conversion_ConvertSlice_InterfaceToInt(t *testing.T) {
	// Arrange
	var output []int
	input := []interface{} {
		rg.Int(),
		rg.Int(),
		rg.Int(),
	}

	// Act
	err := core.ConvertSlice(input, &output)

	// Assert
	Assert(t).That(err, is.Nil, "No error")
	for i, v := range input {
		Assert(t).That(output[i], is.EqualTo(v), fmt.Sprintf("Item %d", i))
	}
}

func Test_Conversion_ConvertSlice_Error(t *testing.T) {
	// Arrange
	var output []int
	var input []string
	rg.Value(&input)

	// Act
	err := core.ConvertSlice(input, &output)

	// Assert
	fmt.Println(err)
	Assert(t).That(err, is.NotNil, "error")
}
