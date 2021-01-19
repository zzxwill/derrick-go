package core

const Platform = "BaseRigging"

type Rigging interface {
	Detect(workspace string) (bool, string)
}
