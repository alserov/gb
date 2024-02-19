package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestMergeSlices(t *testing.T) {
	tests := []struct {
		s1     []int
		s2     []int
		expect []int
		name   string
	}{
		{
			s1:     []int{1, 3, 6},
			s2:     []int{2, 4, 5},
			expect: []int{1, 2, 3, 4, 5, 6},
			name:   "default",
		},
		{
			s1:     []int{2, 2, 2},
			s2:     []int{1, 2, 3, 4},
			expect: []int{1, 2, 2, 2, 2, 3, 4},
			name:   "with equal els 1",
		},
		{
			s1:     []int{1, 4, 6},
			s2:     []int{2, 4, 5},
			expect: []int{1, 2, 4, 4, 5, 6},
			name:   "with equal els 2",
		},
		{
			s1:     []int{},
			s2:     []int{2, 4, 5},
			expect: []int{2, 4, 5},
			name:   "s1 empty",
		},
		{
			s1:     []int{1, 4, 6},
			s2:     []int{},
			expect: []int{1, 4, 6},
			name:   "s2 empty",
		},
		{
			s1:     []int{},
			s2:     []int{},
			expect: []int{},
			name:   "empty",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := mergeSortedSlices(tc.s1, tc.s2)
			eq := reflect.DeepEqual(tc.expect, res)
			require.True(t, eq)
		})
	}
}

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		in   []int
		out  []int
		name string
	}{
		{
			in:   []int{4, 5, 1, 2, 3},
			out:  []int{1, 2, 3, 4, 5},
			name: "default",
		},
		{
			in:   []int{},
			out:  []int{},
			name: "empty",
		},
		{
			in:   []int{1},
			out:  []int{1},
			name: "1 el",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			bubbleSort(tc.in)
			eq := reflect.DeepEqual(tc.out, tc.in)
			require.True(t, eq)
		})
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		in   []int
		out  []int
		name string
	}{
		{
			in:   []int{4, 5, 1, 2, 3},
			out:  []int{1, 2, 3, 4, 5},
			name: "default",
		},
		{
			in:   []int{},
			out:  []int{},
			name: "empty",
		},
		{
			in:   []int{1},
			out:  []int{1},
			name: "1 el",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			quickSort(tc.in)
			eq := reflect.DeepEqual(tc.out, tc.in)
			require.True(t, eq)
		})
	}
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		in   []int
		out  []int
		name string
	}{
		{
			in:   []int{4, 5, 1, 2, 3},
			out:  []int{1, 2, 3, 4, 5},
			name: "default",
		},
		{
			in:   []int{},
			out:  []int{},
			name: "empty",
		},
		{
			in:   []int{1},
			out:  []int{1},
			name: "1 el",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.in = mergeSort(tc.in)
			fmt.Println(tc.in)
			eq := reflect.DeepEqual(tc.out, tc.in)
			require.True(t, eq)
		})
	}
}
