package blogs

import (
	"fmt"
	"testing"
	"time"
)

func TestDateParsing(t *testing.T) {
	if loc, err := time.LoadLocation("Asia/Dhaka"); err != nil {
		t.Error(err)
	} else {
		nowDhaka := time.Now().In(loc)
		fmt.Println("Current time in Dhaka:", nowDhaka.Format(time.RFC1123))
	}
}

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

	e := (*entries)[0]
	t.Log(e.CreatedAt)
}
