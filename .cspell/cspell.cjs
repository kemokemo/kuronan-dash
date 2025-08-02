/** @type { import("@cspell/cspell-types").CSpellUserSettings } */
module.exports = {
    $schema:
        "https://raw.githubusercontent.com/streetsidesoftware/cspell/main/cspell.schema.json",
    version: "0.2",
    // dictionaryDefinitions[number].name と同じ名前を指定する
    dictionaries: ["words"],
    // カスタム辞書ファイル配列、複数指定可能
    dictionaryDefinitions: [
        // 辞書ファイル
        {
            name: "words",
            path: "./words.txt",
            addWords: true,
        },
    ],
    // cSpell 除外設定
    ignorePaths: ["./words.txt"],
};
