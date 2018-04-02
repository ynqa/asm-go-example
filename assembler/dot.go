package assembler

//go:generate python -m peachpy.x86_64 -S -o avx_dot32.s -mabi=goasm dot.py
func Dot(length uint, x, y *float32) float32
