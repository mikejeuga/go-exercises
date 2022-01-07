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

## Part 15

Add a naive authentication to the endpoint. Only requests containing a header `Authorization: Bearer SUPER_SECRET_API_KEY` (that exact API key) should be allowed in. 

Research and implement the behaviour for when:

1. The header is not present
2. The API key is incorrect

## Part 16

Instead of a hard-coded API key (should have never committed the API key to source control anyway), read in a list of API keys from an environment variable. Any request with one of those API keys should be allowed; all other requests should be rejected.

## Part 17

Add a new endpoint `GET /fibonacci/:n` which will return the nth number in the Fibonacci sequence (0,1,1,2,3,5,8,13 etc.) where 0 is the 1st number, and 13 is the 8th number.

## Part 18

Extend the API authentication done earlier to also protect the fibonacci endpoint.

## Part 19

For each request that comes in, log out the following info to STDOUT. 

- timestamp
- HTTP method
- path and query params (`GET /fibonacci/8`, `POST /add?num=9&num=12`)
- first 8 chars of the API key
- request body size
- response code
- time taken to send the response

## Part 20

For all endpoints, support a query param `?flakiness=:p`, where 0 <= p <= 1

Based on that probability, return a response code of 500 for that request.  

e.g. if a request comes in as `?flakiness=1`, it's a 100% probability, so send a 500.
For `?flakiness=0.2`, the probability should be 1 in 5 of returning a 500.

## Part 21

Extend the flakiness param to specify the flaky response code.

e.g. `?flakiness=0.2,404` should return a 404 with a probability of 20%

## Part 22

Extend the flakiness param to simulate a slow server by specifying a delay.

e.g. `?flakiness=0.33,500,3s` will introduce a delay of 3 seconds before sending a 500, all with a probability of 33%

Support `s` and `ms` as units for the delay

## Part 23

Build a program `fiboclient` that calls the fibonacci endpoint (running in a separate process) for the value `10`.


