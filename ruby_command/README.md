## [学習中によく使うコマンド](https://github.com/DaisukeKarasawa/command/blob/master/ruby_command/command.rb)

### Stringクラス

#### 特定の文字の削除 #delete

引数に指定した文字と一致する全てを削除し、新しい文字列を返す。
```
str = 'RubyonRails'

str.delete('on') # RubyRails
```

#### 大文字・小文字へ置き換える #upcase/#downcase

全ての文字を対応する大文字・小文字に置き換えた文字列を返す。
```
lowerCaseLetter = 'my name is daisuke.'
capitalLetter = 'I STUDY RUBY.'

lowerCaseLetter.upcase  # MY NAME IS DAISUKE.
capitalLetter.downcase  # i study ruby.
```

#### 特定文字の文字数カウント #scan/Array#length

特定の文字列に対して引数にマッチした部分文字列の配列を返す。（String#scan）

配列の長さを返す。配列が空の時は０を返す。（Array#length）
```
'RubyonRails'.scan('R')         # ["R", "R"]
['Ruby', nil, 'Go'].length      # 3

'RubyonRails'.scan('R').length  # 2
```

#### 特定文字列が何文字目に含まれているか #index

文字列を右に向かって検索し、最初に見つかった部分文字列の左側のインデックスを返す。

見つからなければ、nil を返す。

第二引数 n をしてした場合、n から右に向かって検索する。

n が負の場合、末尾から数えた位置から検索する。
```
'abcdefabc'.index('abc')      # 0
'abcdefabc'.index('abc', 0)   # 0
'abcdefabc'.index('abc', -4)  # 6
```

#### 文字列から文字抽出 #slice

※ slice! を使用する場合は、破壊的メソッドになるので、元の文字列にも影響がある。

先頭(後ろ)から n 文字目抽出
```
string = 'Hello Ruby'

string.slice(0)   # H
string.slice(-1)  # y
```
先頭(後ろ)から n 番目で、そこからの m 文字取得
```
string = 'Hello Ruby'

string.slice(0, 2)    # He
string.slice(-4, 4)   # Ruby
```
文字の直接指定
```
string = 'Hello Ruby'

string.slice('H') # H
```
範囲演算子での指定
```
string = 'Hello Ruby'

string.slice(0..2)  # Hel
```

### Arrayクラス

#### 配列の先頭の要素を取り除く #shift

引数を指定した場合はその個数だけ取り除き、それを配列で返す。
```
webLanguages = ['Ruby', 'Go', 'Java', 'Python']

webLanguages.shift # ['Go', 'Java', 'Python']
```

#### 配列内の最大値・最小値の取得 #max/#min

最大の要素、最小の要素を返す。

引数に数字 n を指定した場合、n 要素が降順(昇順)に入った配列を返す。
```
maxMinArr = [1, 324, 87, 23]

maxMinArr.max     # 324
maxMinArr.max(2)  # [324, 87]
```

#### 範囲オブジェクトを配列にする #to_a

範囲オブジェクトから配列を作成して返す。
```
(1..5).to_a # [1, 2, 3, 4, 5]
```

#### 配列内の要素をランダムに返す #sample

配列の要素を1個ランダムで返す。

引数を指定した場合は自身の要素数を越えない範囲で n 個返す。
```
randArr = (1.10).to_a

randArr.sample    # 2
randArr.sample(3) # [6, 10, 1]
```

#### 配列内をソートする #sort

配列の内容を昇順でソートした配列を返す。（sort! は破壊的メソッドになる）

**注意点**

数字のソートは数の大きさ、文字列の場合、文字コードに基づいてソートされる。
```
arr = [3, 323, 12, 21, 0]
arr.sort                                # 0 3 12 21 323

arr2 = ['3', '323', '12', '21', '0']
arr2.sort                               # 0 12 21 3 323
# Integer に変換したうえで、<=> を使用して大小比較をする
arr2.sort {|a, b| a.to_i <=> b.to_i}    # 0 3 12 21 323
```

### Hashクラス

#### ハッシュからキー、値を取り出す #keys/values

ハッシュの全てのキー、または値を配列で返す。
```
hash = {"ruby": 1, "go": 2, "javascript": 3}

hash.keys    # ruby go javascript
hash.values  # 1 2 3
```