# eShopGo
Go parallel product handler

More info:

[Go official](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable)


### Set Environment variables
```export GOPATH=$HOME/myproject/```

```export GOBIN=$HOME/myproject/bin/```

```export GOROOT=/my/golang/root/install/dir/```

### Dependencies

To get dependencies just run:

``` go get ```

### Generate bin

```go install```


### Build and run

``` go run main.go ```



## Try the api

```curl "localhost:9000/products/?q={mysearch}"```
