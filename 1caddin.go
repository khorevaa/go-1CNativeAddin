package main

import (
	"./internal/wrapper"
)

type NativeComponent struct {
	wrapper.NativeObjectWrapper
}

func NewNativeComponent(name string) *NativeComponent {

	return &NativeComponent{
		wrapper.NativeObjectWrapper{Namespace: name},
	}
}
