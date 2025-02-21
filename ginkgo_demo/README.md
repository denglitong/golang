# Ginkgo testing framework

**Wiki**

- https://github.com/onsi/ginkgo/blob/master/README.md
- https://onsi.github.io/ginkgo/

**Install**

```shell
go install github.com/onsi/ginkgo/v2/ginkgo
go get github.com/onsi/ginkgo/v2/ginkgo
go get github.com/onsi/gomega/...
```

**Getting Started**

Ginkgo hooks into Go's existing `testing` infrastructure and use
the Ginkgo and Gomega DSLs instead of using `func TextX(t *testing.T)`.

We call a collection of Ginkgo specs in a given package a `Ginkgo suite`;
and we use the word `spec` to talk about individual Ginkgo tests contained in the suite.

Though they're functionally interchangeable, we'll use the word "spec" instead of "test" to make a distinction between Ginkgo tests and traditional testing tests.

In most Ginkgo suites there is only one TestX function - the entry point for Ginkgo.

**Run tests(specs)**

```shell
cd ./books
ginkgo
```

**Walk through & understand**

- [books_suite_test.go](./books/books_suite_test.go)
- [book.go](./books/book.go)
- [book_test.go](./books/book_test.go)
