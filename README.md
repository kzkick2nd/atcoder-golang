## AtCoderをGoで解いてるやつ

### 問題選び
- [AtCoder に登録したら次にやること ～ これだけ解けば十分闘える！過去問精選 10 問 ～ - Qiita](https://qiita.com/drken/items/fd4e5e3630d0f5859067)

### MEMO
1. まず入力を定義してScan
2. 次に出力を定義してTest
3. 問題を解く

### NOTE
- 文字列はスライス風にアクセスできる text[0]
    - 長さ len()
    - 連結 a + b
    - 大小比較 a > b
    - 文字列の番号の話しhttps://qiita.com/seihmd/items/4a878e7fa340d7963fee
    - 075B string("text"[0])
    - 075B 文字列追加+=
    - 075B 文字列の指定index置換★
- スライスの初期化 []int{1,2,3}
- 複数Scanにもコツがいる
    - fmt.Scanは改行と空白の次の文字を掴む
    - 長さが事前にわかる複数"1 2 3"入力は[]intに格納すると便利
- forの抜け方 break
- 各桁の計算は%10した余りのループ
- 数字の桁数はsliceにしてlen()
- []n桁のパターンは再帰
    - 再起関数はカウンタをポインタ渡しすると考え方が楽
- どこで切るといいかわからないときにはDPが使える
- abs intは結構使う
- 複数の数字のセットをテストで比較したりまとめて扱うのどうする？
- スライスの比較はreflect.DeepEqual
- 文字列リピートはstrings.Repeat
- intスライスの最大最小はsortと位置指定
- テストを書いておくとAtCoder開かないで実行できる
- 075B 数字を文字列 strconv.Itoa(int)
- 096C スライスの全探査にはrange使うと便利
- 096C boolの0値はfalse
- 096C グリッド問題は位置ではなく、配列の添字で考える?グリッド位置むずい
- 096C []byteには""文字列ではなく''文字列が使える
- 096C []pair構造体つくると探査位置
- TODO 096C なにを間違えていたか確認
- 079C fmt.Sprintf()
- 079C rune数字のint32化