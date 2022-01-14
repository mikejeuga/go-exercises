package cli_test

import (
	"bytes"
	"github.com/matryer/is"
	"github.com/mikejeuga/go-exercises/src/adapters/cli"
	"testing"
)

func TestAdd(t *testing.T) {
	is := is.New(t)
	t.Run("", func(t *testing.T) {
		command := cli.NewCommand()
		buffer := bytes.Buffer{}
		ints := []byte(`1
2
3
4
6
9`)
		data := bytes.NewReader(ints)


		command.Add(data, &buffer)
		want := `25
`
		is.Equal(buffer.String(), want)

	})

	//TODO: table test the below.

	t.Run("Prints the sum of integers and print thousands with commas", func(t *testing.T) {
		command := cli.NewCommand()
		buffer := bytes.Buffer{}
		ints := []byte(`10000
10000`)

		data := bytes.NewReader(ints)
		command.Add(data, &buffer)
		want := `20,000
`
		is.Equal(buffer.String(), want)

	})

	t.Run("Prints the sum of integers and print millions with commas", func(t *testing.T) {
		command := cli.NewCommand()
		buffer := bytes.Buffer{}
		ints := []byte(`1000000
1000000`)

		data := bytes.NewReader(ints)

		 command.Add(data, &buffer)
		want := `2,000,000
`
		is.Equal(buffer.String(), want)
	})

	t.Run("Prints the sum of the integer arguments taken", func(t *testing.T) {
		command := cli.NewCommand()
		buffer := bytes.Buffer{}
		ints := []byte("1 2 3 4 6 9")


		data := bytes.NewReader(ints)
		command.Add(data, &buffer)

		want := `25
`
		is.Equal(buffer.String(), want)



	})

	t.Run("Prints the sum of the integer arguments taken", func(t *testing.T) {
		command := cli.NewCommand()
		buffer := bytes.Buffer{}
		ints := []byte("1,2,3,4,6,9")


		data := bytes.NewReader(ints)
		command.Add(data, &buffer)

		want := `25
`
		is.Equal(buffer.String(), want)



	})
}
