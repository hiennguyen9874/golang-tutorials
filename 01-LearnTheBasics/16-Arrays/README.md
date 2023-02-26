# Arrays

- Create array
- Go vs C:
  - Arrays are values. Assigning one array to another copies all the elements.
  - In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
  - The size of an array is part of its type. The types [10]int and [20]int are distinct.
- pass a pointer to the array
