package shared

type Config interface {
	GetURL() string
}

type env struct {
}

func (env *env) GetURL() string {
	return "haha"
}
