package loaders

type Loader interface {
	Load() error
	GetMap() (map[string]interface{}, error)
}
