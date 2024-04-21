package blogs

import (
	"fmt"
	"testing"
)

func TestGetBlogEntries(t *testing.T) {
	cbs := NewCachedBlogService()
	entries, err := cbs.GetAllArticleEntries()
	fmt.Println(entries)
	if err != nil {
		t.Fatal(err)
	}
	if len(*entries) < 1 {
		t.Fatal("incorrect amount of entries")
	}
}
