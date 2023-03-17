# 学習中によく使うコマンド

# 特定削除 delete()
str = 'RubyonRails'
deleteSpace = 'I study Ruby.'

str.delete('on') # RubyRails
deleteSpace.delete(' ')  # IstudyRuby :空白削除

# 先頭削除 shift()
webLanguages = ['Ruby', 'Go', 'Java', 'Python']

webLanguages.shift # ['Go', 'Java', 'Python']

# 小大文字変換 downcase/upcase
lowerCaseLetter = 'my name is daisuke.'
capitalLetter = 'I STUDY JAVASCRIPT.'

lowerCaseLetter.downcase  # MY NAME IS DAISUKE.
capitalLetter.upcase       # i study javascript.

# 最大値,最小値 max/min
maxMinArr = [1, 324, 87, 23]

maxMinArr.max  # 324
maxMinArr.min  # 1

# 特定の文字カウント scan/length
word = 'hi! hideyuki. hidemi.'

word.scan('hi').length # 3

# 範囲オブジェクトを配列にする to_a
(1..5).to_a # [1, 2, 3, 4, 5]

# 配列の要素をランダムに返す sample
randArr = (1.10).to_a

randArr.sample  # 2
randArr.sample  # 6

# 特定文字列が何文字目に含まれているか
'abcdef'.index('abc')  # 0

# 文字列から文字抽出 slice
string = 'Hello Ruby'
# 先頭(後ろ)から n 文字目抽出
string.slice(0) # H
string.slice(-1) # y
# 先頭(後ろ)から n 番目で、そこからの m 文字取得
string.slice(0, 2)    # He
string.slice(-4, 4)   # Ruby
# 文字の直接指定
string.slice('H')     # H
# 範囲演算子での指定
string.slice(0..2)    # Hel