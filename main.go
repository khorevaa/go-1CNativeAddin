package main

import "unsafe"
import "./internal/types"

var component *NativeComponent

func main() {

	component = NewNativeComponent("khorevaa.golang")

}

func GetClassNames() types.Wchar_t {

	return component.GetClassNames()

}

func GetClassObject(clsName types.Wchar_t, pIntf *unsafe.Pointer) int64 {

	return component.GetClassObject(clsName, pIntf)

}

func DestroyObject(pIntf *unsafe.Pointer) int64 {

	return component.DestroyObject(pIntf)

}

func SetPlatformCapabilities(pIntf *unsafe.Pointer) int64 {

	return component.DestroyObject(pIntf)

}
