package core

import (
	"reflect"
)

type ExtensionPoint struct {
	Rigging Rigging
}

func getRiggingName(platform interface{}) string {
	return reflect.TypeOf(platform).Name()
}

func loadRigging(detect interface{}) reflect.Value {
	return reflect.ValueOf(detect)
}

func Register(rig Rigging) ExtensionPoint {
	return ExtensionPoint{
		Rigging: rig,
	}
}
