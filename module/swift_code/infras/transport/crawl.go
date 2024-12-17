package transport

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (https *HttpService) Crawl(c *gin.Context) {
	//return func(c *gin.Context) {

	pageUrl := c.Param("page")
	countries := c.Param("country")
	page := c.Param("nth-page")
	domainUrl := c.Param("domain")

	var completedPageUrl string
	fmt.Println("page nha:", page)
	// page > 1
	if page != "0" {
		completedPageUrl = "https://" + pageUrl + "/" + countries + "/page/" + page + "/"
	} else {
		completedPageUrl = "https://" + pageUrl + "/" + countries + "/"
	}

	//pageUrlEncode, _ := url.QueryUnescape(pageUrl)
	fmt.Println("pageUrl: ", completedPageUrl)
	fmt.Println("domainUrl: ", domainUrl)

	if pageUrl == "" || domainUrl == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "page_url and domain_url are required"})
		return
	}

	if err := https.crawlingDataUseCase.Execute(completedPageUrl, domainUrl); err != nil {
		panic(err)
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "success"})
	fmt.Println("Crawling data from ", completedPageUrl, " to ", domainUrl)
	//}
}
