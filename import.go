package main

/*
#include <wchar.h>
const size_t SIZEOF_WCHAR_T = sizeof(wchar_t);
void gowchar_set (wchar_t *arr, int pos, wchar_t val)
{
	arr[pos] = val;
}
wchar_t gowchar_get (wchar_t *arr, int pos)
{
	return arr[pos];
}
*/
import "C"
