# 黒菜んダッシュ :dash:

[![License](https://img.shields.io/github/license/kemokemo/kuronan-dash)](https://opensource.org/licenses/Apache-2.0) [![Go Version](https://img.shields.io/github/go-mod/go-version/kemokemo/kuronan-dash)](https://github.com/kemokemo/kuronan-dash/blob/main/go.mod) [![Go Report Card](https://goreportcard.com/badge/github.com/kemokemo/kuronan-dash)](https://goreportcard.com/report/github.com/kemokemo/kuronan-dash) [![LatestVersion](https://img.shields.io/github/v/release/kemokemo/kuronan-dash?color=8783f7)](https://github.com/kemokemo/kuronan-dash/releases/latest) [![OpenIssues](https://img.shields.io/github/issues-raw/kemokemo/kuronan-dash?color=fca438)](https://github.com/kemokemo/kuronan-dash/issues)
[![Test and Build](https://github.com/kemokemo/kuronan-dash/actions/workflows/test-and-build.yml/badge.svg)](https://github.com/kemokemo/kuronan-dash/actions/workflows/test-and-build.yml) ![Deploy GitHub pages](https://github.com/kemokemo/kuronan-dash/workflows/Deploy%20GitHub%20pages/badge.svg)

## 概要

月刊COMICリュウで人気連載中のコミック、「[ねこむすめ道草日記](http://www.comic-ryu.jp/_nekomusume/)」の同人ゲームです。

黒菜が跳ぶ！独楽の拳が唸る！獅子丸が走る！  
今日もみんなで駆け抜けろ！

## ゲーム紹介

タイトル画面です。`Space`キーまたはマウスの左クリックでキャラクター選択画面へと進みます。

![TitleScreen](media/title_screen.png)

キャラクターを1人選びます。選択可能なキャラクターは「黒菜」「独楽」「獅子丸」の3人です。
キーボードの矢印キーの左右でキャラクターを選択して、`Space`キーで選択します。
マウスをキャラクターの枠の内側に持っていって左クリックしても選択できます。

![SelectScreen](media/select_screen.png)

コースをダッシュで走り抜けます。`Space`キーまたはマウスの左クリックでゲームを開始します。
キーボードの矢印キーの上下または、マウスの左クリックでレーンを移動しながら進み、タイムがゼロになる前にゴールまでたどり着けたらステージクリアです。岩などの障害物に当たっている間やスタミナが`0`の間は走れなくなり、速度がゆっくりになります。

![GameScreenKurona](media/game_screen_kurona.png)

`Space`キーまたはマウスの右クリックで一時停止ができます。

### プレイ方法

[こちらのページ](https://kemokemo.github.io/kuronan-dash/)から、ブラウザ上で遊べます。

PCにダウンロードして遊んでいただける場合は、[最新版のリリースページ](https://github.com/kemokemo/kuronan-dash/releases/latest)からお使いのOSに応じた実行ファイルをダウンロードして実行してください。

## 開発者向けの情報

- [Developers Guide](documents/developers-guide.md)
- [Specification List](documents/spec-list.md)

## 作者

:cat: [kemokemo](https://github.com/kemokemo)

## ライセンス

:orange_book: [Apache License Version 2.0](https://github.com/kemokemo/kuronan-dash/blob/main/LICENSE)

ソースコードだけでなく、`assets`ディレクトリ以下の画像や音楽、効果音データなども上記のライセンスです。なお、フォントは[患者長ひっくさん](https://twitter.com/hicchicc)の[ザ・ストロングゲーマー](http://www17.plala.or.jp/xxxxxxx/00ff/)フォントを使わせていただいています。素晴らしいフォントです。

## スペシャルサンクス

[@hajimehoshi](https://github.com/hajimehoshi)さんが作っておられる素敵なGo言語の2Dゲームライブラリ[ebiten](https://github.com/hajimehoshi/ebiten)を使っています。  
この場をお借りしてお礼申し上げます。

そして、大好きな「[ねこむすめ道草日記](http://www.comic-ryu.jp/_nekomusume/)」の作者である「[いけ先生](https://twitter.com/ikenokappa)」に感謝申し上げます。
