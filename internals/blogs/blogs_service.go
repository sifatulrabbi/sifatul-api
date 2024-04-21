package blogs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sifatulrabbi/sifatul-api/internals/caching"
)

type ArticleItemBody struct {
	ContentType string `json:"category_type"`
	Content     string `json:"content"`
}

type ArticleItem struct {
	ID        string            `json:"id"`
	Title     string            `json:"title"`
	Summary   string            `json:"summary"`
	Category  string            `json:"category"`
	Tags      []string          `json:"tags"`
	Body      []ArticleItemBody `json:"body"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type ArticleEntry struct {
	ID       string `json:"id"`
	Metadata string `json:"metadata"`
}

type ArticleEntries []ArticleEntry

func (be *ArticleEntry) GetFullBlog() (*ArticleItem, error) {
	return nil, nil
}

type IBlogService interface {
	GetAllArticleEntries() (*ArticleEntries, error)
	QueryArticles(q, t string) (*ArticleEntries, error)
	FindArticleById(id string) (*ArticleItem, error)
}

type CachedBlogService struct {
	cachingService caching.ICachingService
}

var _ IBlogService = &CachedBlogService{}

func NewCachedBlogService() CachedBlogService {
	s := CachedBlogService{
		cachingService: caching.NewCustomExpiringCachingService(time.Hour * 1),
	}
	return s
}

func (s *CachedBlogService) GetAllArticleEntries() (*ArticleEntries, error) {
	blogEntries := ArticleEntries{}
	url := "https://raw.githubusercontent.com/sifatulrabbi/blogs/main/index.json"

	if cachedEntries, err := s.cachingService.Get(url); err != nil {
		log.Println("Error while getting cached data:", err)
	} else if d, ok := cachedEntries.(ArticleEntries); ok {
		return &d, nil
	} else {
		log.Println("corrupted cached data:", d, err)
	}

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&blogEntries); err != nil {
		return nil, err
	}

	fmt.Println("caching article entries")
	if err = s.cachingService.Set(url, blogEntries); err != nil {
		log.Println("failed to cache article entries:", err)
	}
	return &blogEntries, nil
}

func (s *CachedBlogService) QueryArticles(q, t string) (*ArticleEntries, error) {
	return nil, nil
}

func (s *CachedBlogService) FindArticleById(id string) (*ArticleItem, error) {
	return nil, nil
}