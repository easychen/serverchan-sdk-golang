# ServerChan SDK for Golang

这是用于调用 ServerChan 推送服务的 Go SDK，包名为 `serverchan-sdk`，仓库地址为 `github.com/easychen/serverchan-sdk-golang`。

## 安装

```bash
go get github.com/easychen/serverchan-sdk-golang
```

## 使用示例

```go
package main

import (
    "fmt"
    "github.com/easychen/serverchan-sdk-golang"
)

func main() {
    sendkey := "your-sendkey"
    title := "Test Message"
    desp := "This is a test message"
    
    resp, err := serverchan_sdk.ScSend(sendkey, title, desp, nil)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Response:", resp)
    }
}
```

## License

MIT License