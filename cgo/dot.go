package cgo

/*
#cgo CFLAGS: -mavx -std=c99
#include <stdio.h>
#include <stdlib.h>
#include <immintrin.h>//AVX: -mavx

float avx_dot32(const size_t n, const float* x, const float* y) {
	static const size_t single_precision_size = 8;
	const size_t end = n / single_precision_size;

	__m256 res = _mm256_setzero_ps();

	__m256 *vx = (__m256 *)x;
	__m256 *vy = (__m256 *)y;
	
	for(size_t i=0; i<end; i++) {
		__m256 tmp = _mm256_mul_ps(vx[i], vy[i]);
		res = _mm256_add_ps(res, tmp);
	}	
	return res[0] + res[1] + res[2] + res[3] +
		res[4] + res[5] + res[6] + res[7];
}
*/
import "C"

import (
	"reflect"
	"unsafe"
)

const align = 32

func Malloc32(length int) []float32 {
	size := length * align
	ptr := C._mm_malloc((C.size_t)(size), align)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ptr)),
		Len:  size,
		Cap:  size,
	}
	goSlice := *(*[]float32)(unsafe.Pointer(&hdr))
	return goSlice
}

func Free32(v []float32) {
	C._mm_free(unsafe.Pointer(&v[0]))
}

func Dot(length int, x, y []float32) float32 {
	return float32(C.avx_dot32((C.size_t)(length), (*C.float)(&x[0]), (*C.float)(&y[0])))
}
