# 黒菜んダッシュ Developers Guide

## 前提条件

- golang Ver. 1.9 以降

## 実行方法

リビジョン情報を渡しながら起動してみましょう。

```sh
go run -ldflags="-X 'main.Revision=$(git rev-parse --short HEAD)'" *.go
```

## ビルド方法

### 通常ビルド

タイトル画面にリビジョン情報を含むバージョン情報を表示していますので、以下のようにビルドします。

```sh
go build -ldflags="-X 'main.Revision=$(git rev-parse --short HEAD)'"
```

### WebAssembly

`WebAssembly`形式へとビルドして、ブラウザで遊ぶこともできます。以下のようにビルドます。

```sh
GOOS=js GOARCH=wasm go build -ldflags="-X 'main.Revision=$(git rev-parse --short HEAD)'" -o public/kuronan-dash.wasm
```

あとは、`public`フォルダをブラウザで閲覧可能にすれば、ブラウザでゲームをプレイできます。Dockerコンテナの`nginx`でも良いですし、拙作の[miniweb](https://github.com/kemokemo/miniweb)などの小さなWebサービスでも良いです。

```sh
miniweb -p 9000 public
```

上記を実行した場合、ブラウザで `http://localhost:9000/` を開きましょう。

## 設計について

非常に低スペックな環境でも動作することを想定しているため、極力ローカル変数を作らない設計にしています。特に、`ebiten.Game`の`Update`関数で呼び出される処理はゲームプレイ中に非常に多くの回数呼び出される処理のため、実装のしやすさを犠牲にしてもポインタで保持して管理しながら使うという方針にしています。

また、音声再生の処理などキャラクターのステータスに応じた処理では、`goroutine`を使ってもっと見通しの良いコードにしたかったのですが、2022/11/26現在`goroutine`を使うと`WebAssembly`版が応答なしになってしまい動作しないため断念しました。
