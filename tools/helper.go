package tools

import "sync/atomic"

func Value(value interface{}) uint64 {
	switch value := value.(type) {
	case uint64:
		return value
	case uint32:
		return uint64(value)
	case uint16:
		return uint64(value)
	case uint8:
		return uint64(value)
	case uint:
		return uint64(value)
	case int64:
		return uint64(value)
	case int32:
		return uint64(value)
	case int16:
		return uint64(value)
	case int8:
		return uint64(value)
	case int:
		return uint64(value)
	case float64:
		return uint64(value)
	case float32:
		return uint64(value)
	default:
		return 0
	}
}

func LoadValue(variables Variables, value interface{}) uint64 {
	switch value := value.(type) {
	case string:
		return atomic.LoadUint64(variables.Fetch(value))
	default:
		return Value(value)
	}
}
