# Slices

- Slices are similar to arrays, but are more powerful and flexible.
- Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array
- If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller, analogous to passing a pointer to the underlying array
- len, cap