@echo off
set outDir=%~1
set input=%~2
::echo gen %input%
echo.|protoc.exe --proto_path="global" --csharp_wxb_out=%outDir%  global/%input%
::echo finish... 

::pause