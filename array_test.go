package logus

import (
	"testing"
	"time"

	"go.uber.org/zap/zapcore"

	"github.com/stretchr/testify/assert"
)

func BenchmarkBoolsArrayMarshaler(b *testing.B) {
	// Keep this benchmark here to capture the overhead of the ArrayMarshaler
	// wrapper.
	bs := make([]bool, 50)
	enc := zapcore.NewJSONEncoder(zapcore.EncoderConfig{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bools("array", bs...).AddTo(enc.Clone())
	}
}

func BenchmarkBoolsReflect(b *testing.B) {
	bs := make([]bool, 50)
	enc := zapcore.NewJSONEncoder(zapcore.EncoderConfig{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reflect("array", bs).AddTo(enc.Clone())
	}
}

func TestArrayWrappers(t *testing.T) {
	tests := []struct {
		desc     string
		field    Field
		expected []interface{}
	}{
		{"empty bools", Bools(""), []interface{}{}},
		{"empty byte strings", ByteStrings(""), []interface{}{}},
		{"empty complex128s", Complex128s(""), []interface{}{}},
		{"empty complex64s", Complex64s(""), []interface{}{}},
		{"empty durations", Durations(""), []interface{}{}},
		{"empty float64s", Float64s(""), []interface{}{}},
		{"empty float32s", Float32s(""), []interface{}{}},
		{"empty ints", Ints(""), []interface{}{}},
		{"empty int64s", Int64s(""), []interface{}{}},
		{"empty int32s", Int32s(""), []interface{}{}},
		{"empty int16s", Int16s(""), []interface{}{}},
		{"empty int8s", Int8s(""), []interface{}{}},
		{"empty strings", Strings(""), []interface{}{}},
		{"empty times", Times(""), []interface{}{}},
		{"empty uints", Uints(""), []interface{}{}},
		{"empty uint64s", Uint64s(""), []interface{}{}},
		{"empty uint32s", Uint32s(""), []interface{}{}},
		{"empty uint16s", Uint16s(""), []interface{}{}},
		{"empty uint8s", Uint8s(""), []interface{}{}},
		{"empty uintptrs", Uintptrs(""), []interface{}{}},
		{"bools", Bools("", true, false), []interface{}{true, false}},
		{"byte strings", ByteStrings("", [][]byte{{1, 2}, {3, 4}}...), []interface{}{[]byte{1, 2}, []byte{3, 4}}},
		{"complex128s", Complex128s("", []complex128{1 + 2i, 3 + 4i}...), []interface{}{1 + 2i, 3 + 4i}},
		{"complex64s", Complex64s("", []complex64{1 + 2i, 3 + 4i}...), []interface{}{complex64(1 + 2i), complex64(3 + 4i)}},
		{"durations", Durations("", []time.Duration{1, 2}...), []interface{}{time.Nanosecond, 2 * time.Nanosecond}},
		{"float64s", Float64s("", 1.2, 3.4), []interface{}{1.2, 3.4}},
		{"float32s", Float32s("", 1.2, 3.4), []interface{}{float32(1.2), float32(3.4)}},
		{"ints", Ints("", 1, 2), []interface{}{1, 2}},
		{"int64s", Int64s("", 1, 2), []interface{}{int64(1), int64(2)}},
		{"int32s", Int32s("", 1, 2), []interface{}{int32(1), int32(2)}},
		{"int16s", Int16s("", 1, 2), []interface{}{int16(1), int16(2)}},
		{"int8s", Int8s("", 1, 2), []interface{}{int8(1), int8(2)}},
		{"strings", Strings("", "foo", "bar"), []interface{}{"foo", "bar"}},
		{"times", Times("", time.Unix(0, 0), time.Unix(0, 0)), []interface{}{time.Unix(0, 0), time.Unix(0, 0)}},
		{"uints", Uints("", 1, 2), []interface{}{uint(1), uint(2)}},
		{"uint64s", Uint64s("", 1, 2), []interface{}{uint64(1), uint64(2)}},
		{"uint32s", Uint32s("", 1, 2), []interface{}{uint32(1), uint32(2)}},
		{"uint16s", Uint16s("", 1, 2), []interface{}{uint16(1), uint16(2)}},
		{"uint8s", Uint8s("", 1, 2), []interface{}{uint8(1), uint8(2)}},
		{"uintptrs", Uintptrs("", 1, 2), []interface{}{uintptr(1), uintptr(2)}},
	}

	for _, tt := range tests {
		enc := zapcore.NewMapObjectEncoder()
		tt.field.Key = "k"
		tt.field.AddTo(enc)
		assert.Equal(t, tt.expected, enc.Fields["k"], "%s: unexpected map contents.", tt.desc)
		assert.Equal(t, 1, len(enc.Fields), "%s: found extra keys in map: %v", tt.desc, enc.Fields)
	}
}
