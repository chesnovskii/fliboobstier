package bot_helpers

import (
	"errors"
	"strings"
)

func ErrListToError(err_list []string) error {
	non_empty_err_list := []string{}
	for _, err := range err_list {
		if err != "" {
			non_empty_err_list = append(non_empty_err_list, err)
		}
	}
	if len(non_empty_err_list) == 0 {
		return nil
	}
	return errors.New(strings.Join(non_empty_err_list, "\n"))
}
