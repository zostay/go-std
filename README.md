# Zostays Go-Std

There are a lot of common operations simply missing from the Golang builtin
library. This makes up for that deficiency. Here's a summary of some of the
provided tools:

## Byte Slice Operations

Here's a few additional operations for working with slices of bytes:

* `ContainsOnly(b, chars)`
* `TrimToOnly(b, chars)`
* `FromRange(a, z)`
* `Reverse(b)`
* `Indent(b, indent)`

## Deferred Handling

Handling deferred functions is a bit of a pain. This provides a way to handle
certain deferred calls with a little less hassle:

* `Error` (allows capturing both errors that may occur during defer)

## FileSystem Operations

The built-in `io/fs` package is fine for reading files, but is missing write interfaces. This provides them:

* `FileWriter` (analogous to `fs.File`, but for writes)
* `CreateFS`
* `WriteFileFS`

And then provides some additional general interfaces:

* `ReaderFS`
* `ReaderWriterFS`
* `WriterFS`

And provides some convenience functions for working with them:

* `MkdirAll(fsys, name, perm)`
* `WriteFile(fsys, name, data, perm)`

## Generic/Interface Comparison

The built-in comparable generic is pretty weak. Just good enough for map key
types, but not much else. We provide a compare.Able interface which covers all
the types that can be compared using `<`, `>`, `<=`, and `>=` operators, in
addition to `==` and `!=`. Then we provide the following operations on them:

 * `CountDeltas(a, b, delta)`
 * `Max(a, b)`
 * `Min(a, b)`
 * `Less(a, b)`
 * `Zero()`
 * `FirstNonZero(...vals)`
 * `FirstNonNil(...vals)`

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

 * `FromRange(a, z, delta)`
 * `Delete(slice, index)`
 * `DeleteValue(slice, value)`
 * `DeleteAllValues(slice, value)`
 * `Insert(slice, index, value)`
 * `Push(slice, value)`
 * `Pop(slice)`
 * `Shift(slice)`
 * `Unshift(slice, value)`
 * `Concat(slice...)`
 * `Map(slice, mapper)`
 * `MapSlice(slice, mapper)`
 * `MapMap(slice, mapper)`
 * `Reduce(slice, reducer)`
 * `Reductions(slice, reducer)`
 * `ReduceAcc(slice, start, reducer)`
 * `ReductionsAcc(slice, start, reducer)`
 * `Sum(slice)`
 * `Product(slice)`
 * `Grep(slice, predicate)`
 * `GrepIndex(slice, predicate)`
 * `Any(slice, predicate)`
 * `All(slice, predicate)`
 * `None(slice, predicate)`
 * `NotAll(slice, predicate)`
 * `First(slice, predicate)`
 * `FirstIndex(slice, predicate)`
 * `FirstOr(slice, missed, predicate)`
 * `Head(slice, n)`
 * `Tail(slice, n)`
 * `Shuffle(slice)`
 * `Sample(slice, n)`
 * `Uniq(slice)`
 * `UniqInPlace(slice)`

## Generic Map Transformations

Frequently, you need to get just the keys or values off of a slice. This
provides those operations among others for maps:

 * `Copy(dst, src)`
 * `CopyInit(dst, src)`
 * `Keys(map)`
 * `Values(map)`
 * `KVs(map)`

## Generic Map Operations

There are a lot of common map operations that are missing from the built-in
standard library:

 * `Merge(...maps)`
 * `MergeInPlace(base, ...maps)`
 * `Diff(a, b)`

## Generic Sets

There's no built-in set operation. This library provides one. It provides the
following methods on sets:

 * `Contains(val)`
 * `Insert(val)`
 * `Delete(val)`
 * `Len()`
 * `SubsetOf(set)`
 * `Keys()`

And then it also provides these functions to work on sets:

 * `Intersects(a, b)`
 * `Intersection(b, b)`
 * `Union(a, b)`
 * `Difference(a, b)`
 * `Copy(dst, src)`
 * `CopyInit(dst, src)`
 * `Diff(a, b)`

## String Operations

Some additional string operations are provided that are missing from the
built-in `strings` package:

 * `ContainsOnly(s, chars)`
 * `FromRange(a, z)`
 * `Reverse(s)`
 * `Increment(s)`
 * `IncrementWithSets(s, sets)`
 * `Indent(s, indent)`

The latter two operations might warrant their own library, but I put them here for now.
