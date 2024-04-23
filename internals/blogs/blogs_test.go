package blogs

import (
	"fmt"
	"testing"
)

func TestGetBlogEntries(t *testing.T) {
	s := NewCachedBlogService()
	entries, err := s.GetAllArticleEntries()
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}
	if len(*entries) < 1 {
		t.Fatal("incorrect amount of entries")
	}
}

func TestParseMdText(t *testing.T) {
	s := NewCachedBlogService()
	mockId := "mock-article-on-tree-plantation"
	article, err := s.FindArticleById(mockId)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println(article.Body)
}
