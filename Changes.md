## WIP  TBD

 * :boom: Breaking Change :boom:: Now requires Go 1.20.
 * Adding deferred.Error for helping with deferred error handling.

## v0.6.0  2023-08-12

 * Adding slices.Insert

## v0.5.0  2023-07-25

 * Adding fs.CreateFS, fs.WriteFileFS, fs.ReaderFS, fs.ReaderWriterFS, and fs.WriterFS
 * Adding fs.MkdirAll and fs.WriteFile

## v0.4.0  2023-07-14

 * Added maps.MergeInPlace.
 * Added slices.MapSlice and slices.MapMap.
 * maps.Merge now allocates room equal to the size of all input maps in the output map.

## v0.3.0  2023-07-13

 * Added slices.FirstIndex and slices.GrepIndex.
 * Added set.NewSized and set.Diff.
 * Added maps.Diff.
 * Fixed missing test coverage on generic.FirstNonNil and generic.FirstNonZero.

## v0.2.0  2023-07-08

 * Added generic.FirstNonZero and generic.FirstNonNil.
 * Added maps.Merge.

## v0.1.2  2023-07-07

 * Fix a silly release typo.

## v0.1.1  2023-07-07

 * Added set.Set.Keys method.

## v0.1.0  2023-05-27

 * Added generic.CountDeltas function.
 * Added slices.FromRange function.
 * Added strings.Increment, strings.IncrementWithSets, strings.IncrementSet, strings.NumericSet, strings.NumericSetRange, strings.LetterSet, strings.LetterSetRange, and strings.SeqSet.

## v0.0.4  2023-05-27

 * Added a LICENSE file.

## v0.0.3  2023-05-09

 * Added structs.ApplyDefaults function.

## v0.0.2  2023-04-21

 * Small improvement to the behavior/performance of slices.Delete
 * Added slices.Map, slices.Reduce, slices.ReduceAcc, slices.Reductions,
   slices.ReductionAcc, slices.Sum, slices.Product, slices.Grep, slices.Any,
   slices.All, slices.None, slices.NotAll, slices.First, slices.FirstOr
   functions.
 * Added slices.Head and slices.Tail functions.
 * Added slices.Shuffle, slices.Sample, slices.Uniq, and slices.UniqInPlace
   functions.

## v0.0.1  2023-04-17

 * Added slices.Concat

## v0.0.0  2023-04-08

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
