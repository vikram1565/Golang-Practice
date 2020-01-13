1) Slice
2) Map
3) packages
4) Interface
   1) an interface is a type that specifies a method set and all the methods for an interface type are considered to      be the interface.
   2) 
5) Methods
   1) function with receiver value or struct 
6) pointers
7) Embedding
8) Concurrency
9)  Goroutine with internal working
10) WaitGroup
11) Scheduling
12) Memory allocation (stack and heap)
13) defer with internal working
    1)  When a new defer method is created, it is attached to the current Goroutine and the previous one is linked to      the new one as the next function to be executed
    2)  function returns then executes defer
    3)  working as LIFO
    4)  _defer struct with 7 fields
    5)  newDefer create above struct for each defer func call
    6)  runtime.deferproc - to execute current func
    7)  runtime.deferreturn - return current proc result
    8)  runtime.gopanic - call panic function
14) Panic , Recover
15) Garbage collection
16) channel
17) Testing (Unit , Benchmark , Mock)
18) Echo , Gin 
19) Context
20) Struct tags , omitempty
21) Error handling
22) Unsafe Pointer
23) Profiling
24) Escape Analysis
25) Worker pool
26) Data races
27) Reflect
28) Pragmas
29) GRPC
30) 