package external

type Finder interface {
	Process(flags interface{}) error
}
