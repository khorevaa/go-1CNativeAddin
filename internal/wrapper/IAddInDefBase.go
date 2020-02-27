package wrapper

import "../types"

type IAddInDefBaseWrapper interface {
	AddError(wcode int, source, descr types.Wchar_t, scode int64) types.ADDIN_BOOL

	Read(wszPropName types.Wchar_t, pVal uintptr, pErrCode int64, errDescriptor types.Wchar_t) types.ADDIN_BOOL
	Write(wszPropName types.Wchar_t, pVal uintptr, pErrCode int64, errDescriptor types.Wchar_t) types.ADDIN_BOOL

	RegisterProfileAs(wszProfileName types.Wchar_t) types.ADDIN_BOOL

	SetEventBufferDepth(lDepth int64) types.ADDIN_BOOL
	GetEventBufferDepth() int64

	ExternalEvent(wszSource, wszMessage, wszData types.Wchar_t) types.ADDIN_BOOL
	CleanEventBuffer()

	SetStatusLine(wszStatusLine types.Wchar_t) types.ADDIN_BOOL
	ResetStatusLine()
}

type IAddInDefBase struct {
}

//////////////////////////////////////////////////////////////////////////////////
///**
// * This class serves as representation of a platform for external
// * components External components use it to communicate with a platform.
// *
// */
///// Base interface for object components.
//class IAddInDefBase
//{
//public:
//virtual ~IAddInDefBase() {}
///// Adds the error message
///**
// *  @param wcode - error code
// *  @param source - source of error
// *  @param descr - description of error
// *  @param scode - error code (HRESULT)
// *  @return the result of
// */
//virtual bool ADDIN_API AddError(unsigned short wcode, const WCHAR_T* source,
//const WCHAR_T* descr, long scode) = 0;
//
///// Reads a property value
///**
// *  @param wszPropName -property name
// *  @param pVal - value being returned
// *  @param pErrCode - error code (if any error occured)
// *  @param errDescriptor - error description (if any error occured)
// *  @return the result of read.
// */
//virtual bool ADDIN_API Read(WCHAR_T* wszPropName,
//tVariant* pVal,
//long *pErrCode,
//WCHAR_T** errDescriptor) = 0;
///// Writes a property value
///**
// *  @param wszPropName - property name
// *  @param pVar - new property value
// *  @return the result of write.
// */
//virtual bool ADDIN_API Write(WCHAR_T* wszPropName,
//tVariant *pVar) = 0;
//
/////Registers profile components
///**
// *  @param wszProfileName - profile name
// *  @return the result of
// */
//virtual bool ADDIN_API RegisterProfileAs(WCHAR_T* wszProfileName) = 0;
//
///// Changes the depth of event buffer
///**
// *  @param lDepth - new depth of event buffer
// *  @return the result of
// */
//virtual bool ADDIN_API SetEventBufferDepth(long lDepth) = 0;
///// Returns the depth of event buffer
///**
// *  @return the depth of event buffer
// */
//virtual long ADDIN_API GetEventBufferDepth() = 0;
///// Registers external event
///**
// *  @param wszSource - source of event
// *  @param wszMessage - event message
// *  @param wszData - message parameters
// *  @return the result of
// */
//virtual bool ADDIN_API ExternalEvent(WCHAR_T* wszSource,
//WCHAR_T* wszMessage,
//WCHAR_T* wszData) = 0;
///// Clears event buffer
///**
// */
//virtual void ADDIN_API CleanEventBuffer() = 0;
//
///// Sets status line contents
///**
// *  @param wszStatusLine - new status line contents
// *  @return the result of
// */
//virtual bool ADDIN_API SetStatusLine(WCHAR_T* wszStatusLine) = 0;
///// Resets the status line contents
///**
// *  @return the result of
// */
//virtual void ADDIN_API ResetStatusLine() = 0;
//};
