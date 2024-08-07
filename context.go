package context

import (
	"context"
	"log"
	"time"
)

func ProcessSomething(ctx context.Context) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			return
		default:
			i++
		}
		log.Println("i ::", i)
		time.Sleep(1 * time.Second)
	}
}

func GetValueFromContext(ctx context.Context, key any) {
	v := ctx.Value(key)
	log.Println("value: ", v)
}
