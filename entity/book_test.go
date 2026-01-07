package entity
 
import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestBookCorrect(t *testing.T) {
	g := NewGomegaWithT(t)

	book := Book{
		Title : "Test Book",
		Author : "Test Author",
		Price : 100,
	}
	ok, err := govalidator.ValidateStruct(book)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}

func TestBookTitleRequired(t *testing.T) {
	g := NewGomegaWithT(t)

	book := Book{
		Title : "",
		Author : "Test Author",
		Price : 100,
	}
	ok, err := govalidator.ValidateStruct(book)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Title is required"))
}

func TestBookAuthorRequired(t *testing.T) {
	g := NewGomegaWithT(t)

	book := Book{
		Title : "Test Book",
		Author : "",
		Price : 100,
	}
	ok, err := govalidator.ValidateStruct(book)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Author is required"))
}

func TestBookPriceRange(t *testing.T) {
	g := NewGomegaWithT(t)
	book := Book{
		Title : "Test Book",
		Author : "Test Author",
		Price : 0,
	}
	ok, err := govalidator.ValidateStruct(book)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Price is required"))
}

