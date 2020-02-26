package main

import "./internal/types"

type NativeAddin interface {
	GetClassNames() types.Wchar_t
	GetClassObject(clsName types.Wchar_t, pIntf *IComponentBase) int64
}

type NativeAddinObject struct {
	classNames string
}
