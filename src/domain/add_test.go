package domain_test

import (
	"github.com/matryer/is"
	add "github.com/mikejeuga/go-exercises/src/domain"
	"testing"
)

func TestSum(t *testing.T) {
    t.Parallel()
	is := is.New(t)
	for _, tt := range []struct {
		name string
		values []int
		got int
		want int


	}{
		{
			name: "The Add function adds 2 parameters and return the correct value",
			values: []int{},
			want: 0,
		},

		{
			name: "The Add function adds 1 and 1 and returns 2",
			values: []int{1, 1},
			want: 2,
		},

		{
			name: "The Add function adds 3 and 3 and returns 6",
			values: []int{3, 3},
			want: 6,
		},
		{
			name: "The Add function adds more than 3 parameters and return the correct value",
			values: []int{3, 3,3},
			want: 9,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			service := add.Service{}
			tt.got = service.Add(tt.values...)
			is.Equal(tt.got, tt.want)
		})
	}
 }

