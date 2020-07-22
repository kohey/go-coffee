## defer 
後に入れたものが、先に処理される
`Last In First Out`

## Task と Region と Log
前提: main() のパフォーマンスが知りたい
1. _main()に処理を移して、main() の中で trace.Start(*file) → これで、ざっくりしたフローがわかる
2. 関数中のそれぞれの処理に対して、どれくらいのパフォーマンスが出ているのか、気になってくるはず。
3. そこで、_main() 中で、task, region を定義する
```
*main()
ctx, task := context.Newtask(context.Background(), "hoge")
defer task.End()

*_main()
defer trace.StartRegion(ctx, "リージョン").End()
or 
region := trace.StartRegion(ctx, "リージョン")
defer region.End()
```
のような感じ

## trace で task1 の プロファイルが表示されなかった
パッチを当てる
https://github.com/kumakichi/patch-go-tool-trace

## lock について
defer で unlock する場合もあるが、for や 再帰呼び出しでデッドロックを引き起こす可能性もあるので、注意

## goruitine の待ち合わせ
- Wait: 複数の goroutine を待ち合わせ
- Add : 追加した分だけ、Done メソッド が呼ばれるまで、処理をブロックする

## sync.Mutex, sync.WaitGroup
