package code

import (
	"container/list"
	"time"
)

type quoteType int

const (
	Bid quoteType = 1 // 买
	Ask quoteType = 2 // 卖
)

type orderType int

const (
	Limit  orderType = 11 // 限价单
	Market orderType = 12 // 市场单
)

type orderData struct {
	idx       int
	time      time.Time
	price     float64
	volume    int
	quoteType quoteType
	orderType orderType
}

// BaseBroker N档记事簿
type BaseBroker struct {
	buyLists    list.List            // 买双向链表
	saleLists   list.List            // 卖双向链表
	marketLists list.List            // 市场单，我自己理解市场单是一个FIFO队列，不知道对不对
	m           map[int]list.Element // key=>idx 快速定位单据所在的位置，方便撤回订单
}

// orderBook 展示最近的N档订单
func (b *BaseBroker) orderBook(level int) {

}

// transact 来了一个订单，处理合并
func (b *BaseBroker) transact(order orderData) {
	if order.orderType == Market {
		b.transactMarket(order)
	} else if order.orderType == Limit {
		b.transactLimit(order)
	}
	return
}

// transactMarket 市场单
func (b *BaseBroker) transactMarket(order orderData) {

}

// transactLimit 限价单
func (b *BaseBroker) transactLimit(order orderData) {

}
