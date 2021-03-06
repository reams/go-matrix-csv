# League Backend Challenge

In `src/main.go` you will find a basic web server written in GoLang. It accepts some requests to perform the following operations

#### Given an uploaded csv file

```
1,2,3
4,5,6
7,8,9
```

#### 1. Echo (given) `/echo`

- Return the matrix as a string in matrix format.

```
// Expected output
1,2,3
4,5,6
7,8,9
```

#### 2. Invert `/invert`

- Return the matrix as a string in matrix format where the columns and rows are inverted

```
// Expected output
1,4,7
2,5,8
3,6,9
```

#### 3. Flatten `/flatten`

- Return the matrix as a 1 line string, with values separated by commas.

```
// Expected output
1,2,3,4,5,6,7,8,9
```

#### 4. Sum `/sum`

- Return the sum of the integers in the matrix

```
// Expected output
45
```

#### 5. Multiply `/multiply`

- Return the product of the integers in the matrix

```
// Expected output
362880
```

The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.

#### Compile and run

```
git clone https://github.com/ivoneijr/go-matrix-csv.git
cd go-matrix-csv
go build -o ./build/main src/main.go

./build/main
```

##### Run web server using `go run`

```
git clone https://github.com/ivoneijr/go-matrix-csv.git
cd go-matrix-csv
go run src/*.go
```

#### cURL tests

##### Request examples:

```
curl -F 'file={CSV_PATH}' "localhost:8080/echo"
curl -F 'file={CSV_PATH}' "localhost:8080/invert"
curl -F 'file={CSV_PATH}' "localhost:8080/flatern"
curl -F 'file={CSV_PATH}' "localhost:8080/sum"
curl -F 'file={CSV_PATH}' "localhost:8080/multiply"
```

`CSV_PATH` is your local csv file path

Remember, the form-data key should be named as `file`.

You can use samples in `/docs/sample`

#### Coverage

You can see coverage test report runing `go tool cover -html=coverage/report.out` inside project file (go-matrix-csv)

### What we're looking for

- The solution runs
- The solution performs all cases correctly
- The code is easy to read
- The code is reasonably documented
- The code is tested
- The code is robust and handles invalid input and provides helpful error messages

Have fun 🤓
