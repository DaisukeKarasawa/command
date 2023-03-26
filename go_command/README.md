## [学習中によく使うコマンド]()

### 二つの数値の最大値・最小値の取得 math.Max/Min

引数で与えられた２つの数値のより大きい(小さい)方を返す。
```
math.Max(10, 20)  // 20
math.Max(10, 20)  // 10
```
**注意点**

引数と、戻り値はfloat64になる。
```
var n int = 100
math.Max(n, 200)  // cannot use n (variable of type int) as float64 value in argument to math.Max

math.Max(10, 20)  // 20 (float64)
```

### 整数値を返す math.Floor/Trunc

引数を整数値で返す。（ 引数、戻り値はfloat64 ）
```
tr := 1.5

math.Trunc(tr)  // 1
math.Floor(tr)  // 1
```
**func Floorとfunc Truncの違い**

*Floor*：引数の値以下の最大の整数値を返す。

*Trunc*：引数の整数値を返す。
```
tr := -1.5

math.Trunc(tr)  // -1
math.Floor(tr)  // -2
```

### 書式指定子を指定した出力 fmt.Printf

書式指定子に従って書式を設定し、フォーマット付き(f)で出力する。

**書式指定子**

「文字列型」

%s    ：文字列をそのまま表示する。

%x(%X)：16進数の小文字(大文字)で表示する。
```
s := "golang"

fmt.Printf("%s\n", s)             // golang
fmt.Printf("%x << %X\n", s, s)    // 676f6c616e67 << 676F6C616E67
```

「整数型」

%d：値をそのまま表示する。
```
n := 100

fmt.Printf("%d\n", n)   // 100
```

「浮動小数点型」

%f(%F)：実数で表示する。

%.2   ：小数点以下２桁に丸める。
```
n := 3.1415

fmt.Printf("%f\n", n)     // 3.141500
fmt.Printf("%.2f\n", n)   // 3.14
```

*「%v」*

さまざまな型の値を柔軟に出力することが出来る。

配列やスライス、マップなどについても有効に作用する。
```
n1 := 100
n2 := 3.1415
s := "golang"
arr1 := []int{1, 2, 3}
arr2 := map[int]string{1: "Go", 2: "Ruby"}

fmt.Printf("%v\n", n1)    // 100
fmt.Printf("%v\n", n2)    // 3.141500
fmt.Printf("%v\n", s)     // golang
fmt.Printf("%v\n", arr1)  // [1 2 3]
fmt.Printf("%v\n", arr2)  // map[1:Go 2:Ruby]
```

**注意点**

改行を含まないので、エスケープシーケンス \n をフォーマットとして指定する必要がある。

### 指定時間のスリープ time.Sleep

実行しているゴルーチンをしてした時間の間一時停止する。負またはゼロを指定すると、すぐに戻る。
```
time.Sleep(100 * Milliseconds)  // 100ミリ秒の間停止

// 時間の種類（一部）
Microseconds                    // マイクロ秒
Milliseconds                    // ミリ秒
Minutes                         // 分
```

### 文字列内の部分文字列の検索 strings

**func Index**

検索対象の文字列の中に指定した部分文字列が含まれるかを検索する。

含まれている場合、その最初の文字が開始するインデックス(int)が返る。

戻り値が -1 の場合、部分文字列が含まれていなかったことを表す。
```
str := "golang"

strings.Index(str, "go")  // 0
strings.Index(str, "Go")  // -1
```
**func LastIndex**

検索対象に指定した部分文字列が複数含まれている場合に、最後の部分文字列が開始するインデックスを返す。

１つも含まれていない場合、-1 が返る。
```
str := "go! golang"

strings.LastIndex(str, "go")  // 4
strings.LastIndex(str, "Go")  // -1
```

**func IndexAny**/**LastIndexAny**

文字列に含まれる、引数のいづれかの文字が開始するインデックスを返す。(func IndexAny)

いづれかの文字が最後に含まれるインデックスを返す。(LastIndexAny)

どの文字も含まれない場合、-1 が返る。
```
str := "go! golang"

strings.IndexAny(str, "go")         // 0
strings.IndexAny(str, "Go")         // 1
strings.LastIndexAny(str, "go")     // 9
strings.LastIndexAny(str, "Go")     // 5
strings.LastIndexAny(str, "Ruby")   // -1
```

**Contains**/**ContainsAny**

検索対象の文字列の中にしてした部分文字列が含まれるかどうかのみを確認する。(Contains)

いづれかの文字が含まれるかどうかのみ検索する。(ContainsAny)
```
str := "go! golang"

strings.Contains(str, "go")       // true
strings.Contains(str, "Go")       // false
strings.ContainsAny(str, "go")    // true
strings.ContainsAny(str, "Go")    // true
```

**HasPrefix**/**HasSuffix**

対象の文字列のが指定した部分文字列から開始されるかどうかをbool型で返す。(HasPrefix)

指定した部分文字列で終了するかどうかを検索する。(HasSuffix)
```
str := "go! golang"

strings.HasPrefix(str, "go")    // true
strings.HasPrefix(str, "Go")    // false
strings.HasSuffix(str, "go")    // false
strings.HasSuffix(str, "Go")    // false
```

### 大文字・小文字への変換 strings.ToUpper/.ToLower

文字列に含まれる文字をそれぞれ大文字・小文字に置換する。
```
str1 := "golang"
str2 := "gRPC"

strings.ToUpper(str1)   // GOLANG
strings.ToLower(str2)   // grpc
```

### 文字列から空白を取り除く strings.TrimSpace

第二引数に重複しない全てを、第三引数で置き換えた第一引数の文字列のコピーを返す。
```
str := "I study golang"

// 空白を取り除く
strings.ReplaceAll(str4, " ", "")   // Istudygolang
```

### [ゴルーチンの終了を待つ sync.Waitgroup/Add/Wait/Done](https://github.com/DaisukeKarasawa/go/tree/master/goroutine_prg/sync_wait)

各ゴルーチンの処理が開始される前に、main関数が終了しないように実行されるゴルーチンの各処理の完了を待つ

