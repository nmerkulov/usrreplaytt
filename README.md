Write an application that reads random integers, in the range of 1 to 10000, from a channel,
and distributes those integers to multiple go-routines.
Each go-routine should sleep for the millisecond duration matching the integer;
so if the goroutine receives a value of 100; it will sleep for 100 milliseconds.
The number of go routines should be a constant, such that it can be easily changed during our testing;
and the number of integers that will be generated should be a constant value too.
The application should wait for all of the go-routines to finish processing before existing.