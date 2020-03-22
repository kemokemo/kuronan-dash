# assets

## アセットの追加方法とプログラムからの使い方

画像などのアセットファイルをGo言語のプログラムから直接使うため[file2byteslice](https://github.com/hajimehoshi/file2byteslice)ツールを使ってバイナリデータに変換して使います。公式ページを見ながらインストールしておきます。`go generate`の仕組みを使って実行します。

1. アセットを `../_assets`フォルダに追加する
1. `file2byteslice`ツールを使った変換処理を`generate.go`ファイルに書く
1. この`README.md`と同じパスをbashなどのシェルで開き、`go generate`を実行する
1. `generate.go`で指定した変数名を、他のプログラムコードから参照して使う
