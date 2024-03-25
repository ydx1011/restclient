package request

type Setter interface {
	Set(key string, value interface{})
}

type Opt func(Setter)
