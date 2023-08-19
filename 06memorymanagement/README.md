# Memory Management in GO

memory allocation and de-allocation happens automatically

## Memory Management

- new()
    - Allocate Memory but No INIT
    - you will get a memory address
    - zeroed storage  (cannot put data initially)

- make()
    - Allocate memory and INIT
    - you will get a memory address
    - non zeroed storage (can put data initially)

GC happens Automatically

out of scope or nil

more information for the same can be found at https://pkg.go.dev/runtime#pkg-functions

> The GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage. The default is GOGC=100. Setting GOGC=off disables the garbage collector entirely. runtime/debug.SetGCPercent allows changing this percentage at run time. 