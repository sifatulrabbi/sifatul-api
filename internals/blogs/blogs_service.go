package blogs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sifatulrabbi/sifatul-api/internals/caching"
)

const StorageRootUrl = "https://raw.githubusercontent.com/sifatulrabbi/blogs/main"

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
	ID        string    `json:"id"`
	Url       string    `json:"url"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	Category  string    `json:"category"`
	Tags      []string  `json:"tags"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	url := fmt.Sprintf("%s/index.json", StorageRootUrl)

	if cachedEntries, err := caching.RetrieveCachedData[ArticleEntries](s.cachingService, url); err != nil {
		log.Println("Error while retrieving cached data:", err)
	} else if len(cachedEntries) > 0 {
		return &cachedEntries, nil
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

	if err = s.cachingService.Set(url, blogEntries); err != nil {
		log.Println("failed to cache article entries:", err)
	} else {
		fmt.Println("caching article entries")
	}
	return &blogEntries, nil
}

func (s *CachedBlogService) QueryArticles(q, t string) (*ArticleEntries, error) {
	return nil, nil
}

func (s *CachedBlogService) FindArticleById(id string) (*ArticleItem, error) {
	allEntries, err := s.GetAllArticleEntries()
	if err != nil {
		return nil, err
	}
	var entry *ArticleEntry = nil
	for _, e := range *allEntries {
		if e.ID == id {
			entry = &e
			break
		}
	}
	if entry == nil {
		return nil, nil
	}

	url := fmt.Sprintf("%s%s", StorageRootUrl, entry.Url)

	if cachedData, err := caching.RetrieveCachedData[ArticleItem](s.cachingService, url); err != nil {
		log.Println("Error while getting the cached data:", err)
	} else if cachedData.ID != "" {
		return &cachedData, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	article := &ArticleItem{}
	if err := json.NewDecoder(res.Body).Decode(article); err != nil {
		return nil, err
	}
	s.cachingService.Set(url, *article)
	return article, nil
}

func (s *CachedBlogService) GetAllCategories() ([]string, error) {
	url := fmt.Sprintf("%s/categories/index.json", StorageRootUrl)

	if cachedData, err := caching.RetrieveCachedData[[]string](s.cachingService, url); err != nil {
		log.Println("Error while getting the cached data:", err)
	} else if len(cachedData) > 0 {
		return cachedData, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	categories := []string{}
	if err := json.NewDecoder(res.Body).Decode(&categories); err != nil {
		return nil, err
	}
	s.cachingService.Set(url, categories)
	return categories, nil
}

func (s *CachedBlogService) GetAllTags() ([]string, error) {
	url := fmt.Sprintf("%s/tags/index.json", StorageRootUrl)

	if cachedData, err := caching.RetrieveCachedData[[]string](s.cachingService, url); err != nil {
		log.Println("Error while getting the cached data:", err)
	} else if len(cachedData) > 0 {
		return cachedData, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	tags := []string{}
	if err := json.NewDecoder(res.Body).Decode(&tags); err != nil {
		return nil, err
	}
	s.cachingService.Set(url, tags)
	return tags, nil
}
