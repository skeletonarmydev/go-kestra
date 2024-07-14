# go-kestra

A simple golang module for [kestra.io](http://kestra.io).

Note: not yet API complete

Install:
```
go get github.com/skeletonarmydev/go-kestra
```

Usage:
```
import (
kestra "github.com/skeletonarmydev/go-kestra/kestra-oss/v1"
...
)

kestraClient, _ := kestra.NewClient("<kestrabaseurl>", nil)
flows, resp, err := kestraClient.Flow.GetAll(context.Background(), "some_namespace")

for _, element := range *flows {
  fmt.Println("Flow: " + element.ID)
}
```
Note: your `<kestrabaseurl>` should be something like: http://localhost:8080/  (with the trailing slash)
