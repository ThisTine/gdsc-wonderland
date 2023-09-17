package value

func Ptr[T any](v T) *T {
	return &v
}

var TruePtr = Ptr(true)
var FalsePtr = Ptr(false)
