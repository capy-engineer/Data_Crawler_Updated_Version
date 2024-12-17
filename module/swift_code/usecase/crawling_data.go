package usecase

import (
	"data_crawler/module/swift_code/model"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"time"
)

type CreateRepository interface {
	CreateInfoSwiftCode(payload []model.Swiftcode) error
}
type CrawlingDataUseCase struct {
	repo CreateRepository
}

func NewCrawlingDataUseCase(repo CreateRepository) *CrawlingDataUseCase {
	return &CrawlingDataUseCase{
		repo: repo,
	}
}

func (uc *CrawlingDataUseCase) Execute(pageUrl string, domain string) error {
	c := colly.NewCollector(
		//colly.AllowedDomains("www.theswiftcodes.com"),
		colly.AllowedDomains(domain),
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
		//colly.Debugger(&debug.LogDebugger{Output: os.Stdout}),
	)
	fmt.Println("done")
	var payload []model.Swiftcode
	// Extract the data-value attribute and text content
	c.OnHTML("tbody", func(e *colly.HTMLElement) {

		//var result []string

		bankId := e.Attr("table-id")
		bankName := e.Attr("table-name")
		bankCity := e.Attr("table-city")
		bankBranch := e.Attr("table-branch")
		bankSwift := e.Attr("table-swift")

		e.ForEach("tr", func(i int, row *colly.HTMLElement) {

			bankIdText := row.ChildText("td:nth-child(1)")
			bankNameText := row.ChildText("td:nth-child(2)")
			bankCityText := row.ChildText("td:nth-child(3)")
			bankBranchText := row.ChildText("td:nth-child(4)")
			bankSwiftText := row.ChildText("td:nth-child(5)")

			payload = append(payload, model.Swiftcode{
				Id:                 bankIdText,
				BankInsitutionName: bankNameText,
				City:               bankCityText,
				Branch:             bankBranchText,
				SwiftCode:          bankSwiftText,
				UpdatedAt:          time.Now().Format("2006-01-02 15:04:05"),
				CreatedAt:          time.Now().Format("2006-01-02 15:04:05"),
			})
		})
		//bankIdText := e.Text
		//bankNameText := e.Text
		//bankCityText := e.Text
		//bankBranchText := e.Text
		//bankSwiftText := e.Text

		fmt.Println("cai nay la attribute nha", bankId, bankName, bankCity, bankBranch, bankSwift)

		//result = append(result, bankIdText, bankNameText, bankCityText, bankBranchText, bankSwiftText)

		//utils := shared.Utils{}
		//err := shared.Utils.WriteCSV(result)

		//fmt.Println("result: ", bankIdText)
		//err := shared.WriteCSV(result)

		if err := uc.repo.CreateInfoSwiftCode(payload); err != nil {
			log.Fatalf("Failed to add payload into db: %v", err)
		}
		//if err != nil {
		//	log.Fatalf("Failed to write CSV: %v", err)
		//}
	})

	//defer func() {
	//	if err := uc.repo.CreateInfoSwiftCode(payload); err != nil {
	//		log.Fatalf("Failed to add payload into db: %v", err)
	//	}
	//}()

	// Handle errors
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Request URL: %s failed with response: %v\nError: %v", r.Request.URL, r, err)
	})

	// Start the scraping process
	//err := c.Visit("https://www.theswiftcodes.com/vietnam/")
	err := c.Visit(pageUrl)
	if err != nil {
		log.Fatalf("Fatal error visiting the URL: %v", err)
	}
	return nil
}
