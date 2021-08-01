package tags

import "fmt"

type Disk int

func (d Disk) String() string {
	return fmt.Sprintf("%d GB", d)
}

type Metrics struct {
	// `json:"cpu"` 是结构体属性的tag，用来标识其序列化、反序列化时对于的名称
	CPU    float32 `json:"cpu" yaml:"cpu"`
	Memory float32 `json:"memory"`
	Core   int16   `json:"core"`
	Disk   Disk    `json:"disk"`
}
