# NOTES
notes i use to keep track of what im thinking/ doing during this course.

- industry standard to read an IO file descriptor == 4(4096) or 8(8192) KB(kilo bytes) per chunk
    - this is to match page/ buffer sizes
    - not 8 bytes!! too slow; inefficient
    - causes excessive read() syscalls => need to context switch from user to kernel modes

-  use 1 goroutine to handle 1 request at each time