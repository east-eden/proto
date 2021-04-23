@echo off
set outDir=%~1
for %%i in (global/*.proto) do (
   #echo gen %%~nxi...
   protoc.exe --csharp_out=%outDir%  global/%%~nxi)

echo finish... 
#pause