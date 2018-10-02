package main

import (
	"testing"

	"../lnklist"
	"github.com/stretchr/testify/require"
)

func Test_Partition_Empty(t *testing.T) {
	result := partition(nil, 5)
	require.Nil(t, result)
}

func Test_Partition_Single(t *testing.T) {
	node := lnklist.New(1, nil)

	result := partition(node, 5)

	require.Equal(t, node, result)
}

func Test_Partition_Double_NoItemsMatchCriteria(t *testing.T) {
	node := lnklist.New(1, lnklist.New(2, nil))
	expectedResult := lnklist.New(2, lnklist.New(1, nil))

	result := partition(node, 5)

	require.Equal(t, expectedResult, result)
}

func Test_Partition_Double_MatchLastItem(t *testing.T) {
	node := lnklist.New(1, lnklist.New(2, nil))
	expectedResult := lnklist.New(1, lnklist.New(2, nil))

	result := partition(node, 2)

	require.Equal(t, expectedResult, result)
}

func Test_Partition_Double_MatchFirstItem(t *testing.T) {
	node := lnklist.New(1, lnklist.New(2, nil))
	expectedResult := lnklist.New(1, lnklist.New(2, nil))

	result := partition(node, 1)

	require.Equal(t, expectedResult, result)
}

func Test_Partition_Multiple_NoItemsMatchCriteria_1(t *testing.T) {
	node := lnklist.New(3, lnklist.New(5, lnklist.New(8, lnklist.New(5, lnklist.New(10, lnklist.New(2, lnklist.New(1, nil)))))))
	expectedResult := lnklist.New(3, lnklist.New(5, lnklist.New(8, lnklist.New(5, lnklist.New(10, lnklist.New(2, lnklist.New(1, nil)))))))

	result := partition(node, 0)

	require.Equal(t, expectedResult, result)
}

func Test_Partition_Multiple_NoItemsMatchCriteria_2(t *testing.T) {
	node := lnklist.New(3, lnklist.New(5, lnklist.New(8, lnklist.New(5, lnklist.New(10, lnklist.New(2, lnklist.New(1, nil)))))))
	expectedResult := lnklist.New(1, lnklist.New(2, lnklist.New(10, lnklist.New(5, lnklist.New(8, lnklist.New(5, lnklist.New(3, nil)))))))
	result := partition(node, 30)

	require.Equal(t, expectedResult, result)
}

func Test_Partition_Multiple_MatchLastItem(t *testing.T) {
	node := lnklist.New(3, lnklist.New(5, lnklist.New(8, lnklist.New(5, lnklist.New(10, lnklist.New(2, lnklist.New(1, nil)))))))
	expectedResult := lnklist.New(3, lnklist.New(5, lnklist.New(8, lnklist.New(5, lnklist.New(10, lnklist.New(2, lnklist.New(1, nil)))))))

	result := partition(node, 1)

	require.Equal(t, expectedResult, result)
}

func Test_Partition_Multiple_MatchFirstItem(t *testing.T) {
	node := lnklist.New(3, lnklist.New(5, lnklist.New(8, lnklist.New(5, lnklist.New(10, lnklist.New(2, lnklist.New(1, nil)))))))
	expectedResult := lnklist.New(1, lnklist.New(2, lnklist.New(3, lnklist.New(5, lnklist.New(8, lnklist.New(5, lnklist.New(10, nil)))))))

	result := partition(node, 3)

	require.Equal(t, expectedResult, result)
}

func Test_Partition_Multiple_Normal(t *testing.T) {
	node := lnklist.New(3, lnklist.New(5, lnklist.New(8, lnklist.New(5, lnklist.New(10, lnklist.New(2, lnklist.New(1, nil)))))))
	expectedResult := lnklist.New(1, lnklist.New(2, lnklist.New(3, lnklist.New(5, lnklist.New(8, lnklist.New(5, lnklist.New(10, nil)))))))

	result := partition(node, 5)

	require.Equal(t, expectedResult, result)
}
