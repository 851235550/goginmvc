
### Install dependencies
```
make install
```

### Install ginkgo
```
go install github.com/onsi/ginkgo/ginkgo
```

### Build
```
make build env=dev   # dev 
make build env=prod  # prod
```

### Run unit tests
```
make test
```

### Others
```
make clean     # clean

make fmt       # format code
make vet       # check code
```