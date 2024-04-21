package blogs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(r *gin.RouterGroup) {
	r.GET("/blogs/articles/all", getAllArticleEntries) // this supports `q="search query"` and `t="tag1,tag2"`
	r.GET("/blogs/articles/single/:id", getArticleById)
	r.GET("/blogs/articles/category/:category", getArticleByCategory)
	r.GET("/blogs/categories", getAllCategories)
	r.GET("/blogs/tags", getAllTags)
}

func getAllArticleEntries(c *gin.Context) {
	var (
		entries *ArticleEntries
		err     error
	)
	blogsService := NewCachedBlogService()

	q := c.Query("q")
	t := c.Query("t")
	if q != "" {
		entries, err = blogsService.QueryArticles(q, t)
	} else {
		entries, err = blogsService.GetAllArticleEntries()
	}
	if err != nil || entries == nil {
		errMsg := ""
		if err != nil {
			errMsg = err.Error()
		} else {
			errMsg = "No blog entries found!"
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Unable to get the blogs at the moment",
			"error":   errMsg,
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Articles found",
		"data":    entries,
	})
}

func getAllCategories(c *gin.Context) {
}

func getAllTags(c *gin.Context) {
}

func getArticleById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "'id' is not found."})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Blog found",
		"data":    "",
	})
}

func getArticleByCategory(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "'id' is not found."})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Blog found",
		"data":    "",
	})
}
