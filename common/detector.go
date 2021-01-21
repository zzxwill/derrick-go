package common

type Detector interface {
	Execute() (map[string]string, error)
	Name() string
}
