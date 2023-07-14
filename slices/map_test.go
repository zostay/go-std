package slices_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/slices"
)

func toString(a int) string {
	return strconv.Itoa(a)
}

func subtract(a, b int) int {
	return a - b
}

func join(a string, b int) string {
	if a == "" {
		return toString(b)
	}
	return fmt.Sprintf("%s-%d", a, b)
}

func isOdd(a int) bool {
	return a%2 == 1
}

func double(a int) int {
	return a * 2
}

func oddly(a int) int {
	return a*2 + 1
}

func TestMap(t *testing.T) {
	in := []int{1, 2, 3}
	ss := slices.Map(in, toString)
	assert.Equal(t, []string{"1", "2", "3"}, ss)
	ss = slices.Map([]int{}, toString)
	assert.Equal(t, []string{}, ss)
}

func TestMapSlice(t *testing.T) {
	halveTheOdds := func(i int) []float64 {
		if i%2 == 1 {
			return []float64{float64(i) / 2.0}
		} else {
			return []float64{}
		}
	}

	in := []int{1, 2, 3}
	ss := slices.MapSlice(in, halveTheOdds)
	assert.Equal(t, []float64{0.5, 1.5}, ss)
	ss = slices.MapSlice([]int{}, halveTheOdds)
	assert.Equal(t, []float64{}, ss)
}

func TestMapMap(t *testing.T) {
	stringToInt := func(in int) map[string]int {
		return map[string]int{
			strconv.Itoa(in): in,
		}
	}
	in := []int{1, 2, 3}
	ss := slices.MapMap(in, stringToInt)
	assert.Equal(t, map[string]int{"1": 1, "2": 2, "3": 3}, ss)
	ss = slices.MapMap([]int{}, stringToInt)
	assert.Equal(t, map[string]int{}, ss)
}

func TestReduce(t *testing.T) {
	in := []int{1, 2, 3}
	s := slices.Reduce(in, subtract)
	assert.Equal(t, -6, s)
	s = slices.Reduce([]int{}, subtract)
	assert.Equal(t, 0, s)

	str := slices.Reduce(in, join)
	assert.Equal(t, "1-2-3", str)
	str = slices.Reduce([]int{}, join)
	assert.Equal(t, "", str)
}

func TestReduceAcc(t *testing.T) {
	in := []int{1, 2, 3}
	s := slices.ReduceAcc(in, 7, subtract)
	assert.Equal(t, 1, s)
	s = slices.ReduceAcc([]int{}, 7, subtract)
	assert.Equal(t, 7, s)

	str := slices.ReduceAcc(in, "steve", join)
	assert.Equal(t, "steve-1-2-3", str)
	str = slices.ReduceAcc([]int{}, "steve", join)
	assert.Equal(t, "steve", str)
}

func TestReductions(t *testing.T) {
	in := []int{1, 2, 3}
	s := slices.Reductions(in, subtract)
	assert.Equal(t, []int{-1, -3, -6}, s)
	s = slices.Reductions([]int{}, subtract)
	assert.Equal(t, []int{}, s)

	str := slices.Reductions(in, join)
	assert.Equal(t, []string{"1", "1-2", "1-2-3"}, str)
	str = slices.Reductions([]int{}, join)
	assert.Equal(t, []string{}, str)
}

func TestReductionsAcc(t *testing.T) {
	in := []int{1, 2, 3}
	s := slices.ReductionsAcc(in, 7, subtract)
	assert.Equal(t, []int{6, 4, 1}, s)
	s = slices.ReductionsAcc([]int{}, 7, subtract)
	assert.Equal(t, []int{}, s)

	str := slices.ReductionsAcc(in, "steve", join)
	assert.Equal(t, []string{"steve-1", "steve-1-2", "steve-1-2-3"}, str)
	str = slices.ReductionsAcc([]int{}, "steve", join)
	assert.Equal(t, []string{}, str)
}

func TestSum(t *testing.T) {
	in := []int{1, 2, 3}
	s := slices.Sum(in)
	assert.Equal(t, 6, s)
	s = slices.Sum([]int{})
	assert.Equal(t, 0, s)
}

func TestProduct(t *testing.T) {
	in := []int{3, 4, 5}
	p := slices.Product(in)
	assert.Equal(t, 60, p)
	p = slices.Product([]int{})
	assert.Equal(t, 1, p)
}

func TestGrep(t *testing.T) {
	in := []int{1, 2, 3}
	assert.Equal(t, []int{1, 3}, slices.Grep(in, isOdd))
	assert.Equal(t, []int{}, slices.Grep([]int{}, isOdd))
}

func TestGrepIndex(t *testing.T) {
	in := []int{1, 2, 3}
	assert.Equal(t, []int{0, 2}, slices.GrepIndex(in, isOdd))
	assert.Equal(t, []int{}, slices.GrepIndex([]int{}, isOdd))
}

func TestAny(t *testing.T) {
	in := []int{1, 2, 3}
	assert.Equal(t, true, slices.Any(in, isOdd))
	assert.Equal(t, false, slices.Any(slices.Map(in, double), isOdd))
	assert.Equal(t, false, slices.Any([]int{}, isOdd))
}

func TestAll(t *testing.T) {
	in := []int{1, 2, 3}
	assert.Equal(t, false, slices.All(in, isOdd))
	assert.Equal(t, true, slices.All(slices.Map(in, oddly), isOdd))
	assert.Equal(t, true, slices.All([]int{}, isOdd))
}

func TestNone(t *testing.T) {
	in := []int{1, 2, 3}
	assert.Equal(t, false, slices.None(in, isOdd))
	assert.Equal(t, true, slices.None(slices.Map(in, double), isOdd))
	assert.Equal(t, true, slices.None([]int{}, isOdd))
}

func TestNotAll(t *testing.T) {
	in := []int{1, 2, 3}
	assert.Equal(t, true, slices.NotAll(in, isOdd))
	assert.Equal(t, false, slices.NotAll(slices.Map(in, oddly), isOdd))
	assert.Equal(t, false, slices.NotAll([]int{}, isOdd))
}

func TestFirst(t *testing.T) {
	in := []int{1, 2, 3}
	v, found := slices.First(in, isOdd)
	assert.Equal(t, 1, v)
	assert.True(t, found)
	v, found = slices.First(slices.Map(in, double), isOdd)
	assert.Equal(t, 0, v)
	assert.False(t, found)
	v, found = slices.First([]int{}, isOdd)
	assert.Equal(t, 0, v)
	assert.False(t, found)
}

func TestFirstIndex(t *testing.T) {
	in := []int{1, 2, 3}
	v := slices.FirstIndex(in, isOdd)
	assert.Equal(t, 0, v)
	v = slices.FirstIndex(slices.Map(in, double), isOdd)
	assert.Equal(t, -1, v)
	v = slices.FirstIndex([]int{}, isOdd)
	assert.Equal(t, -1, v)
}

func TestFirstOr(t *testing.T) {
	in := []int{1, 2, 3}
	v := slices.FirstOr(in, -1, isOdd)
	assert.Equal(t, 1, v)
	v = slices.FirstOr(slices.Map(in, double), -1, isOdd)
	assert.Equal(t, -1, v)
	v = slices.FirstOr([]int{}, -1, isOdd)
	assert.Equal(t, -1, v)
}
