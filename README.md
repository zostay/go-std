# Zostays Go-Std

There are a lot of common operations simply missing from the Golang builtin
library. This makes up for that deficiency. Here's a summary of some of the
provided tools:

## Generic/Interface Comparison

The built-in comparable generic is pretty weak. Just good enough for map key
types, but not much else. We provide a compare.Able interface which covers all
the types that can be compared using `<`, `>`, `<=`, and `>=` operators, in
addition to `==` and `!=`. Then we provide the following operations on them:

 * `Max(a, b)`
 * `Min(a, b)`
 * `Less(a, b)`
 * `Zero()`

Then we create a generic interface for doing the same thing generally, which
looks similar to what the built-in `sort` package does but then implements the
functions to make such an interface even more useful:

 * `MaxI(a, b)`
 * `MinI(a, b)`
 * `LessI(a, b)`

## Generic Iteration

Iteration in Golang is provided by using a `for`-loop with the `range` keyword.
This has some deficiencies, however. This iterator cannot be passed around
in-progress. Also, when iterating over maps, the iterator is over key/value
pairs rather than index/value pairs. Sometimes it's also helpful to have that
index.

## Generic List Operations for Slices

Many common slice operations are a bit tedious to implement. It's easy to get
off-by-one errors, so the following ops are provided:

 * `Delete(slice, index)`
 * `Push(slice, value)`
 * `Pop(slice)`
 * `Shift(slice)`
 * `Unshift(slice, value)`

## Generic Map Transformations

Frequently, you need to get just the keys or values off of a slice. This
provides those operations among others for maps:

 * `Copy(dst, src)`
 * `CopyInit(dst, src)`
 * `Keys(map)`
 * `Values(map)`
 * `KVs(map)`

## Generic Sets

There's no built-in set operation. This library provides one. It provides the
following methods on sets:

 * `Contains(val)`
 * `Insert(val)`
 * `Delete(val)`
 * `Len()`
 * `SubsetOf(set)`

And then it also provides these functions to work on sets:

 * `Intersects(a, b)`
 * `Intersection(b, b)`
 * `Union(a, b)`
 * `Difference(a, b)`
 * `Copy(dst, src)`
 * `CopyInit(dst, src)`

## String Operations

Some additional string operations are provided that are missing from the
built-in `strings` package:

 * `ContainsOnly(s, chars)`
 * `FromRange(a, z)`
 * `Reverse(s)`
