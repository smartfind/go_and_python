package main

import "net/http"

//panic的坑
func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		//	go func() {
		//		panic("error")
		//	}()
		//	w.Write([]byte("hello world"))
	})
	//panic("error") //panic自带recover，在父协程中可以捕获时，不会造成主协程挂掉
	//比如操作redis,有人觉得这段代码可以放到协程中去运行，可能有很大隐患
	http.ListenAndServe("127.0.0.1:8080", nil)
	//	panic会硬气主线程挂掉，同时会导致其他协程挂掉
	//	在父协程中无法捕获的子协程中出现的异常

}
