package serverchan_sdk

import (
	"fmt"
	"testing"
)

// TestScSendRealRequest 发送真实请求到 ServerChan API
func TestScSendRealRequest(t *testing.T) {
	// 请在这里填写你自己的 sendkey
	sendkey := "" // 需要替换为真实的 sendkey
	title := "测试消息"
	desp := "这是一条来自 Go 的测试消息。"

	// 调用 ScSend 函数发送请求
	resp, err := ScSend(sendkey, title, desp, nil)
	if err != nil {
		t.Fatalf("ScSend failed: %v", err)
	}

	// 打印响应结果，便于检查
	fmt.Printf("Response: %+v\n", resp)

	// 验证返回结果
	if resp.Code != 0 {
		t.Errorf("Expected success response, but got: %+v", resp)
	}
}
