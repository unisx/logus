package logus

import (
	"time"
)

func ExampleDebug() {
	Debug("Debug level message",
		Time("time", time.Now().Local()),
		Ints("Ints", 1, 2),
		Any("Any", []int{1, 2}))
	// Output:
}

func ExampleInfo() {
	Info("Info level message", Time("time", time.Now().Local()))
	// Output:
}

func ExampleWarn() {
	Warn("Warn level message", Time("time", time.Now().Local()))
	// Output:
}

func ExampleError() {
	Error("Error level message", Time("time", time.Now().Local()))
	// Output:
}

func ExampleDPanic() {
	defer func() {
		recover()
	}()

	DPanic("DPanic level message", Time("time", time.Now().Local()))
	// Output:
}

func ExamplePanic() {
	defer func() {
		recover()
	}()

	Panic("Panic level message", Time("time", time.Now().Local()))
	// Output:
}

func ExampleFatal() {
	Error("Fatal level message", Time("time", time.Now().Local()))
	// Output:
}
