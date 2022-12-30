package customError
import "fmt"

type PlayerMoveError struct {
	PlayerName	string
	PlayerX		int
	PlayerY		int
	ErrorStr	string
}

func (error PlayerMoveError) Error() string {
	return fmt.Sprintf(
		"Can not move player \"%v\" to position (%v, %v): %v",
		error.PlayerName,
		error.PlayerX,
		error.PlayerY,
		error.ErrorStr)
}