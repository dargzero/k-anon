package hierarchy

import (
	"bitbucket.org/dargzero/k-anon/partition"
)

func GetGradeHierarchy() Hierarchy {
	h, err := Build(partition.NewSet("A+", "A", "A-", "B+", "B", "B-", "C+", "C", "C-"),
		N(partition.NewSet("A+", "A", "A-"),
			N(partition.NewSet("A+")),
			N(partition.NewSet("A")),
			N(partition.NewSet("A-"))),
		N(partition.NewSet("B+", "B", "B-"),
			N(partition.NewSet("B+")),
			N(partition.NewSet("B")),
			N(partition.NewSet("B-"))),
		N(partition.NewSet("C+", "C", "C-"),
			N(partition.NewSet("C+")),
			N(partition.NewSet("C")),
			N(partition.NewSet("C-"))))
	if err != nil {
		panic(err)
	}
	return h
}