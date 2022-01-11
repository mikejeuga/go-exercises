package add_test

import (
	"github.com/matryer/is"
	"github.com/mikejeuga/go-exercises/add"
	"testing"
)

func Test(t *testing.T) {
    t.Parallel()
	is := is.New(t)
	for _, tt := range []struct {
		name string
		arg1, arg2 int
		got int
		want int


	}{
		{
			name: "The Add function add 2 values",
			arg1: 0,
			arg2: 0,
			want: 0,
		},

		{
			name: "The Add function add 2 values",
			arg1: 1,
			arg2: 1,
			want: 2,
		},

		{
			name: "The Add function add 2 values",
			arg1: 3,
			arg2: 3,
			want: 6,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			tt.got = add.Add(tt.arg1, tt.arg2)
			is.Equal(tt.got, tt.want)
		})
	}
 }