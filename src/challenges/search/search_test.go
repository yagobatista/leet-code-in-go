package search

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSearch(t *testing.T) {
	suite.Run(t, new(SearchSuite))
}

type SearchSuite struct {
	suite.Suite
}

func (this *SearchSuite) TestExample1() {
	expected := []string{"Word", "Microsoft Word", "Wordpress", "1Password", "Google AdWords"}
	output := search("word", []string{"Google AdWords", "Wordpress", "Word", "1Password", "Microsoft Word"})

	this.Equal(expected, output)
}

func (this *SearchSuite) TestOneWord() {
	expected := []string{"Word"}
	output := search("word", []string{"Word"})

	this.Equal(expected, output)
}

func (this *SearchSuite) TestOneExactMatch() {
	expected := []string{"Microsoft Word"}
	output := search("word", []string{"Microsoft Word"})

	this.Equal(expected, output)
}

func (this *SearchSuite) TestStartWith() {
	expected := []string{"Wordpress"}
	output := search("word", []string{"Wordpress"})

	this.Equal(expected, output)
}

func (this *SearchSuite) TestOrder() {
	expected := []string{"1Password", "Google AdWords"}
	output := search("word", []string{"Google AdWords", "1Password"})

	this.Equal(expected, output)
}
