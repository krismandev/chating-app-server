package utility

import (
	"errors"

	"gorm.io/gorm"
)

func CheckErrorResult(result *gorm.DB) error {
	var err error
	err = result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &NotFoundError{Code: 404, Message: "Data not Found"}
		} else {
			return err
		}
	}
	return err
}
