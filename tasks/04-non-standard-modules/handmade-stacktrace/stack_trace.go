package stacktrace

const maxStacktraceDepth = 32

type Frame uintptr

func (f Frame) pc() uintptr {
	return uintptr(f) - 1
}

func (f Frame) String() string {
	// Реализуй меня.
	return ""
}

type StackTrace []Frame

func (s StackTrace) String() string {
	// Реализуй меня.
	return ""
}

// Trace возвращает стектрейс глубиной не более maxStacktraceDepth.
// Возвращаемый стектрейс начинается с того места, где была вызвана Trace.
func Trace() StackTrace {
	// Реализуй меня.
	return nil
}
