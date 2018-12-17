package logus

import (
	"go.uber.org/zap"
	"time"
)

func ExampleDebug() {
	Debug("Debug level message", zap.Time("time", time.Now().Local()))
	// Output:
}

func ExampleInfo() {
	Info("Info level message", zap.Time("time", time.Now().Local()))
	// Output:
}

func ExampleWarn() {
	Warn("Warn level message", zap.Time("time", time.Now().Local()))
	// Output:
}

func ExampleError() {
	Error("Error level message", zap.Time("time", time.Now().Local()))
	// Output:
}

func ExampleDPanic() {
	defer func() {
		recover()
	}()

	DPanic("DPanic level message", zap.Time("time", time.Now().Local()))
	// Output:
}

func ExamplePanic() {
	defer func() {
		recover()
	}()

	Panic("Panic level message", zap.Time("time", time.Now().Local()))
	// Output:
}

func ExampleFatal() {
	Error("Fatal level message", zap.Time("time", time.Now().Local()))
	// Output:
}
