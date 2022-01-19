package cli_test

import (
	"bytes"
	"context"
	"github.com/matryer/is"
	"github.com/mikejeuga/go-exercises/src/adapters/cli"
	"testing"
)

func TestAdd(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
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
		command.Add(ctx, &buffer, data)

		want := `1
2
3
4
6
9
Total: 25
`
		is.Equal(buffer.String(), want)

	})

	//TODO: table test the below.

	t.Run("Prints the sum of integers and print thousands with commas", func(t *testing.T) {
		command := cli.NewCommand()
		buffer := bytes.Buffer{}
		ints := []byte(`10000
20000`)

		data := bytes.NewReader(ints)
		command.Add(ctx, &buffer, data)
		want := `10000
20000
Total: 30,000
`
		is.Equal(buffer.String(), want)

	})

	t.Run("Prints the sum of integers and print millions with commas", func(t *testing.T) {
		command := cli.NewCommand()
		buffer := bytes.Buffer{}
		ints := []byte(`2000000
1000000`)

		data := bytes.NewReader(ints)

		 command.Add(ctx, &buffer, data)
		want := `2000000
1000000
Total: 3,000,000
`
		is.Equal(buffer.String(), want)
	})

	t.Run("Prints the sum of the integer arguments taken", func(t *testing.T) {
		command := cli.NewCommand()
		buffer := bytes.Buffer{}
		ints := []byte("1 2 3 4 6 9 9")


		data := bytes.NewReader(ints)
		command.Add(ctx, &buffer, data)

		want := `1
2
3
4
6
9
9
Total: 25
`
		is.Equal(buffer.String(), want)



	})

	t.Run("Prints the sum of the integer arguments taken", func(t *testing.T) {
		command := cli.NewCommand()
		buffer := bytes.Buffer{}
		ints := []byte("1,2,3,4,6,9")


		data := bytes.NewReader(ints)
		command.Add(ctx, &buffer, data)

		want := `1
2
3
4
6
9
Total: 25
`
		is.Equal(buffer.String(), want)
	})
}

