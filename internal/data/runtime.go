package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%dmins", r)
	quotedJsonValue := strconv.Quote(jsonValue)
	return []byte(quotedJsonValue), nil

}
