@echo off
cd src
..\tools\fyne bundle icon.png > bundled.go
windres --output-format=coff -o icon.syso icon.rc
go build -o ../sudokusolver.exe -ldflags="-H windowsgui"
rem go build -o ../sudokusolver.exe 
cd ..