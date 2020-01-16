# 這是註解
# target:dependencies
#	命令

all:  win64 win86 linux64 linux86 prepareStasticFiles

clean:
	rm -rf distributions
	rm -rf archivements

prepareStasticFiles: distributions archivements
	cp -r static/ distributions/static
	cp -r view/ distributions/view

distributions:
	@mkdir distributions

archivements:
	@mkdir archivements

tar: all
	tar cvf archivements\OnePageNote_win64.tar.gz distributions\view distributions\static distributions\OnePageNote_win64.exe
	tar cvf archivements\OnePageNote_win386.tar.gz distributions\view distributions\static distributions\OnePageNote_win386.exe
	tar cvf archivements\OnePageNote_linux64.tar.gz distributions\view distributions\static distributions\OnePageNote_linux64.exe
	tar cvf archivements\OnePageNote_linux386.tar.gz distributions\view distributions\static distributions\OnePageNote_linux386.exe

win86: prepareStasticFiles
	set GOOS=windows
	set GOARCH=386
	go build -o distributions/OnePageNote_win386.exe ./cmd/desktopApp/main.go

win64: prepareStasticFiles
	set GOOS=windows
	set GOARCH=amd64
	go build -o distributions/OnePageNote_win64.exe ./cmd/desktopApp/main.go

linux64: prepareStasticFiles
	set GOOS=linux
	set GOARCH=amd64
	go build -o distributions/OnePageNote_linux64.exe ./cmd/desktopApp/main.go

linux86: prepareStasticFiles
	set GOOS=linux
	set GOARCH=386
	go build -o distributions/OnePageNote_linux386.exe ./cmd/desktopApp/main.go
