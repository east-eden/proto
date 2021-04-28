@echo off
set outDir=%~1

for %%i in (global/*.proto) do (
    echo gen %%i
    echo.|protoc.exe --proto_path="global" --csharp_wxb_out=%outDir%  global/%%i
)
::echo finish... 
::pause