package add_test

import (
	"bytes"
	"github.com/matryer/is"
	"github.com/mikejeuga/go-exercises/add"
	"testing"
)

func TestAdd(t *testing.T) {
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
			tt.got = add.Add(tt.values...)
			is.Equal(tt.got, tt.want)
		})
	}
 }

func TestPrintAdd(t *testing.T) {
	t.Run("Prints the sum of the integer arguments taken", func(t *testing.T) {
		is := is.New(t)
		buffer := bytes.Buffer{}
		ints := []byte("1 2 3 4 6 9")
		data := bytes.NewReader(ints)

		add.PrintAdd(data, &buffer)

		want := "25"

		is.Equal(buffer.String(), want)



	})

	//TODO: table test the below.

	t.Run("Prints the sum of the integer arguments taken", func(t *testing.T) {
		is := is.New(t)
		buffer := bytes.Buffer{}
		ints := []byte("10000 10000")
		data := bytes.NewReader(ints)

		add.PrintAdd(data, &buffer)

		want := "20,000"

		is.Equal(buffer.String(), want)

	})

	t.Run("Prints the sum of the integer arguments taken", func(t *testing.T) {
		is := is.New(t)
		buffer := bytes.Buffer{}
		ints := []byte("1000000 1000000")
		data := bytes.NewReader(ints)

		add.PrintAdd(data, &buffer)

		want := "2,000,000"

		is.Equal(buffer.String(), want)
	})
}
