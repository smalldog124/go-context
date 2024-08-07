package context

import (
	"context"
	"testing"
	"time"
)

func TestDeadline(t *testing.T) {
	stopTime := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), stopTime)
	//Inter- nally, context.WithTimeout creates a goroutine that will be retained in memory for 5 seconds or until cancel is called.
	defer cancel()

	ProcessSomething(ctx)

	t.Errorf("Run TestDeadline")
}

func TestCancellationSignals(t *testing.T) {
	stopTime := 3 * time.Second
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(stopTime)
		cancel()
	}()

	ProcessSomething(ctx)

	t.Errorf("Run TestDeadline")
}

type myKey string
type userKey string

func TestValue(t *testing.T) {
	ctx := context.Background()
	var key myKey = "name"
	var name = "lek"
	var user userKey = "name"
	var userData = struct {
		Name string
		Age  int
	}{Name: "lek", Age: 20}
	ctx = context.WithValue(ctx, key, name)
	ctx = context.WithValue(ctx, user, userData)

	GetValueFromContext(ctx, key)
	GetValueFromContext(ctx, user)

	t.Errorf("Run TestValue")
}
