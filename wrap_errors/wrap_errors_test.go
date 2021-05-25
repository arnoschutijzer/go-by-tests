package wraperrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://golang.org/pkg/errors/
func TestWrapErrors(t *testing.T) {
	t.Run("Unwrap with Is", func(t *testing.T) {
		err := ReturnsInnerErrorWhenValueIs10(10)

		// This check will break when the inner message changes!
		assert.Equal(t, "something went horribly wrong in the outer function: oh no, value is 10", err.Error())
		// returns true if the error is in the chain
		assert.True(t, errors.Is(err, ErrValueIs10))
	})

	t.Run("Unwrap with Is with the error not in the chain", func(t *testing.T) {
		err := ReturnsInnerErrorWhenValueIs10(10)

		assert.False(t, errors.Is(err, errors.New("this isn't real")))
	})
}
