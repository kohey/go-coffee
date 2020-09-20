# Go をやる上で知っておいた方がいい知識です

## パフォーマンス測定
- trace package を使う

### 各関数に関する概要
- 計測したい関数の中に、`defer trace.StartRegion().End()` を追記する
- それを扱う部分(main)で、開始: `trace.Start()`、終了: `defer trace.End()`

### 全体(main)のパフォーマンス測定
- Region があるなら、Task もある
- _main() を作り、開始: `trace.NewTask()`、終了: `task.End()`
- Log も取りたいので、特にエラーでは `trace.Log()`で吐き出し

-------------------------

## エラーと終了
- main(): この中のエラーは、関数を実行した「結果」のエラーなので、`log.Fprintln(os.Stderr, err)` などを呼んで、標準出力に書き出してからExit(1)
- 関数: この中のエラーは、関数実行中で「想定された」エラーなので、return などでハンドリングして抜ける(プロセスを終了する必要は特にない)。返り値の err は main で扱ってくれる

## エラーと goroutine
- ある goroutine がエラーしたとき、他の goroutine もキャンセルしたいことがある。
- context package でキャンセルできる

## 複数のエラー
- 準標準 package の errgroup を使う

ex.
```
var eg errgroup.Group()

// Go の中は goroutine で回る
eg.Go(
  func()error{
    a ,e := hoge1()
    if e != nil {
      return e
    }
    b, e := fuga1()
    if e != nil {
      return e
    }
    return nil
  })
  // }) としないと、lint エラーする

eg.Go(
  func()error{
    a ,e := hoge1()
    if e != nil {
      return e
    }
    b, e := fuga1()
    if e != nil {
      return e
    }
    return nil
  })

// ここで eg に入ったエラーを処理
if err := eg.Wait(); err != nil {
  // もし main や _main なら
  log.Fprintln(os.Stderr, err)
}
....
```
-------------------------

### WithCancel(parent Context) (ctx Context,cancel CancelFunc)
- 親の context を渡すと、子の context を作ってくれる
- cancel で 子の context をキャンセルできる

--------------------------

### goroutine の待ち合わせ
- sync.WaitGroup を使う
- goroutine の
- 前 -> Add
- defer -> Done
- 待ち合わせ -> Wait

------------------------
### データの競合を防ぐ
- goroutine や 排他的処理
- sync.Mutex を使う
- `mu.Lock()` と `defer mu.Unlock()` で OK

### sync.Mutex 
- ロックをとって、整合性を保つ方法

ex. <- 不整合
```
func main() {
    c := 0
    for i := 0; i < 1000; i++ {
        go func() {
            c++
        }()
    }
    time.Sleep(time.Second)
    fmt.Println(c)
}
```

ex. <- 整合
```
func main() {
    var mu sync.Mutex
    c := 0
    for i := 0; i < 1000; i++ {
        go func() {
            mu.Lock()         // 排他ロック取得
            defer mu.Unlock() // 関数終了時に排他ロック解除
            c++
        }()
    }
    time.Sleep(time.Second)
    fmt.Println(c)
}
```