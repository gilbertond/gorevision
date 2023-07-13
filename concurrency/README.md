### My concurrency

1. Why Go: has inherent concurrency unlike most languages that use 3rd party libraries
2. Has a robust set of tools to test and build concurrent, parallel, and distributed code
3. Use-cases of channels
    1. Go-routine safe
    2. Implement FIFO behaviour
    3. Exchange data between go-routines
    4. Block and Unblock go-routines
4. Channel types
    1. Buffered (asynchronous)
    2. Unbuffered (synchronous)
    3. Semaphores. Async channels with 0-sized elements. e.g. `chan struct{}`. `nil` buffer and `O(1)` memory.
5. Channel properties
    1. FIFO circular queue behaviour
    2. By default, sends and receives block until the other side is ready. allows to synchronize without explicit locks
    3. Transfer a copy of the object
    4. If routine G1 is to send to a channel but never does, routine G2 reading from the channel blocks indefinitely
6. Channel unbuffered by default
7. Characteristics of unbuffered channels
    1. Sends on this are blocking until another routine reads from it
    2. Reads on this are blocking until another routine sends to it
8. Characteristics of buffered channels
    1. Sends to channel block only when the buffer is full
    2. Receives block when the buffer is empty
    3. If buffer is full or nothing to read, behaves like unbuffered channel 