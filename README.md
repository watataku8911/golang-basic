# golang-basic

[参照](https://qiita.com/watataku8911/items/f12163ade0d820d00bba)

##　変数

変数は以下のとうりで宣言します。
*＊ちなみにGo言語ではセミコロン(;)はいりません！！*

```
var 変数名　データ型
```

またこのような書き方もあります。

```
var 変数名　= 初期値
```

さらにこのような書き方もできます。

```
変数名　:= 初期値
```
3番目の記述方法が多く使われていて、参考書とかもこの書き方が多い。

## 定数

定数は以下のように記述します。

```
const 定数名　= 初期値
```
*定数も変数と同じく型指定できるが、通常は型を指定しません。*

## 条件分岐

- if文

```
if 条件式 {
    //処理
} else if 条件式 {
    //処理
} else {
    //処理
}
```

[例](https://github.com/watataku8911/golang-basic/blob/master/variable-if.go)

- switch文

```
switch 条件の値 {
    case 値１:
         //処理
    case 値２:
         //処理
    case 値３:
         //処理
        ・
        ・
        ・
　　 default:
　　　　　 //処理
}
```
*＊CやJavaなどのswitch文は，1つのcaseが実行されるとその次のcaseに処理が移るため，単一のcaseの実行で終わらせたい場合に，caseごとにbreakを書く必要がありました。しかしGoのswitch文では，caseが1つ実行されると次のcaseに実行が移ることなくswitch文が終了するため，breakをいちいち書く必要はありません。*

[例](https://github.com/watataku8911/golang-basic/blob/master/switch.go)

## ループ処理
*＊Goでのループはwhile文がなくfor文のみ。*

```
for 初期値; 条件式; 初期値をどうする？ {
    //処理
}
```

[例](https://github.com/watataku8911/golang-basic/blob/master/for.go)

## 配列

[参照](https://qiita.com/watataku8911/items/50de84b191f73d281be9)

*Goの配列は固定長。可変長はスライス。（後述）*

```
var 配列名 [要素数] データ型
```

またこのような書き方もできます。

```
配列名 := [要素数] データ型　{初期値1, 初期値2, 初期値3,　・・・}
```

*＊[要素数]を[...]にすると配列の長さは暗黙的に指定できる。*

## スライス

```
var スライス名 [] データ型
```

またこのような書き方もできます。

```
スライス名 := [] データ型　{初期値1, 初期値2, 初期値3,　・・・}
```

### スライスの末尾に値を追加する方法

```
スライス名 = append(スライス名, "値")
```

## マップ
値をkey-valueの対応で保存するデータ構造。属に言う連想配列。

```
var マップ名 map[キーの型] 値の型
```

またこのような書き方もできます。

```
マップ名 := map[キーの型] 値の型 {
    キー : 値
        ・
        ・
        ・
}
```

### マップからデータを削除する

```
delete(マップ名, 削除したいキー)
````

*＊Goで、mapをrange(後述)でイテレーションすると、取り出す順番は実行ごとに異なるので注意。*
理由：Goでは乱択アルゴリズムが使われているようです。

## range

配列やスライスなどに格納された値を先頭から順番に処理していく。for文と一緒に使う

```
for 添字, 値 := range 配列名（ここがスライス名などの名前が変わる） {
    //処理
}
```

[例](https://github.com/watataku8911/golang-basic/blob/master/array.go)

## 関数

[参照](https://qiita.com/watataku8911/items/79f5e49d08000838150d)

```
func 関数名(引数 データ型) 戻り値の型 {
    return 戻り値
}
```

*＊Goでは複数の戻り値が指定できる。(カンマ(,)区切り)で。その際、戻り値のデータ型もカンマ区切りで()でくくる。*

[例](https://github.com/watataku8911/golang-basic/blob/master/func.go)

## ポインタ

[参照](https://qiita.com/watataku8911/items/ef518d59a4ae31b45b63)

[例](https://github.com/watataku8911/golang-basic/blob/master/pointer.go)

## 構造体

構造体は下記のように``「struct」``を使用して定義します。
*＊Goの場合、大文字から始まる名前は、他のパッケージからアクセス可能となります。
一方、小文字で始まる名前の場合、他のパッケージからアクセス不可能です(カプセル化)。同じパッケージ内であれば小文字でもアクセス可能です。*

```
type 構造体名 struct {
   フィールド名 データ型
}
```

### 構造体にアクセス

```
var 変数名 構造体名

変数名.フィールド名
```

### メソッド

```
func (レシーバ名  レシーバの型) 関数名(引数) (戻り値) {
    //処理
}
```

[例](https://github.com/watataku8911/golang-basic/blob/master/struct.go)

## 時間の扱い方

**timeパッケージ**を使う。

[参照](https://qiita.com/watataku8911/items/7c32df8a8c6b59d7401c)

[例](https://github.com/watataku8911/golang-basic/blob/master/time.go)

