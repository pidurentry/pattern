package pattern

import "sync/atomic"

func LoadValue(variables Variables, value interface{}) uint64 {
	switch value := value.(type) {
	case string:
		return atomic.LoadUint64(variables.Fetch(value))
	case uint64:
		return value
	case uint32:
		return uint64(value)
	case uint:
		return uint64(value)
	case int64:
		return uint64(value)
	case int32:
		return uint64(value)
	case int:
		return uint64(value)
	default:
		return 0
	}
}
