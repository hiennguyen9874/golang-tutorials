# For loop

- Like a C for: `for init; condition; post { }`
- Like a C while: `for condition { }`
- Like a C for(;;): `for { }`
- If you're looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop:
  ```go
  for key, value := range oldMap {
      newMap[key] = value
  }
  ```
- If you only need the first item in the range (the key or index), drop the second:
  ```go
  for key := range m {
      if key.expired() {
          delete(m, key)
      }
  }
  ```
- If you only need the second item in the range (the value), use the blank identifier, an underscore, to discard the first:
  ```go
  sum := 0
  for _, value := range array {
      sum += value
  }
  ```
- `continue, break`
