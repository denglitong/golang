// `ginkgo bootstrap`
// generate module test entrypoint
package books_test

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"testing"
)

// Entrypoint for Ginkgo specs to run,
// we must include this otherwise we got `testing: warning: no tests to run`.
func TestBooks(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	// RunSpecs is the entry point for the Ginkgo spec runner.
	ginkgo.RunSpecs(t, "Books Suite")
}
