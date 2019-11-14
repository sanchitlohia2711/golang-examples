1.WithValue() - Passing Request Scoped Values

The complete signature of the function is 
withValue(parent Context, key, val interface{}) (ctx Context)

It takes in a parent context, key, value and returns a derived context aIt returns a derived context. This derived context has key associated with the value. Here the parent context can be either context.Background() or any other context. Further any context which is derived from this context will have this value. 

ctxRoot := context.Background() - #Root context 
ctxChild := context.WithValue(ctxRoot, "a", "x") #It has acess to only one pair {"a":"x"}
ctxChildofChild := context.WithValue(ctxChild, "b", "y") #It has access to both pairs {"a":"x", "b" :"y"} as it is derived from ctxChild

Complete Working example of withValue(). In below example we are injecting a msgId for each incoming request. If you notice in below program
inejctMsgID is a net http middleware function that populates the "msgID" field in context
HelloWorld is the handler function for api "localhost:8080/welcome" which gets this msgID from context and sends it back as response headers
package main
import (
    "context"
    "net/http"
    "github.com/google/uuid"
)
func main() {
    helloWorldHandler := http.HandlerFunc(HelloWorld)
    http.Handle("/welcome", inejctMsgID(helloWorldHandler))
    http.ListenAndServe(":8080", nil)
}
//HelloWorld hellow world handler
func HelloWorld(w http.ResponseWriter, r *http.Request) {
    msgID := ""
    if m := r.Context().Value("msgId"); m != nil {
        if value, ok := m.(string); ok {
            msgID = value
        }
    }
    w.Header().Add("msgId", msgID)
    w.Write([]byte("Hello, world"))
}
func inejctMsgID(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        msgID := uuid.New().String()
        ctx := context.WithValue(r.Context(), "msgId", msgID)
        req := r.WithContext(ctx)
        next.ServeHTTP(w, req)
        
    })
}

2.WithCancel() - with cancellation signals
Below is the signature of WithCancel() function
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
context.WithCancel() function returns two things
Copy of the parentContext with the new done channel.
A cancel function which when called closes this done channel

Only the creator of this context should call the cancel function. It is highly not recommended to pass around the cancel function. Lets understand withCancel with an example.

package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx := context.Background()
    cancelCtx, cancelFunc := context.WithCancel(ctx)
    go task(cancelCtx)
    time.Sleep(time.Second * 3)
    cancelFunc()
    time.Sleep(time.Second * 1)
}

func task(ctx context.Context) {
    i := 1
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Gracefully exit")
            fmt.Println(ctx.Err())
            return
        default:
            fmt.Println(i)
            time.Sleep(time.Second * 1)
            i++
        }
    }
}

Output:
1
2
3
Gracefully exit
context canceled

In the above program
task function will gracefully exit once the cancelFunc is called. Once the cancelFunc is called , the error string is set to "context cancelled" by the context package. That is why the output of ctx.Err() is "context cancelled"

3.WithTimeout() - time based cancellation
The signature of the function is 
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
context.WitTimeout() function 
Will return copy of the parentContext with the new done channel.
Accept a timeout duration after which this done channel will be closed and context will be cancelled
A cancel function which can be called in case the context needs to be cancelled before timeout.

Lets see an example
package main
import (
    "context"
    "fmt"
    "time"
)
func main() {
    ctx := context.Background()
    cancelCtx, cancel := context.WithTimeout(ctx, time.Second*3)
    defer cancel()
    go task1(cancelCtx)
    time.Sleep(time.Second * 4)
}
func task1(ctx context.Context) {
    i := 1
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Gracefully exit")
            fmt.Println(ctx.Err())
            return
        default:
            fmt.Println(i)
            time.Sleep(time.Second * 1)
            i++
        }
    }
}

Output:
1
2
3
Gracefully exit
context deadline exceeded

In the above program
task function will gracefully exit once the timeout of 3 second is completed. The error string is set to "context deadline exceeded" by the context package. That is why the output of ctx.Err() is "context deadline exceeded"
4.WithDeadline() - time based cancellation
The signature of the function is 
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
context.WithDeadline() function 
Will return copy of the parentContext with the new done channel.
Accept a deadline after which this done channel will be closed and context will be cancelled
A cancel function which can be called in case the context needs to be cancelled before the deadline is reached.

Lets see an example
package main
import (
    "context"
    "fmt"
    "time"
)
func main() {
    ctx := context.Background()
    cancelCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
    defer cancel()
    go task(cancelCtx)
    time.Sleep(time.Second * 4)
}
func task(ctx context.Context) {
    i := 1
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Gracefully exit")
            fmt.Println(ctx.Err())
            return
        default:
            fmt.Println(i)
            time.Sleep(time.Second * 1)
            i++
        }
    }
}

Output:
1
2
3
Gracefully exit
context deadline exceeded

In the above program
task function will gracefully exit once the timeout of 3 second is completed. The error string is set to "context deadline exceeded" by the context package. That is why the output of ctx.Err() is "context deadline exceeded"


All ways of creating a error in Golang

Before learning about different ways of creating a error in golang lets first understand about the error interface. Below is the signature of the error interface

type error iterface{
    Error() string
}

So any type which defines the Error function will be of type error

Now lets see different types

1. sampleErr := errors.New("werwer")

errors.New("") returns a struct which implements the Error function. What ever you pass as an argument to the errors.New() it will returned when calling sampleErr.Error()

2. sampleErr := fmt.Errorf("error is %s", "connection issue")
creates error with formatting. Else everything is same as errors.New()

3. log error and exit

log.errorf("Unable to connect to databases")
Above function will log error first and then call os.Exit(1). So the process will exit 