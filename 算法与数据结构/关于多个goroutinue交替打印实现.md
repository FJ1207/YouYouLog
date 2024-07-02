<h1 align=center>关于多个goroutinue交替打印的实现</h1>

# 交替打印数字和字母，各十次

```
func maxEqualFreq() {
	var (
		C    = make(chan struct{}, 1)
		C2   = make(chan struct{}, 1)
		wg   sync.WaitGroup
		N    = 10
		Char = 'A'
	)
	wg.Add(2)
	C <- struct{}{}
go func() {
		defer wg.Done()
		for i := 1; i <= N; i++ {
			<-C
			fmt.Println(i)
			C2 <- struct{}{} //交替的关键，通知另外一个
		}
	}()
go func() {
		defer wg.Done()
		for i := 1; i <= N; i++ {
			<-C2
			fmt.Println(string(Char))
			Char++
			C <- struct{}{}
		}
	}()
	wg.Wait()
}
```

# 交替打印1-100的奇数和偶数

```
func print() {
        var (
                //无缓存
                ch = make(chan struct{})
                wg sync.WaitGroup
        )
        wg.Add(2)
        go func() {
                defer wg.Done()
                for i := 1; i <= 100; i++ {
                        // 打印奇数
                        <-ch
                        if i%2 == 1 {
                                fmt.Printf("奇数goroutine: %d\n", i)
                        }
                }
        }()

        go func() {
                defer wg.Done()
                for i := 1; i <= 100; i++ {
                        // 打印偶数
                        ch <- struct{}{}
                        if i%2 == 0 {
                                fmt.Printf("偶数goroutine: %d\n", i)
                        }
                }
        }()
        wg.Wait()
}
```
## 十个goroutinue异步计算1-10000的累加和
```
func add() {
	var (
		wg   sync.WaitGroup
		res  = 0
		l    = 10000
		ch   = make(chan int, 10)
		res1 = 0
	)
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func(i int) {
			defer wg.Done()
			s := 0
			for j := (i-1)*1000 + 1; j <= i*1000; j++ {
				s = s + j
			}
			fmt.Println(s)
			ch <- s
		}(i)
	}
	wg.Wait()
	close(ch)
	for v := range ch {
		res += v
	}
	fmt.Println(res)
	for q := 1; q <= l; q++ {
		res1 = res1 + q
	}
	fmt.Println(res1)
}
```