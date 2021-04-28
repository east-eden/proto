@echo off
::set outDir=%~1
::set input=%~2
for %%i in (global/*.proto) do (
   echo gen %%i
   protoc.exe --proto_path="global" --csharp_wxb_out=output  input
)
go skip
for /f "delims=" %%d in ('dir /ad/b/x "global"') do (
   echo %%d
   for %%i in (global/%%d/*.proto) do (
	  echo gen %%d/%%i
	  protoc.exe --csharp_wxb_out=output  global/%%d/%%i
	)
)
:skip
echo finish... 
pause