package assets

//go:generate file2byteslice -package=images -input=../_assets/images/koma_taiki.png -output=./images/koma_taiki.go -var=Koma_taiki_png
//go:generate file2byteslice -package=images -input=../_assets/images/koma_00.png -output=./images/koma_00.go -var=Koma_00_png
//go:generate file2byteslice -package=images -input=../_assets/images/koma_01.png -output=./images/koma_01.go -var=Koma_01_png
//go:generate file2byteslice -package=images -input=../_assets/images/koma_02.png -output=./images/koma_02.go -var=Koma_02_png
//go:generate file2byteslice -package=images -input=../_assets/images/koma_03.png -output=./images/koma_03.go -var=Koma_03_png
//go:generate file2byteslice -package=images -input=../_assets/images/title_bg.png -output=./images/title_bg.go -var=Title_bg_png
//go:generate gofmt -s -w .
