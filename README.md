# 黒菜んダッシュ :dash:

[![License](https://img.shields.io/github/license/kemokemo/kuronan-dash)](https://opensource.org/licenses/Apache-2.0) [![Build Status](https://app.travis-ci.com/kemokemo/kuronan-dash.svg?branch=main)](https://app.travis-ci.com/kemokemo/kuronan-dash) [![Go Version](https://img.shields.io/github/go-mod/go-version/kemokemo/kuronan-dash)](https://github.com/kemokemo/kuronan-dash/blob/main/go.mod) [![Go Report Card](https://goreportcard.com/badge/github.com/kemokemo/kuronan-dash)](https://goreportcard.com/report/github.com/kemokemo/kuronan-dash) [![LatestVersion](https://img.shields.io/github/v/release/kemokemo/kuronan-dash?color=8783f7)](https://github.com/kemokemo/kuronan-dash/releases/latest) [![OpenIssues](https://img.shields.io/github/issues-raw/kemokemo/kuronan-dash?color=fca438)](https://github.com/kemokemo/kuronan-dash/issues) ![Deploy GitHub pages](https://github.com/kemokemo/kuronan-dash/workflows/Deploy%20GitHub%20pages/badge.svg)

## 概要

月刊COMICリュウで人気連載中のコミック、「[ねこむすめ道草日記](http://www.comic-ryu.jp/_nekomusume/)」の同人ゲームです。

黒菜が跳ぶ！独楽の拳が唸る！獅子丸が走る！  
今日もみんなで駆け抜けろ！

## ゲーム紹介

こんなゲームを作っています。:blush:

キャラクターを1人選んで遊びます。選択可能なキャラクターは「黒菜」「独楽」「獅子丸」の3人です。
![SelectScreen](media/select_screen.png)

コースをダッシュで走り抜けます。岩などの障害物に当たっている間やスタミナが0の間は走れなくなって歩いてしまいます。
![GameScreenKurona](media/game_screen_kurona.png)

### プレイ方法

[最新版のリリースページ](https://github.com/kemokemo/kuronan-dash/releases/latest)からお使いのOSに応じた実行ファイルをダウンロードして実行してください。[こちらのページ](https://kemokemo.github.io/kuronan-dash/)で、ブラウザ上で遊ぶこともできます。

### 操作方法

現在は、キーボードでの操作のみ可能です。

- タイトル画面
  - `Space`キーでキャラクター選択画面へと進みます。
- キャラクター選択画面
  - キーボードの矢印キーの左右でキャラクターを選択します。
  - `Space`キーで選択します。
- ゲーム画面
  - `Space`キーでゲームを開始します。
  - キーボードの矢印キーの上下でレーンを移動しながら進みます。
  - タイムがゼロになる前に、ゴールまでたどり着けたらステージクリアです。

## 開発者向けの情報

### ビルド方法

`Go Modules`の仕組みを使っており、Go Ver. 1.16以上を使います。`kuronan-dash`バイナリを実行するとゲーム画面が開きます。

```sh
go build
```

`WebAssembly`形式へとビルドして、ブラウザで遊ぶこともできます。以下のようにビルドます。

```sh
GOOS=js GOARCH=wasm go build -o public/kuronan-dash.wasm
```

あとは、`public`フォルダを小さなWebサービスツールでserveすれば、ブラウザでゲームをプレイできます。[miniweb](https://github.com/kemokemo/miniweb)を使うと便利です。

```sh
miniweb -p 9000 public
```

上記を実行した場合、ブラウザで `http://localhost:9000/` を開きましょう。

### ゲームの仕様と実装状況のメモ

- [x] 制限時間内にゴールできないとゲームオーバーです。
- [x] 道は上中下の3レーンあり、ぴょんぴょん飛び移りながら走ります。
- [x] 雲などのフィールドを彩るパーツや障害物は、毎回ランダムな位置と速度で生成されます。
- [x] 道には、岩などの障害物があります。
  - [x] 障害物に当たるとキャラクターは動きが遅くなります。
  - [ ] キャラクターは障害物に攻撃できます。
  - [ ] キャラクターは攻撃が障害物に当たると、スタミナを消費します。
  - [ ] 障害物はそれぞれ耐久値を持っており、攻撃による累積ダメージが上回ると壊れて消えます。
- [x] キャラクターごとに上昇と下降の速度が異なります。
- [ ] キャラクターごとに加速度が異なります。
- [x] キャラクターはスタミナゲージをもっており、スタミナがなくなると歩くようになります。
  - [x] キャラクターごとにスタミナの総量と減少レートが異なります。
- [x] 道にはいろんな食べ物が落ちています。
  - [x] 食べ物を食べるとキャラクターはスタミナが回復します。
  - [ ] キャラクターごとに食べ物の好みがあって、スタミナの回復量は好みに左右されます。
- [ ] キャラクターはテンションゲージをもっています。
  - [ ] テンションゲージは走っても少しずつ増えますし、障害物を砕くとグンと増えます。
  - [ ] テンションMAXで、キャラクター固有のスキルが使えるようになります。:w
  - [ ] 

## 作者

:cat: [kemokemo](https://github.com/kemokemo)

## ライセンス

:orange_book: [Apache License Version 2.0](https://github.com/kemokemo/kuronan-dash/blob/main/LICENSE)

ソースコードだけでなく、`assets`ディレクトリ以下の画像や音楽、効果音データなども上記のライセンスです。なお、フォントは[患者長ひっくさん](https://twitter.com/hicchicc)の[ザ・ストロングゲーマー](http://www17.plala.or.jp/xxxxxxx/00ff/)フォントを使わせていただいています。素晴らしいフォントです。

## スペシャルサンクス

[@hajimehoshi](https://github.com/hajimehoshi)さんが作っておられる素敵なGo言語の2Dゲームライブラリ[ebiten](https://github.com/hajimehoshi/ebiten)を使っています。  
この場をお借りしてお礼申し上げます。

そして、大好きな「[ねこむすめ道草日記](http://www.comic-ryu.jp/_nekomusume/)」の作者である「[いけ先生](https://twitter.com/ikenokappa)」に感謝申し上げます。
