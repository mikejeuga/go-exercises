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
			name: "The Add function add 2 value",
			got: add.Add(),
			want: 0,
		},

	} {
		t.Run(tt.name, func(t *testing.T) {
			is.Equal(tt.got, tt.want)
		})
	}
 }