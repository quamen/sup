# Sup!

Track activity across your github organisation

## Development

Sup is built using:

* A [Golang](http://golang.org) based server component

### Dependencies

* A working [Golang](http://golang.org/doc/install) environment
* A working ruby environment (to use the Procfile)

### Getting setup

1. `$ go get github.com/quamen/sup`
2. `$ cd $GOPATH/src/github.com/quamen/sup`
3. `$ script/bootstrap`
4. Set up your `.env` file
5. `$ foreman start`

## Inspiration and blatent code lifting

The following two repos were super helpful in writing this, and I've definitely stolen ideas and code from both while learning my way around go.

* https://github.com/google/go-github
* https://github.com/kljensen/golang-html5-sse-example
