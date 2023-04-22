v0.0.2  2023-04-21

* Small improvement to the behavior/performance of slices.Delete
* Added slices.Map, slices.Reduce, slices.ReduceAcc, slices.Reductions,
  slices.ReductionAcc, slices.Sum, slices.Product, slices.Grep, slices.Any,
  slices.All, slices.None, slices.NotAll, slices.First, slices.FirstOr
* functions.
* Added slices.Head and slices.Tail functions.
* Added slices.Shuffle, slices.Sample, slices.Uniq, and slices.UniqInPlace
  functions.

v0.0.1  2023-04-17

 * Added slices.Concat

v0.0.0  2023-04-08

 * Initial release.
 * Tools to makeup for limits in the std library.
 * Includes generic package features: Comparison, LessThan, Equal, GreaterThan, 
   ComparableInterface, Comparable, Compare, Max, MaxI, Min, MinI,
   Less, LessI, and Zero
 * Includes iterable package features: Interface
 * Includes maps package features: Copy, CopyInit, Clear, Iterator, 
   Iterator.Len, Iterator.Next, Iterator.Index, Iterator.ID, Iterator.Val, Keys,
   Values, and KVs
 * Includes set package features: Set, New, Set.Contains, Set.Insert, 
   Set.Delete, Set.Len, Set.SubsetOf, Set.Intersects, Set.Intersection,
   Set.Union, Set.Difference, Copy, CopyInit
 * Includes slices package features: Iterator, NewIterator, Iterator.Len,
   Iterator.Next, Iterator.Index, Iterator.ID, Iterator.Val, Delete, 
   DeleteValue, DeleteAllValues, Push, Pop, Shift, Unshift
 * Includes strings package features: ContainsOnly, FromRange, and Reverse
