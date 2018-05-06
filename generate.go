package main

/* package images */
//go:generate file2byteslice -package=images -input=./_assets/images/title_bg.png -output=./assets/images/title_bg.go -var=Title_bg_png

// TODO: 黒菜のフレーム数を増やしながらリファインしたい
//go:generate file2byteslice -package=images -input=./_assets/images/kurona_taiki.png -output=./assets/images/kurona_taiki.go -var=Kurona_taiki_png
//go:generate file2byteslice -package=images -input=./_assets/images/kurona_00.png -output=./assets/images/kurona_00.go -var=Kurona_00_png
//go:generate file2byteslice -package=images -input=./_assets/images/kurona_01.png -output=./assets/images/kurona_01.go -var=Kurona_01_png
//go:generate file2byteslice -package=images -input=./_assets/images/kurona_02.png -output=./assets/images/kurona_02.go -var=Kurona_02_png

//go:generate file2byteslice -package=images -input=./_assets/images/koma_taiki.png -output=./assets/images/koma_taiki.go -var=Koma_taiki_png
//go:generate file2byteslice -package=images -input=./_assets/images/koma_00.png -output=./assets/images/koma_00.go -var=Koma_00_png
//go:generate file2byteslice -package=images -input=./_assets/images/koma_01.png -output=./assets/images/koma_01.go -var=Koma_01_png
//go:generate file2byteslice -package=images -input=./_assets/images/koma_02.png -output=./assets/images/koma_02.go -var=Koma_02_png
//go:generate file2byteslice -package=images -input=./_assets/images/koma_03.png -output=./assets/images/koma_03.go -var=Koma_03_png

// TODO: 獅子丸の待機画像も作りたい
//go:generate file2byteslice -package=images -input=./_assets/images/shishimaru_00.png -output=./assets/images/shishimaru_taiki.go -var=Shishimaru_taiki_png
//go:generate file2byteslice -package=images -input=./_assets/images/shishimaru_00.png -output=./assets/images/shishimaru_00.go -var=Shishimaru_00_png
//go:generate file2byteslice -package=images -input=./_assets/images/shishimaru_01.png -output=./assets/images/shishimaru_01.go -var=Shishimaru_01_png
//go:generate file2byteslice -package=images -input=./_assets/images/shishimaru_02.png -output=./assets/images/shishimaru_02.go -var=Shishimaru_02_png
//go:generate file2byteslice -package=images -input=./_assets/images/shishimaru_03.png -output=./assets/images/shishimaru_03.go -var=Shishimaru_03_png

/* package audios */
//go:generate file2byteslice -package=audios -input=./_assets/audios/hashire_kurona.mp3 -output=./assets/audios/hashire_kurona.go -var=Hashire_kurona_mp3
//go:generate file2byteslice -package=audios -input=./_assets/audios/shibugaki_no_kuroneko.mp3 -output=./assets/audios/shibugaki_no_kuroneko.go -var=Shibugaki_no_kuroneko_mp3

/* package se */
//go:generate file2byteslice -package=se -input=./_assets/se/jump.wav -output=./assets/se/jump.go -var=Jump_wav

//go:generate gofmt -s -w .
