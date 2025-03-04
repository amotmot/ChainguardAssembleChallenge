.text
.section .text.startup,"ax",@progbits
.p2align 4
.globl main
.type main, @function
main:
.LFB11:
.cfi_startproc
    subq $24, %rsp       
    .cfi_def_cfa_offset 32
    movdqa LC1(%rip), %xmm0
    movq %rsp, %rdi        
    movaps %xmm0, (%rsp)   
    movdqa %xmm0, %xmm1    
    movq .LCO(%rip), %xmm0 
    xorl $33686018, 8(%rsp) 
    pxor %xmm1, %xmm0       
    xorw $514, 12(%rsp)     
    movb $53, 14(%rsp)      
    movq %xmm0, (%rsp)      
    call puts               
    xorl %eax, %eax  
    addq $24, %rsp   
    .cfi_def_cfa_offset 8
    ret              
.cfi_endproc
.LFE11:
.size main,.-main
.section .rodata.cst8,"aM",@progbits,8
.align 8
.LCO:
    .byte 2          
    .byte 2
    .byte 2
    .byte 2
    .byte 2
    .byte 2
    .byte 2
    .byte 2
.section .rodata.cst16,"aM",@progbits, 16
.align 16
LC1:
    .quad 8388875062886221384     
    .quad 15534430226834279
.section .note.GNU-stack,"",@progbits