[![Build Status](https://travis-ci.org/kemokemo/kuronan-dash.svg?branch=master)](https://travis-ci.org/kemokemo/kuronan-dash) [![Coverage Status](https://coveralls.io/repos/github/kemokemo/kuronan-dash/badge.svg?branch=master)](https://coveralls.io/github/kemokemo/kuronan-dash?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/kemokemo/kuronan-dash)](https://goreportcard.com/report/github.com/kemokemo/kuronan-dash) [![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

# 黒菜んダッシュ

月刊COMICリュウで人気連載中のコミック、「[ねこむすめ道草日記](http://www.comic-ryu.jp/_nekomusume/)」の同人ゲームです。

黒菜が跳ぶ！独楽の拳が唸る！獅子丸は優しくて強い子！  
今日もみんなで駆け抜けろ！

## どんなゲーム？

こんな特徴をもったゲームを作ろうとしています。

* それぞれ個性をもったキャラを選択できる
  * もちろん登場キャラクターは「ねこむすめ道草日記」の子たち
* 制限時間内にコースをダッシュで走り抜ける
* 道は上中下の3レーンあり、ぴょんぴょん飛び移りながら走る
* 道には山あり谷あり岩もあり。岩は砕いてもよし、避けてもよし、スタミナと相談して進む
* キャラクターはスタミナゲージをもっており、スタミナがなくなるとダッシュできなくなる
  * 道中、いろんなアイテムを食べてスタミナを回復させる
* キャラクターはテンションゲージをもっており、テンションMAXでスキルを発動できるようになる
  * キャラの個性がキラリと光るスキルで一気に走り抜けろ！

## つくってる人

[kemokemo](https://github.com/kemokemo)

## らいせんす

ソースコードもassetsディレクトリ以下に入っているキャラクターのドット絵や音楽なども、  
ぜ～んぶ[Apache License Version 2.0](https://github.com/kemokemo/kuronan-dash/blob/master/LICENSE)です。

## すぺしゃるさんくす

[hajimehoshi](https://github.com/hajimehoshi)さんが作っておられる素敵なGo言語の2Dゲームライブラリ[ebiten](https://github.com/hajimehoshi/ebiten)を使っています。  
この場をお借りしてお礼申し上げます。

そして大好きな「[ねこむすめ道草日記](http://www.comic-ryu.jp/_nekomusume/)」の作者である[いけ先生](https://twitter.com/ikenokappa)に感謝申し上げます。

## びるど方法

依存パッケージ管理のために `govendor` ツールを使っています。  
以下のようにしてビルドします。

```sh
$ go get -d github.com/kemokemo/kuronan-dash
$ go get -u github.com/kardianos/govendor
$ cd $GOPATH/src/github.com/kemokemo/kuronan-dash
$ govendor sync
$ go build
```
