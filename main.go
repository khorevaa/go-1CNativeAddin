package main

import "unsafe"

//GetClassNames
//
//Синтаксис:
//
//Копировать в буфер обмена
//const WCHAR_T* GetClassNames()
//
//Возвращаемое значение:
//
//const WCHAR_T*
//
//Описание:
//
//Получение списка имен объектов компоненты.
func GetClassNames() {

}

//GetClassObject
//
//Синтаксис:
//
//Копировать в буфер обмена
//long GetClassObject(const WCHAR_T* clsName, IComponentBase** pIntf)
//
//Параметры:
//
//<clsName> Тип: const WCHAR_T*. Имя создаваемого объекта.
//<pIntf> Тип: IComponentBase**. Указатель на переменную, в которую нужно записать адрес вновь созданного объекта.
//Возвращаемое значение:
//
//long – не нулевое значение сигнализирует о успешном создании объекта
//Описание:
//
//Создание экземпляра объекта компоненты. Если объект не может быть создан или не найден объект с указанным именем – возвращается 0.

func GetClassObject(clsName string, pIntf *IComponentBase) int64 {

	return 0
}

//DestroyObject
//
//Синтаксис:
//
//Копировать в буфер обмена
//long DestroyObject(IComponentBase** pIntf)
//
//Параметры:
//
//<pIntf> Тип: IComponentBase**. Указатель на объект компонеты.
//Возвращаемое значение:
//
//long – успешность выполнения
//Описание:
//
//Удаление экземпляра ранее созданного объекта.
//Компонента должна своими средствами удалить объект и освободить используемую им память.
//При успешном завершении возвращается 0, иначе – код ошибки (Runtime error).
func DestroyObject(pIntf *ComponentObject) int64 {

	return 0
}

//SetPlatformCapabilities
//
//Синтаксис:
//
//Копировать в буфер обмена
//AppCapabilities SetPlatformCapabilities(const AppCapabilities capabilities)
//
//Параметры:
//
//<capabilities> Тип: перечисление AppCapabilities. Значения перечисления:
//eAppCapabilitiesInvalid = -1, eAppCapabilities1 = 1, eAppCapabilitiesLast = eAppCapabilities1,
//Описание:
//
//Устанавливает версию поддерживаемых платформой возможностей.
//Компонента должна вернуть версию, с которой она может работать.
//Если функция не реализована, то для компоненты не будут доступны возможности вывода сообщений, запроса информации о платформе.

func SetPlatformCapabilities(pIntf *ComponentObject) int {

	return 1

}

type IComponentBase unsafe.Pointer

type ComponentObject interface {
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

func main() {

}
