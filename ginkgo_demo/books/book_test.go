// `ginkgo generate book`
// generate specs for separate files
package books_test

import (
	"fmt"
	"github.com/denglitong/golang/ginkgo_demo/books"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {
	var book *books.Book
	now := time.Now()

	// setup nodes
	BeforeEach(func() {
		book = books.NewBook(
			"Les Miserables",
			"Victor Hugo",
			&now,
		)
	})

	// container nodes
	Describe("Description books", func() {
		// container nodes
		// Context is an alias for Describe
		Context("without title, display no title", func() {
			// setup nodes for this context only
			BeforeEach(func() {
				book.Title = ""
			})
			// subject nodes, to write a spec that makes assertions about the subject
			// Each It node corresponds to an individual Ginkgo spec.
			// You cannot nest any other Ginkgo nodes within an It node's closure.
			It("should display No Title", func() {
				Expect(book.Description()).To(Equal("No Title"))
			})
		})
		Context("with title, name", func() {
			// For each spec, Ginkgo will run the closures attached to any associated setup nodes
			// and then run the closure attached to the subject node.
			BeforeEach(func() {
				book.PublishedDate = nil
			})
			It("should display name and author", func() {
				Expect(book.Description()).To(Equal(fmt.Sprintf("%s, %s", book.Title, book.Author)))
			})
			// When is an alias for Describe
			When("when publishedDate is nil", func() {
				It("isNewBook should be false", func() {
					Expect(book.IsNewBook()).To(Equal(false))
				})
			})
		})
		// Context & When can have hierarchy architecture
		When("with normal title, author and publish date", func() {
			// assertions are only allowed in It nodes
			It("should display normal", func() {
				Expect(book.Description()).To(Equal(
					fmt.Sprintf("%s, %s %s", book.Title, book.Author, book.PublishedDate)),
				)
			})
			It("should be a new book", func() {
				Expect(book.IsNewBook()).To(Equal(true))
			})
		})
	})

	Describe("Yet another spec suite", func() {
		BeforeEach(func() {
			// context-specified setup
		})
		// specs...
	})
})
