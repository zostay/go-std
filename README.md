# Zostays Go-Std

There are a lot of common operations simply missing from the Golang builtin
library. This makes up for that deficiency. Here's a summary of some of the
provided tools:

# Generic Iteration

Iteration in Golang is provided by using a `for`-loop with the `range` keyword.
This has some deficiencies, however. This iterator cannot be passed around
in-progress. Also, when iterating over maps, the iterator is over key/value
pairs rather than index/value pairs. Sometimes it's also helpful to have that
index.

# Generic List Operations for Slices

Many common slice operations are a bit tedious to implement. It's easy to get
off-by-one errors, so the following ops are provided:

 * Delete(slice, index)
 * Push(slice, value)
 * Pop(slice)
 * Shift(slice)
 * Unshift(slice, value)

# Generic Map Transformations

Frequently, you need to get just the keys or values off of a slice.

 * Keys(map)
 * Values(map)
 * KVs(map)

# Generic Sets

There's no built-in set operation. This library provides one.
