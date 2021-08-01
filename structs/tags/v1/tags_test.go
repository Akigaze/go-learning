package tags

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMetrics(t *testing.T) {
	metrics := Metrics{
		CPU:    0.8,
		Memory: 1.4,
		Core:   2,
		Disk:   500,
	}
	fmt.Println(metrics)
	// json是GO标准库中的一个包 Marshal 方法用于序列化 对象
	got, err := json.Marshal(metrics)
	want := "{\"cpu\":0.8,\"memory\":1.4,\"core\":2,\"disk\":500}"
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if string(got) != want {
		t.Errorf("expect %s but got %s", want, got)
	}

}
