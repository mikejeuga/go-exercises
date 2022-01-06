# Track 1

Each part builds on top of the previous one. After completing each part, tag the commit as `track-1-part-1` etc.

If you make any assumptions about the feature (e.g. negative numbers are not supported), mention it in the readme.


## Part 1

Write a function to add two integers. 

e.g. `add(5, 4)` should return `9`

## Part 2

Modify the function to add a variable number of integers.

e.g. `add(5, 4, 3, 2, -10)` should return `4`

## Part 3

Create a program `add` that will take any number of integer arguments and print out the sum. Ignore any non-integers in the list.

Example usage of the program: `add 4 5 32 100 867543` should output `867684`

## Part 4

Modify the output of the program so numbers larger than 9999 are formatted in groups of thousands. e.g. 9999 will be shown as `9999`, 10000 will be shown as `10,000`, and 1234567890 will be shown as `1,234,567,890`

## Part 5

Change the program so that when it is called with no arguments (`add`), it read the numbers from a file called `input.txt` in the same directory. The file will have one number per line. Ignore any non-integers in the file.

Note that the old behaviour of the program (taking a list of integers as arguments) still works.

Example of `input.txt`:

```
4
5
32
100
867543
```

## Part 6

Change the program to specify the input file name as an argument. e.g. `add --input-file data/input.txt`. The file can be anywhere on the file system.

## Part 7

Change the program so it can also read the numbers from a CSV file. e.g. `add --input-file data/input2.csv` where `input2.csv` looks like this:

```
4,5,32,100,867543
```

You can assume the file is a CSV if the extension is .csv

## Part 8

Change the program so it does not rely on the .csv extension. It should auto-detect if the file is a CSV or newline-separated file. Therefore `add --input-file input.txt` should work regardless of whether the numbers in `input.txt` are separated by commas or newlines.

## Part 9

Change the program so it accepts multiple files as input, each of which can be any supported format.

e.g. `add --input-file one.txt --input-file data/two.txt`

The output is the sum of all the numbers from all the specified files.

## Part 10

Write a CI pipeline in Github Actions to run all your tests when you push code to the trunk.

## Part 11

Change the program so it discards any duplicate numbers. e.g. `add 2 2 2 3 4 4 4` will output `9` (2+3+4). 

## Part 12

Create a new program `math` so when it is invoked as `math --web-server`, it starts a web server with an endpoint `POST /add?num=4&num=5&num=32` that returns the response `41` as text. 

(Ignore why it is a POST method, for now)

## Part 13

Change the math web server so it can also read a form-urlencoded body from the request. 

Example request:

```
POST /add

num=4&num=5&num=32
```

will still return `41` as text.

## Part 14

Change the math web server so it can also accept numbers as a JSON array.

Example request:

```
POST /add

{
    "nums": [4, 5, 32]
}
```

Research the request headers that a client will send that will allow different forms of the request body to be accepted by the server.

