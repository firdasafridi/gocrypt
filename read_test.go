package gocrypt

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	Field1 string `gocrypt:"aes"`
	Field2 string `json:"field2"`
	Field3 string `gocrypt:"des"`
	Field4 int
	Field5 string `gocrypt:"rc4"`
}

type NestedStruct struct {
	Level1 TestStruct
	Level2 *TestStruct
	Level3 string `gocrypt:"aes"`
}

type LargeStruct struct {
	F1  string `gocrypt:"aes"`
	F2  string
	F3  string `gocrypt:"des"`
	F4  int
	F5  string
	F6  string `gocrypt:"rc4"`
	F7  string
	F8  string `gocrypt:"aes"`
	F9  string
	F10 string `gocrypt:"des"`
	F11 int
	F12 string
	F13 string `gocrypt:"rc4"`
	F14 string
	F15 string `gocrypt:"aes"`
}

// Benchmark tests for current optimized implementation
func BenchmarkInspectField_Optimized(b *testing.B) {
	testData := &TestStruct{
		Field1: "test1",
		Field2: "test2",
		Field3: "test3",
		Field4: 123,
		Field5: "test5",
	}

	encDec := func(algo string, text string) (string, error) {
		return "encrypted_" + text, nil
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = inspectField(reflect.ValueOf(testData), encDec)
	}
}

func BenchmarkInspectFieldNested_Optimized(b *testing.B) {
	testData := &NestedStruct{
		Level1: TestStruct{
			Field1: "test1",
			Field2: "test2",
			Field3: "test3",
			Field5: "test5",
		},
		Level2: &TestStruct{
			Field1: "test1",
			Field3: "test3",
		},
		Level3: "nested",
	}

	encDec := func(algo string, text string) (string, error) {
		return "encrypted_" + text, nil
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = inspectField(reflect.ValueOf(testData), encDec)
	}
}

func BenchmarkLargeStruct_Optimized(b *testing.B) {
	testData := &LargeStruct{
		F1:  "field1",
		F3:  "field3",
		F6:  "field6",
		F8:  "field8",
		F10: "field10",
		F13: "field13",
		F15: "field15",
	}

	encDec := func(algo string, text string) (string, error) {
		return "encrypted_" + text, nil
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = inspectField(reflect.ValueOf(testData), encDec)
	}
}

func BenchmarkInspectField_WithCache(b *testing.B) {
	testData := &TestStruct{
		Field1: "test1",
		Field2: "test2",
		Field3: "test3",
		Field4: 123,
		Field5: "test5",
	}

	encDec := func(algo string, text string) (string, error) {
		return "encrypted_" + text, nil
	}

	// Pre-warm cache
	_ = inspectField(reflect.ValueOf(testData), encDec)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = inspectField(reflect.ValueOf(testData), encDec)
	}
}

func BenchmarkLargeStruct_WithCache(b *testing.B) {
	testData := &LargeStruct{
		F1:  "field1",
		F3:  "field3",
		F6:  "field6",
		F8:  "field8",
		F10: "field10",
		F13: "field13",
		F15: "field15",
	}

	encDec := func(algo string, text string) (string, error) {
		return "encrypted_" + text, nil
	}

	// Pre-warm cache
	_ = inspectField(reflect.ValueOf(testData), encDec)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = inspectField(reflect.ValueOf(testData), encDec)
	}
}
