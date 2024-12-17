package swift_code

import (
	"data_crawler/module/swift_code/infras/repository"
	"data_crawler/module/swift_code/infras/transport"
	usecase "data_crawler/module/swift_code/usecase"
	"github.com/gin-gonic/gin"
)

func SetupService(router *gin.RouterGroup) {
	repo := &repository.MySQLStorage{}

	crawlUsecase := usecase.NewCrawlingDataUseCase(repo)

	crawlService := transport.NewHttpService(crawlUsecase)

	//router.POST("/crawl/:page/:country/:domain", func(c *gin.Context) {
	//	crawlService.Crawl(c)
	//	fmt.Println("module file check")
	//})

	router.POST("/crawl/:page/:country/:domain/:nth-page", crawlService.Crawl)

}
