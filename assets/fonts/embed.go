package fonts

import _ "embed"

var (
	// この素晴らしいフォントは、患者長ひっくさんが作った「ザ・ストロングゲーマー」です。
	// ファミコンゲームのような雰囲気を目指す本ゲームに非常にマッチするフォントとして
	// 使わせていただいております。以下の公式ページからダウンロードが可能です。
	//
	//   http://www17.plala.or.jp/xxxxxxx/00ff/
	//

	//go:embed x8y12pxTheStrongGamer.ttf
	the_strong_gamer_ttf []byte
)
