# Sudoku Backend
The backend is written in Go and inspired by [eddycjy/go-gin-example](https://github.com/eddycjy/go-gin-example).

## Quick start
Run the API with Golang CLI:

```bash
$ go run main.go
```

Or, run the API with Docker:

```bash
$ docker build . -t sudoku-api && docker run -p 8080:8080 sudoku-api
```

The API is now available on http://127.0.0.1:8080.

To test the API, run:

```bash
$ curl 127.0.0.1:8080/api/v1/solve \
    -v \
    -X POST \
    -d '{"sudoku": [0,0,0,9,2,0,0,0,0,0,4,0,8,5,1,0,0,0,2,5,6,0,0,3,0,9,1,1,0,0,0,8,5,4,0,9,0,9,8,7,3,0,1,6,2,0,0,0,2,0,0,5,3,0,0,0,7,0,6,0,9,0,0,9,0,0,0,0,2,6,8,0,0,8,0,0,9,0,0,5,4]}'
```

## Advanced sudoku strategies

- https://www.sudokuonline.io/tips/advanced-sudoku-strategies
- https://www.kristanix.com/sudokuepic/sudoku-solving-techniques.php
