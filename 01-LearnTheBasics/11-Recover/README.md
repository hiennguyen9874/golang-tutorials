# Recover

- Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
- Stop panic from aborting the program and let it continue with execution instead
- `recover` must be called within a deferred function
