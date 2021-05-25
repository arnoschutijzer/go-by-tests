package wraperrors

import (
	"errors"
	"fmt"
)

var ErrValueIs10 = errors.New("oh no, value is 10")

func returnsErrorWhenValueIs10(value int) error {
	if value == 10 {
		return ErrValueIs10
	}

	return nil
}

func ReturnsInnerErrorWhenValueIs10(value int) error {
	err := returnsErrorWhenValueIs10(value)

	if err != nil {
		return fmt.Errorf("something went horribly wrong in the outer function: %w", err)
	}

	return nil
}
