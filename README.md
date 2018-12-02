## AtCoderをGolangで解いてるやつ

### 問題選び
- [AtCoder に登録したら次にやること ～ これだけ解けば十分闘える！過去問精選 10 問 ～ - Qiita](https://qiita.com/drken/items/fd4e5e3630d0f5859067)

### MEMO
1. まず入力を定義してScan
2. 次に出力を定義してTest
3. 問題を解く

### NOTE
- 文字列はスライス風にアクセスできる text[0]
    - []長さ
    - []連結
    - []大小比較
- 複数Scanにもコツがいる
    - fmt.Scanは改行と空白の次の文字をつかむ
    - 長さがわかる複数"1 2 3"入力なら[]intに格納すると便利
- forの抜け方 break