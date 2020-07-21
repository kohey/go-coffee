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
