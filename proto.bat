@echo off
set curDir=%~dp0
set globalDir=%curDir%global\
for /r global\ %%i in (*.proto) do (
   #echo gen %%~nxi
   protoc.exe --proto_path="global" --csharp_wxb_out=output  global\common\%%~nxi
)
   
echo finish... 
#pause