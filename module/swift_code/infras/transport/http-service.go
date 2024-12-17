package transport

type CrawlingDataUseCase interface {
	Execute(pageUrl string, domainUrl string) error
}

type HttpService struct {
	crawlingDataUseCase CrawlingDataUseCase
}

func NewHttpService(useCase CrawlingDataUseCase) *HttpService {
	return &HttpService{
		crawlingDataUseCase: useCase,
	}
}
