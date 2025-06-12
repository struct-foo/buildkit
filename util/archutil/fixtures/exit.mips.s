    .set    noreorder       # disable branch-delay reordering
    .option    big_endian
    .text
    .globl  _start
_start:
    li      $a0, 0         # exit code 0
    li      $v0, 4001      # __NR_exit = 4001 on MIPS
    syscall
