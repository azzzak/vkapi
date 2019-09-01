package vkapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_structToMap(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
		want map[string]string
	}{
		{
			name: "string",
			arg: struct {
				Field string
			}{
				Field: "string",
			},
			want: map[string]string{
				"field": "string",
			},
		},
		{
			name: "int",
			arg: struct {
				Field int
			}{
				Field: 5,
			},
			want: map[string]string{
				"field": "5",
			},
		}, {
			name: "int zero",
			arg: struct {
				Field int
			}{
				Field: 0,
			},
			want: map[string]string{},
		}, {
			name: "int64",
			arg: struct {
				Field int64
			}{
				Field: 512312414235,
			},
			want: map[string]string{
				"field": "512312414235",
			},
		}, {
			name: "uint",
			arg: struct {
				Field uint
			}{
				Field: 7,
			},
			want: map[string]string{
				"field": "7",
			},
		}, {
			name: "float32",
			arg: struct {
				Field float32
			}{
				Field: 3.1415926535,
			},
			want: map[string]string{
				"field": "3.1415927",
			},
		}, {
			name: "bool true",
			arg: struct {
				Field bool
			}{
				Field: true,
			},
			want: map[string]string{
				"field": "1",
			},
		}, {
			name: "bool false",
			arg: struct {
				Field bool
			}{
				Field: false,
			},
			want: map[string]string{},
		}, {
			name: "[]int",
			arg: struct {
				Field []int
			}{
				Field: []int{1, 2, 3, 10},
			},
			want: map[string]string{
				"field": "1,2,3,10",
			},
		},
		{
			name: "[]int empty",
			arg: struct {
				Field []int
			}{
				Field: []int{},
			},
			want: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := structToMap(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_toBool(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want bool
	}{
		{
			name: "zero",
			arg:  0,
			want: false,
		}, {
			name: "negative",
			arg:  -1,
			want: false,
		},
		{
			name: "one",
			arg:  1,
			want: true,
		}, {
			name: "big positive",
			arg:  10,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toBool(tt.arg); got != tt.want {
				t.Errorf("toBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
