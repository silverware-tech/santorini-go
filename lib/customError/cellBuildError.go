package customError

import "fmt"

type CellBuildError struct {
	ErrorStr string
}

func (error CellBuildError) Error() string {
	return fmt.Sprintf(
		"Can not build in cell: %v",
		error.ErrorStr,
	)
}
