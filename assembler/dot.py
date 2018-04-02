from peachpy import *
from peachpy.x86_64 import *

x = Argument(ptr(const_float_), name='x')
y = Argument(ptr(const_float_), name='y')

# python -m peachpy.x86_64 -mabi=goasm -S -o avx_dot32.s dot.py
with Function('Dot', (x, y), float_, target=uarch.default + isa.avx):
    reg_x = GeneralPurposeRegister64()
    reg_y = GeneralPurposeRegister64()
    LOAD.ARGUMENT(reg_x, x)
    LOAD.ARGUMENT(reg_y, y)
    VMOVAPS(xmm0, [reg_x])
    VMOVAPS(xmm1, [reg_y])
    VDPPS(xmm2, xmm0, xmm1, 0xFF)
    RETURN(xmm2)
