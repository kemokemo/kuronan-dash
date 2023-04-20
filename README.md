# 黒菜んダッシュ :dash:

[![License](https://img.shields.io/github/license/kemokemo/kuronan-dash)](https://opensource.org/licenses/Apache-2.0) [![Go Version](https://img.shields.io/github/go-mod/go-version/kemokemo/kuronan-dash)](https://github.com/kemokemo/kuronan-dash/blob/main/go.mod) [![Go Report Card](https://goreportcard.com/badge/github.com/kemokemo/kuronan-dash)](https://goreportcard.com/report/github.com/kemokemo/kuronan-dash) [![LatestVersion](https://img.shields.io/github/v/release/kemokemo/kuronan-dash?color=8783f7)](https://github.com/kemokemo/kuronan-dash/releases/latest) [![OpenIssues](https://img.shields.io/github/issues-raw/kemokemo/kuronan-dash?color=fca438)](https://github.com/kemokemo/kuronan-dash/issues)

## 概要

月刊COMICリュウで人気連載中のコミック、「[ねこむすめ道草日記](http://www.comic-ryu.jp/_nekomusume/)」の同人ゲームです。

黒菜が跳ぶ！独楽の拳が唸る！獅子丸が走る！  
今日もみんなで駆け抜けろ！

## ゲーム紹介

タイトル画面です。キーボードの`Space`キーを押すか、`スタート`ボタンをマウスやタッチで押します。

![TitleScreen](media/title_screen.png)

キャラクター選択画面です。「黒菜」「独楽」「獅子丸」の3人から1人選びましょう。
キーボードの矢印キーの左右を押すか、キャラクターの枠をマウスやタッチで押してキャラクターを選択します。
キーボードの`Space`キーを押すか、`しゅっぱつ!`ボタンをマウスやタッチで押して進みます。

![SelectScreen](media/select_screen.png)

コースをダッシュで走り抜けます。`Space`キーまたはマウスやタッチで開始ボタンを押してゲームを開始します。

キーボードの矢印キーの上下、マウスで別の高さのエリアを左クリック、`うえへ`や`したへ`のボタンをタッチしてレーンを上下に移動しながら進み、タイムがゼロになる前にゴールまでたどり着けたらステージクリアです。岩などの障害物に当たっている間やスタミナが`0`の間は走れなくなり、速度がゆっくりになります。

![GameScreenKurona](media/game_screen_kurona.png)

`Space`キーまたは右上のポーズボタンをマウスやタッチで押すと一時停止ができます。

岩などの障害物は`A`キーを押すかマウスの右クリック、`こうげき`ボタンをタッチするなどして攻撃を繰り出すことで壊せます。テンションゲージがMaxになったら、`S`キーを押すかマウスダブルクリック、`スキル`ボタンをタッチしてスキルを発動して速度などをあげて走ることができます。

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
