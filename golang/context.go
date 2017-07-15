package main

//This is a test of the context package

import (
	"context"
	"fmt"
	"time"
)

//Context包为上下文管理提供了便利，
//一般情况下通过继承context.Background生成满足我们自己需求的context, 然后就可以用来管理我们的过程调用上下文了。
//这里当然不局限于web应用.

func main() {
	withValue := context.WithValue(context.Background(), "SID", "12345465")
	ctx, cancel := context.WithCancel(withValue)
	go ContextWithValueAndCancelTest(ctx)

	<-time.Tick(time.Second * 5)
	cancel()
	<-time.Tick(time.Second * 5)

	//这里既可以通过withTimeout的超时来取消，也可以主动调用cancel回调函数来取消
	//withTimeout, cancel := context.WithTimeout(withValue, time.Duration(time.Second*13))
	withTimeout, _ := context.WithTimeout(withValue, time.Duration(time.Second*13))
	ContextWithValueAndTimeoutTest(withTimeout)
}

func ContextWithValueAndCancelTest(ctx context.Context) {
	sid := ctx.Value("SID")
	sessionid, _ := sid.(string)
	fmt.Println(sessionid)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("With Cancel Cancelled")
			return
		}
	}
}

func ContextWithValueAndTimeoutTest(ctx context.Context) {
	sid := ctx.Value("SID")
	sessionid, _ := sid.(string)
	fmt.Println(sessionid)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("With timeout Cancelled")
			return
		}
	}
}
