package main

import (
	"reflect"
	"testing"
)

func TestReadBinaryWatch(t *testing.T) {

	checkResult := func(t *testing.T, got, want [][]int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want %#v", got, want)
		}
	}

	t.Run("test1", func(t *testing.T) {

		candidates := []int{2, 3, 6, 7}
		target := 7

		want := [][]int{
			{2, 2, 3},
			{7},
		}

		got := combinationSum(candidates, target)

		checkResult(t, got, want)
	})

	t.Run("test2", func(t *testing.T) {

		candidates := []int{2, 3, 5}
		target := 8

		want := [][]int{
			{2, 2, 2, 2},
			{2, 3, 3},
			{3, 5},
		}

		got := combinationSum(candidates, target)

		checkResult(t, got, want)
	})

	t.Run("test3", func(t *testing.T) {

		candidates := []int{2}
		target := 1

		var want [][]int

		got := combinationSum(candidates, target)

		checkResult(t, got, want)
	})

	t.Run("test4", func(t *testing.T) {

		candidates := []int{8, 2, 6, 3}
		target := 13

		want := [][]int{
			{8, 2, 3},
			{2, 2, 2, 2, 2, 3},
			{2, 2, 6, 3},
			{2, 2, 3, 3, 3},
		}

		got := combinationSum(candidates, target)

		checkResult(t, got, want)
	})
}
