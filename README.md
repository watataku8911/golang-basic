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