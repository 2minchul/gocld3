//go:generate protoc --cpp_out=. feature_extractor.proto sentence.proto task_spec.proto

// Package cld3 implements language detection using the Compact Language Detector v3.
//
// This packages includes the relevant sources from the CLD3 project, so it doesn't require any external dependencies. For more information on CLD3, see https://github.com/google/cld3/ .
package cld3

// #cgo CXXFLAGS: -std=c++11
// #cgo pkg-config: protobuf
// #include <stdlib.h>
// #include "cld3.h"
import "C"
import "unsafe"

const UnknownLang = "und"

// FindLanguageOfValidUTF8 detects the language in a given text. The Result's
// Language will be "und" if it is unknown.
func FindLanguage(text string) Result {
	cs := C.CString(text)
	defer C.free(unsafe.Pointer(cs))
	res := C.FindLanguage(cs, C.int(len(text)))
	r := Result{}
	r.Language = C.GoStringN(res.language, res.len_language)
	r.Probability = float32(res.probability)
	r.IsReliable = bool(res.is_reliable)
	r.Proportion = float32(res.proportion)
	return r
}

type Result struct {
	Language string

	// Probability is the probability from 0 to 1 of the text being in the
	// returned Language.
	Probability float32

	// IsReliable is true when the prediction is reliable.
	IsReliable bool

	// Proportion of bytes associated with the language. If FindLanguage is
	// called, this variable is set to 1.
	Proportion float32
}
