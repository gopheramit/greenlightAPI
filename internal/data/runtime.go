package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJsonValue := strconv.Quote(jsonValue)
	return []byte(quotedJsonValue), nil

}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquptedJsonValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		fmt.Println("I am here")
		return ErrInvalidRuntimeFormat
	}
	parts := strings.Split(unquptedJsonValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		fmt.Println("I am here2")
		return ErrInvalidRuntimeFormat
	}
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		fmt.Println("I am here3")
		return ErrInvalidRuntimeFormat
	}
	*r = Runtime(i)
	return nil
}
