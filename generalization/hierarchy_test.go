package generalization

import (
	"k-anon/testutil"
	"strings"
	"testing"
)

func Test_Levels(t *testing.T) {
	h := GetGradeHierarchy1()
	expected := 3
	actual := h.GetLevelCount()
	testutil.AssertEquals(expected, actual, t)
}

func Test_GetLevel(t *testing.T) {
	h := GetGradeHierarchy1()
	expected := []*Partition{NewPartition("A+", "A", "A-", "B+", "B", "B-", "C+", "C", "C-")}
	actual := h.GetLevel(2)
	testutil.AssertEquals(len(expected), len(actual), t)
	for i := range expected {
		if !expected[i].Equals(actual[i]) {
			t.Errorf("element mismatch at index %d", i)
		}
	}
}

func Test_Valid(t *testing.T) {
	h := GetGradeHierarchy1()
	assertValid(h, t)
}

func Test_InvalidMultipleValuesOnLevel(t *testing.T) {
	h := &Hierarchy{Partitions: [][]*Partition{
		{
			NewPartition(1, 2, 3),
			NewPartition(5, 6, 3), // <= error: 3 is present in both partitions in the same level
		},
	}}
	assertInvalid(h, t)
}

func Test_InvalidItemsDoNotAddUp(t *testing.T) {
	h := &Hierarchy{Partitions: [][]*Partition{
		{
			NewPartition(1),
			NewPartition(2),
			NewPartition(3),
			NewPartition(4),
		},
		{
			NewPartition(1, 2),
			NewPartition(3, 5), // <= error: 5 is not part of the hierarchy
		},
		{
			NewPartition(1, 2, 3, 4),
		},
	}}
	assertInvalid(h, t)
}

func Test_Find(t *testing.T) {
	tests := []struct {
		name     string
		item     interface{}
		level    int
		expected *Partition
	}{
		{"Exists GetLevel 0", "C", 0, NewPartition("C")},
		{"Exists GetLevel 1", "C", 1, NewPartition("C+", "C", "C-")},
		{"Exists GetLevel 2", "C", 2, NewPartition("A+", "A", "A-", "B+", "B", "B-", "C+", "C", "C-")},
		{"Missing GetLevel 0", "X", 0, nil},
		{"Missing GetLevel 1", "X", 1, nil},
		{"Missing GetLevel 2", "X", 2, nil},
	}
	h := GetGradeHierarchy1()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := h.Find(test.item, test.level)
			if test.expected == nil {
				testutil.AssertNil(actual, t)
			} else if !test.expected.Equals(actual) {
				t.Errorf("item was not located in the correct partition: %v", actual)
			}
		})
	}
}

func Test_GetLevelUnderIndex(t *testing.T) {
	h := GetGradeHierarchy1()
	idx := -1
	actual := h.GetLevel(idx)
	testutil.AssertNil(actual, t)
}

func Test_GetLevelOverIndex(t *testing.T) {
	h := GetGradeHierarchy1()
	idx := h.GetLevelCount() // max index + 1
	actual := h.GetLevel(idx)
	testutil.AssertNil(actual, t)
}

func Test_FindUnderIndex(t *testing.T) {
	h := GetGradeHierarchy1()
	idx := -1
	actual := h.Find("C", idx)
	testutil.AssertNil(actual, t)
}

func Test_FindOverIndex(t *testing.T) {
	h := GetGradeHierarchy1()
	idx := h.GetLevelCount() // max index + 1
	actual := h.Find("C", idx)
	testutil.AssertNil(actual, t)
}

func Test_StringEmpty(t *testing.T) {
	h := &Hierarchy{}
	expected := ""
	actual := h.String()
	testutil.AssertEquals(expected, actual, t)
}

func Test_StringSinglePartition(t *testing.T) {
	p := NewPartition(1, 2)
	h := &Hierarchy{Partitions: [][]*Partition{{p}}}
	actual := h.String()
	expected1 := "[1, 2]"
	expected2 := "[2, 1]"
	if !strings.Contains(actual, expected1) && !strings.Contains(actual, expected2) {
		t.Errorf("expected %s to contain partition %s", actual, p.String())
	}
}

func assertValid(h *Hierarchy, t *testing.T) {
	if !h.IsValid() {
		t.Errorf("hierarchy should be valid")
	}
}

func assertInvalid(h *Hierarchy, t *testing.T) {
	if h.IsValid() {
		t.Errorf("hierarchy should be invalid")
	}
}

func assertHierarchyEquals(h1 *Hierarchy, h2 *Hierarchy, t *testing.T) {
	if !Equals(h1, h2) {
		t.Errorf("expected:\n%s\nactual:\n%s\n", h1, h2)
	}
}

func buildIntHierarchy(items ...int) *Hierarchy {
	builder := &IntegerHierarchyBuilder{
		Items: items,
	}
	return builder.NewIntegerHierarchy()
}
