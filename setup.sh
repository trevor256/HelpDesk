#!\bin\bash

sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev
go get fyne.io/fyne/v2
go mod tidy