package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/set"
)

func TestCopy(t *testing.T) {
	s1 := set.New(1, 2, 3)
	s2 := set.New[int]()
	set.Copy(s2, s1)

	assert.Equal(t, s1, s2)
}

func TestCopyInit(t *testing.T) {
	s1 := set.New(1, 2, 3)
	var s2 set.Set[int]
	set.CopyInit(&s2, s1)

	assert.Equal(t, s1, s2)
}

func TestNew(t *testing.T) {
	ss := []set.Set[int]{
		set.New(1, 2, 3),
		set.New(3, 2, 1),
		set.New(1, 3, 2),
		set.New(2, 1, 3),
		set.New(2, 3, 1),
		set.New(3, 1, 2),
	}

	for i, s1 := range ss {
		for j, s2 := range ss {
			assert.Equalf(t, s1, s2, "ss[%d] == ss[%d]", i, j)
		}
	}
}

func TestSet_Contains(t *testing.T) {
	s := set.New(1, 2, 3)

	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))
	assert.False(t, s.Contains(4))
}

func TestSet_Delete(t *testing.T) {
	s := set.New(1, 2, 3)

	s.Delete(0)
	assert.Equal(t, set.New(1, 2, 3), s)

	s.Delete(1)
	assert.Equal(t, set.New(2, 3), s)

	s.Delete(2)
	assert.Equal(t, set.New(3), s)

	s.Delete(3)
	assert.Equal(t, set.New[int](), s)

	s.Delete(4)
	assert.Equal(t, set.New[int](), s)
}

func TestSet_Insert(t *testing.T) {
	s := set.New[int]()

	assert.Equal(t, set.New[int](), s)

	s.Insert(1)
	assert.Equal(t, set.New(1), s)

	s.Insert(2)
	assert.Equal(t, set.New(1, 2), s)

	s.Insert(1)
	assert.Equal(t, set.New(1, 2), s)

	s.Insert(3)
	assert.Equal(t, set.New(3, 2, 1), s)
}

func TestSet_Len(t *testing.T) {
	s := set.New(1, 2, 3)
	assert.Equal(t, len(s), s.Len())
}

func TestSet_SubsetOf(t *testing.T) {
	s1 := set.New(1, 2, 3)
	s2 := set.New(2, 3)
	s3 := set.New(3, 4)

	assert.True(t, s1.SubsetOf(s1))
	assert.False(t, s1.SubsetOf(s2))
	assert.False(t, s1.SubsetOf(s3))
	assert.True(t, s2.SubsetOf(s1))
	assert.True(t, s2.SubsetOf(s2))
	assert.False(t, s2.SubsetOf(s3))
	assert.False(t, s3.SubsetOf(s1))
	assert.False(t, s3.SubsetOf(s2))
	assert.True(t, s3.SubsetOf(s3))
}

func TestDifference(t *testing.T) {
	s1 := set.New(1, 2, 3)
	s2 := set.New(3, 4, 5)

	assert.Equal(t, set.New[int](), set.Difference(s1, s1))
	assert.Equal(t, set.New(1, 2), set.Difference(s1, s2))
	assert.Equal(t, set.New(4, 5), set.Difference(s2, s1))
}

func TestIntersection(t *testing.T) {
	s1 := set.New(1, 2, 3)
	s2 := set.New(3, 4, 5, 6)

	assert.Equal(t, set.New(1, 2, 3), set.Intersection(s1, s1))
	assert.Equal(t, set.New(3), set.Intersection(s1, s2))
	assert.Equal(t, set.New(3), set.Intersection(s2, s1))
}

func TestIntersects(t *testing.T) {
	s1 := set.New(1, 2, 3)
	s2 := set.New(3, 4, 5)
	s3 := set.New(5, 6, 7)

	assert.True(t, set.Intersects(s1, s1))
	assert.True(t, set.Intersects(s1, s2))
	assert.False(t, set.Intersects(s1, s3))
	assert.True(t, set.Intersects(s2, s1))
	assert.True(t, set.Intersects(s2, s2))
	assert.True(t, set.Intersects(s2, s3))
	assert.False(t, set.Intersects(s3, s1))
	assert.True(t, set.Intersects(s3, s2))
	assert.True(t, set.Intersects(s3, s3))
}

func TestUnion(t *testing.T) {
	s1 := set.New(1, 2, 3)
	s2 := set.New(3, 4, 5, 6)

	assert.Equal(t, set.New(1, 2, 3, 4, 5, 6), set.Union(s1, s2))
	assert.Equal(t, set.New(1, 2, 3, 4, 5, 6), set.Union(s2, s1))
}
