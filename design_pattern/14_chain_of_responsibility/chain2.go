package chain

import "fmt"

type chainReq struct {
	money int64
}

type chainHandler interface {
	pre(req chainReq) chainReq
	process(req chainReq) bool
	after(req chainReq)
}

type chainManager struct {
	chainList []chainHandler
}

func NewChainManager() *chainManager {
	return &chainManager{chainList: []chainHandler{}}
}

func (cm *chainManager) addHandler(newHandler chainHandler) {
	cm.chainList = append(cm.chainList, newHandler)
}

func (cm *chainManager) process(req chainReq) bool {
	type funcType func(r chainReq, list []chainHandler) bool
	var f funcType
	f = func(r chainReq, list []chainHandler) bool {
		if len(list) == 0 {
			return true
		}
		req = list[0].pre(r)
		defer list[0].after(req)
		if len(list) == 1 {
			return list[0].process(req)
		} else {
			return f(req, list[1:])
		}
	}

	return f(req, cm.chainList)
}

type ProjectHandler struct{}

func NewProjectHandlerChain() chainHandler {
	return &ProjectHandler{}
}

func (*ProjectHandler) pre(req chainReq) chainReq {
	req.money = 100
	fmt.Printf("Project Handler change fee request:%d\n", req.money)
	return req
}

func (*ProjectHandler) process(req chainReq) bool {
	if req.money >= 100 {
		fmt.Printf("Project Handler permit %d fee request\n", req.money)
		return true
	}
	fmt.Printf("Project Handler don't permit %d fee request\n", req.money)
	return false
}

func (*ProjectHandler) after(req chainReq) {
	fmt.Printf("Project Handler after money:%d\n", req.money)
}

type DepHandler struct{}

func NewDepHandlerChain() chainHandler {
	return &DepHandler{}
}

func (*DepHandler) pre(req chainReq) chainReq {
	req.money = 200
	fmt.Printf("Dep Handler change fee request:%d\n", req.money)
	return req
}

func (*DepHandler) process(req chainReq) bool {
	if req.money >= 200 {
		fmt.Printf("Dep Handler permit %d fee request\n", req.money)
		return true
	}
	fmt.Printf("Dep Handler don't permit %d fee request\n", req.money)
	return false
}

func (*DepHandler) after(req chainReq) {
	fmt.Printf("Dep Handler after money:%d\n", req.money)
}

type GeneralHandler struct{}

func NewGeneralHandlerChain() chainHandler {
	return &GeneralHandler{}
}

func (*GeneralHandler) pre(req chainReq) chainReq {
	req.money = 300
	fmt.Printf("General Handler change fee request:%d\n", req.money)
	return req
}

func (*GeneralHandler) process(req chainReq) bool {
	if req.money >= 400 {
		fmt.Printf("General Handler permit %d fee request\n", req.money)
		return true
	}
	fmt.Printf("General Handler don't permit %d fee request\n", req.money)
	return false
}

func (*GeneralHandler) after(req chainReq) {
	fmt.Printf("General Handler after money:%d\n", req.money)
}
