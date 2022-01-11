package add_test

import (
	"bytes"
	"github.com/matryer/is"
	"github.com/mikejeuga/go-exercises/add"
	"testing"
)

func Test(t *testing.T) {
    t.Parallel()
	is := is.New(t)
	for _, tt := range []struct {
		name string
		values []int
		got int
		want int


	}{
		{
			name: "The Sum function adds 2 parameters and return the correct value",
			values: []int{},
			want: 0,
		},

		{
			name: "The Sum function adds 1 and 1 and returns 2",
			values: []int{1, 1},
			want: 2,
		},

		{
			name: "The Sum function adds 3 and 3 and returns 6",
			values: []int{3, 3},
			want: 6,
		},
		{
			name: "The Sum function adds more than 3 parameters and return the correct value",
			values: []int{3, 3,3},
			want: 9,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			tt.got = add.Sum(tt.values...)
			is.Equal(tt.got, tt.want)
		})
	}
 }

func TestAdd(t *testing.T) {
	t.Run("Prints the sum of the integer arguments taken", func(t *testing.T) {
		is := is.New(t)
		buffer := bytes.Buffer{}
		ints := []byte("1 2 3 4 6 9")
		data := bytes.NewReader(ints)

		add.Add(data, &buffer)

		want := "25"

		is.Equal(buffer.String(), want)



	})
}
