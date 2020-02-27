@echo off

rem rd bin /s /q
rem mkdir bin\32
rem mkdir bin\64

set fn=addin
set gn=main.go
set dn=main.def

set tmp=%path%
::使跨平台编译使用CGO
set CGO_ENABLED=1

::设置32位GCC
rem set path=%tmp%;C:\msys64\mingw32\bin
::设置目标平台(32位)
rem set GOARCH=386
::编译32位静态库
rem go build -buildmode=c-archive -o bin\32\%fn%.a %gn%
::生成DLL,LIB文件
rem C:\msys64\mingw32\bin\gcc.exe %dn% bin\32\%fn%.a -shared -lwinmm -lWs2_32 -o bin\32\%fn%-32.dll -Wl,--out-implib,bin\32\%fn%.lib

::设置64位GCC
rem set path=%tmp%;C:\msys64\mingw64\bin
::设置目标平台(64位)
rem set GOARCH=amd64
::编译64位静态库
rem go build -buildmode=c-archive -o bin\64\%fn%.a %gn%
::生成DLL,LIB文件
rem C:\msys64\mingw64\bin\gcc.exe %dn% bin\64\%fn%.a -shared -lwinmm -lWs2_32 -o bin\64\%fn%-64.dll -Wl,--out-implib,bin\64\%fn%.lib

set GOARCH=amd64

set path=C:\msys64\mingw64\bin;%tmp%

go build -buildmode=c-shared -o bin\%fn%-64.dll .

set GOARCH=386

set path=C:\msys64\mingw32\bin;%tmp%

go build -buildmode=c-shared -o bin\%fn%-32.dll .