# GOWNLOAD
command line utility to obtain the hash from url file

# Requirements
* A URL
* A destination path for a file containing a hash
* An optional value for throttling the download Given a url

### Build From Source

1. [Install Go](https://golang.org/doc/install)
2. Clone the repository:

   ```shell
   git clone https://github.com/garciaolais/gownload.git
   ```

3. Run `make` from the source directory

   ```shell
   cd gownload
   make build
   ./gownload -url https://raw.githubusercontent.com/garciaolais/gownload/main/cmd/file12.dat
   ```
### Run Tests

```shell
   cd gownload
   make test
    go test -v cmd/cmd_test.go -short
    === RUN   TestHash
    hash - data [12]
    hash - data result [24 108 90 204 81 189 102 126]
    --- PASS: TestHash (0.00s)
    === RUN   TestCmd
    download - http://localhost:8080/file12.dat
    hash - data [12]
    hash - data result [24 108 90 204 81 189 102 126]   
    create file /tmp/file.dat
    file hex - 186c5acc51bd667e--- PASS: TestCmd (0.01s)
    PASS
    ok      command-line-arguments  0.009s
    go test -v util/util_test.go -short
    === RUN   TestIsURL        
    --- PASS: TestIsURL (0.00s)
    PASS
    ok      command-line-arguments  0.004s
```
### Getting Started

See usage with:
```shell   
    gownload -help
        -path string      
            a destination path for a file containing a hash (default "/tmp/file.dat")
        -throttling
            throttling download
        -url string
            (default "https://raw.githubusercontent.com/garciaolais/gownload/main/cmd/file12.dat")
```