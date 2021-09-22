package simplefactory

import (
	"fmt"
)

type fType int64

const (
	hiFType fType = iota + 1
	helloFType
)

// API is interface
type API interface {
	Say(name string) string
}

// confMap 可以再次优化为从文件/远程配置读取
func confMap() map[fType]API {
	return map[fType]API{
		fType(0):   nil,
		hiFType:    &hiAPI{},
		helloFType: &helloAPI{},
	}
}

// NewAPI return Api instance by type
func NewAPI(t fType) API {
	m := confMap()
	return m[t]
}

// hiAPI is one of API implement
type hiAPI struct{}

// Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// HelloAPI is another API implement
type helloAPI struct{}

// Say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
