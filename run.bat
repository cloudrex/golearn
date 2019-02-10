@echo off
go build
del bin\golearn.exe
move golearn.exe bin/golearn.exe >NUL
cd bin
golearn.exe
cd ..
