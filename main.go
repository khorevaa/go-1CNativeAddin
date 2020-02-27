package main

import "unsafe"
import "./internal/wrapper"

// #include <./include/types.h>
// //typedef long GoInt64;
import "C"

var component *NativeComponent

func main() {

	component = NewNativeComponent("khorevaa.golang")

}

//export GetClassNames
func GetClassNames() *C.wchar_t {

	classNames, _ := wrapper.StringToWcharT("khorevaa.golang.test")
	return (*C.wchar_t)(classNames)

}

//export GetClassObject
func GetClassObject(clsName *C.WCHAR_T, pIntf *unsafe.Pointer) int64 {

	return 0 //component.GetClassObject(clsName, pIntf)

}

//export DestroyObject
func DestroyObject(pIntf *unsafe.Pointer) int64 {

	return 0

}
