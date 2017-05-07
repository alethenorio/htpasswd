@echo off
set NAME=htpasswd
set GOOS=linux
set GOARCH=amd64

set BINARY=dist/%NAME%-%GOOS%-%GOARCH%

for /f %%i in ('git describe --tags') do set VERSION=%%i

for /f %%i in ('git rev-parse HEAD') do set COMMITID=%%i

for /F "usebackq tokens=1,2 delims==" %%i in (`wmic os get LocalDateTime /VALUE 2^>NUL`) do if '.%%i.'=='.LocalDateTime.' set ldt=%%j
set ldt=%ldt:~0,4%-%ldt:~4,2%-%ldt:~6,2%_%ldt:~8,2%:%ldt:~10,2%:%ldt:~12,6%
set BUILDTIME=%ldt%



echo Building version=%VERSION% from commit=%COMMITID% for %GOOS%/%GOARCH%

set CGO_ENABLED=0
go build -a -installsuffix cgo -o "%BINARY%" -ldflags "-X main.version=%VERSION% -X main.buildTime=%BUILDTIME% -X main.commitId=%COMMITID%"

set GOOS=windows
set GOARCH=amd64
set BINARY=dist/%NAME%-%GOOS%-%GOARCH%.exe

echo Building version=%VERSION% from commit=%COMMITID% for %GOOS%/%GOARCH%
go build -o "%BINARY%" -ldflags "-X main.version=%VERSION% -X main.buildTime=%BUILDTIME% -X main.commitId=%COMMITID%"