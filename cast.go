package cast

import (
	"fmt"
	"strconv"
)

//String cast any type to string
func String(i interface{}) (string, error) {
	switch i := i.(type) {
	case string:
		return i, nil
	case int:
		return strconv.Itoa(i), nil
	case float64:
		return fmt.Sprintf(`%f`, i), nil
	case []byte:
		return string(i), nil
	}
	return ``, nil
}

//Int cast any type to int
func Int(i interface{}) (int, error) {
	switch i := i.(type) {
	case string:
		return strconv.Atoi(i)
	case int:
		return i, nil
	case float64:
		return int(i), nil
	}
	return 0, nil
}

//Float64 cast any type to float64
func Float64(i interface{}) (float64, error) {
	switch i := i.(type) {
	case string:
		return strconv.ParseFloat(i, 64)
	case int:
		return float64(i), nil
	case float64:
		return i, nil
	}
	return 0.0, nil
}

//Bytes cast any type to []byte
func Bytes(i interface{}) ([]byte, error) {
	switch i := i.(type) {
	case string:
		return []byte(i), nil
	case []byte:
		return i, nil
	}
	return nil, nil
}
