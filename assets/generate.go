package assets

/* package images */
// Background
//go:generate file2byteslice -package=images -input=../_assets/images/title_bg.png -output=./images/title_bg.go -var=title_bg_png
//go:generate file2byteslice -package=images -input=../_assets/images/select_background.png -output=./images/select_bg.go -var=select_bg_png
//go:generate file2byteslice -package=images -input=../_assets/images/sky_background.png -output=./images/sky_bg.go -var=sky_bg_png

// Characters
//  - Kurona
//go:generate file2byteslice -package=images -input=../_assets/images/kurona_taiki.png -output=./images/kurona_taiki.go -var=kurona_taiki_png
//go:generate file2byteslice -package=images -input=../_assets/images/kurona_00.png -output=./images/kurona_00.go -var=kurona_00_png
//go:generate file2byteslice -package=images -input=../_assets/images/kurona_01.png -output=./images/kurona_01.go -var=kurona_01_png
//go:generate file2byteslice -package=images -input=../_assets/images/kurona_02.png -output=./images/kurona_02.go -var=kurona_02_png
//go:generate file2byteslice -package=images -input=../_assets/images/kurona_03.png -output=./images/kurona_03.go -var=kurona_03_png
//go:generate file2byteslice -package=images -input=../_assets/images/kurona_04.png -output=./images/kurona_04.go -var=kurona_04_png
//go:generate file2byteslice -package=images -input=../_assets/images/kurona_05.png -output=./images/kurona_05.go -var=kurona_05_png

//  - Koma
//go:generate file2byteslice -package=images -input=../_assets/images/koma_taiki.png -output=./images/koma_taiki.go -var=koma_taiki_png
//go:generate file2byteslice -package=images -input=../_assets/images/koma_00.png -output=./images/koma_00.go -var=koma_00_png
//go:generate file2byteslice -package=images -input=../_assets/images/koma_01.png -output=./images/koma_01.go -var=koma_01_png
//go:generate file2byteslice -package=images -input=../_assets/images/koma_02.png -output=./images/koma_02.go -var=koma_02_png
//go:generate file2byteslice -package=images -input=../_assets/images/koma_03.png -output=./images/koma_03.go -var=koma_03_png
//go:generate file2byteslice -package=images -input=../_assets/images/koma_04.png -output=./images/koma_04.go -var=koma_04_png
//go:generate file2byteslice -package=images -input=../_assets/images/koma_05.png -output=./images/koma_05.go -var=koma_05_png

//  - Shishimaru
//go:generate file2byteslice -package=images -input=../_assets/images/shishimaru_taiki.png -output=./images/shishimaru_taiki.go -var=shishimaru_taiki_png
//go:generate file2byteslice -package=images -input=../_assets/images/shishimaru_00.png -output=./images/shishimaru_00.go -var=shishimaru_00_png
//go:generate file2byteslice -package=images -input=../_assets/images/shishimaru_01.png -output=./images/shishimaru_01.go -var=shishimaru_01_png
//go:generate file2byteslice -package=images -input=../_assets/images/shishimaru_02.png -output=./images/shishimaru_02.go -var=shishimaru_02_png
//go:generate file2byteslice -package=images -input=../_assets/images/shishimaru_03.png -output=./images/shishimaru_03.go -var=shishimaru_03_png
//go:generate file2byteslice -package=images -input=../_assets/images/shishimaru_04.png -output=./images/shishimaru_04.go -var=shishimaru_04_png
//go:generate file2byteslice -package=images -input=../_assets/images/shishimaru_05.png -output=./images/shishimaru_05.go -var=shishimaru_05_png

// Field Parts
//  - Prairie
//go:generate file2byteslice -package=images -input=../_assets/images/tilePrairie.png -output=./images/tilePrairie.go -var=tilePrairie_png
//go:generate file2byteslice -package=images -input=../_assets/images/grass1.png -output=./images/grass1.go -var=grass1_png
//go:generate file2byteslice -package=images -input=../_assets/images/grass2.png -output=./images/grass2.go -var=grass2_png
//go:generate file2byteslice -package=images -input=../_assets/images/grass3.png -output=./images/grass3.go -var=grass3_png
//go:generate file2byteslice -package=images -input=../_assets/images/mountain_near.png -output=./images/mountain_near.go -var=mountainNear_png
//go:generate file2byteslice -package=images -input=../_assets/images/mountain_far.png -output=./images/mountain_far.go -var=mountainFar_png
//go:generate file2byteslice -package=images -input=../_assets/images/cloud_near.png -output=./images/cloud_near.go -var=cloud_near_png
//go:generate file2byteslice -package=images -input=../_assets/images/cloud_far.png -output=./images/cloud_far.go -var=cloud_far_png

/* package music */
//go:generate file2byteslice -package=music -input=../_assets/music/shibugaki_no_kuroneko.mp3 -output=./music/shibugaki_no_kuroneko.go -var=shibugaki_no_kuroneko_mp3
//go:generate file2byteslice -package=music -input=../_assets/music/hashire_kurona.mp3 -output=./music/hashire_kurona.go -var=hashire_kurona_mp3

/* package se */
//go:generate file2byteslice -package=se -input=../_assets/se/jump.wav -output=./se/jump.go -var=jump_wav

/* package fonts */
//go:generate file2byteslice -package=fonts -input=../_assets/fonts/x8y12pxTheStrongGamer.ttf -output=./fonts/the_strong_gamer.go -var=the_strong_gamer_ttf

//go:generate gofmt -s -w .
