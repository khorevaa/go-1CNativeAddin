package main

import (
	"./internal/types"
	"unsafe"
)

type NativeAddinWrapper interface {
	GetClassNames() types.Wchar_t
	GetClassObject(clsName types.Wchar_t, pIntf *unsafe.Pointer) int64
	DestroyObject(pIntf *unsafe.Pointer) int64
	SetPlatformCapabilities(pIntf *unsafe.Pointer) int64
}

type ClassBaseWrapper interface {
	SetPlatformCapabilities() int64

	//func Init([MarshalAs(UnmanagedType.IDispatch)] object connection)
	//func Done();
	//func GetInfo([MarshalAs(UnmanagedType.SafeArray, SafeArraySubType = VarEnum.VT_VARIANT)] ref object[] info);
	//
	//void RegisterExtensionAs([MarshalAs(UnmanagedType.BStr)] ref String extensionName);
	//void GetNProps(ref Int32 props);
	//void FindProp([MarshalAs(UnmanagedType.BStr)] String propName, ref Int32 propNum);
	//void GetPropName(Int32 propNum, Int32 propAlias, [MarshalAs(UnmanagedType.BStr)] ref String propName);
	//void GetPropVal(Int32 propNum, [MarshalAs(UnmanagedType.Struct)] ref object propVal);
	//void SetPropVal(Int32 propNum, [MarshalAs(UnmanagedType.Struct)] ref object propVal);
	//void IsPropReadable(Int32 propNum, ref bool propRead);
	//void IsPropWritable(Int32 propNum, ref Boolean propWrite);
	//void GetNMethods(ref Int32 pMethods);
	//void FindMethod([MarshalAs(UnmanagedType.BStr)] String methodName, ref Int32 methodNum);
	//void GetMethodName(Int32 methodNum, Int32 methodAlias, [MarshalAs(UnmanagedType.BStr)] ref String methodName);
	//void GetNParams(Int32 methodNum, ref Int32 pParams);
	//void GetParamDefValue(Int32 methodNum, Int32 paramNum, [MarshalAs(UnmanagedType.Struct)] ref object paramDefValue);
	//void HasRetVal(Int32 methodNum, ref Boolean retValue);
	//void CallAsProc(Int32 methodNum, [MarshalAs(UnmanagedType.SafeArray, SafeArraySubType = VarEnum.VT_VARIANT)] ref object[] pParams);
	//void CallAsFunc(Int32 methodNum, [MarshalAs(UnmanagedType.Struct)] ref object retValue, [MarshalAs(UnmanagedType.SafeArray, SafeArraySubType = VarEnum.VT_VARIANT)] ref object[] pParams);

}

type CreateClassObject func() CreateClassObject

type NativeObject struct {
	namespace string
	classes   map[string]CreateClassObject
	objects   []ClassBaseWrapper
}

func (obj *NativeObject) AddClass(name string, f CreateClassObject) {

	obj.classes[name] = f
}

func (obj *NativeObject) GetClassNames() types.Wchar_t {

	str, _ := types.StringToWcharT(obj.namespace)
	return str

}

func (obj *NativeObject) GetClassObject(clsName types.Wchar_t, pIntf *unsafe.Pointer) int64 {

	className, err := types.WcharTToString(clsName)

	if err != nil {
		return 1
	}
	createObject := obj.classes[className]
	newObject := createObject()

	pIntf = (*unsafe.Pointer)(unsafe.Pointer(&newObject))

	return 0

}

func (obj *NativeObject) DestroyObject(pIntf *unsafe.Pointer) int64 {
	// TODO Удаление созданного класса
	return 0

}

func (obj *NativeObject) SetPlatformCapabilities(pIntf *unsafe.Pointer) int64 {

	//SetPlatformCapabilities(pIntf *unsafe.Pointer) int64

	IObject := (interface{})(pIntf)
	classObject := IObject.(ClassBaseWrapper)
	return classObject.SetPlatformCapabilities()

}
