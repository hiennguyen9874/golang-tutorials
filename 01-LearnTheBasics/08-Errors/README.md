# Errors

- errors are the last return value and have type error, a built-in interface
- errors.New constructs a basic error value with the given error message
- A nil value in the error position indicates that there was no error.
- custom types as errors by implementing the Error() method
