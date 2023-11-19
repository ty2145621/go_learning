package code

import (
	"container/list"
	"math"
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
	buyLists    list.List            // 买双向链表,每个节点是一个同price链表
	saleLists   list.List            // 卖双向链表,每个节点是一个同price链表
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
	var target list.List // 要挂上去的list
	var query list.List  // 要查询成交订单的链接
	if order.quoteType == Bid {
		target = b.buyLists
		query = b.saleLists
	} else {
		target = b.saleLists
		query = b.buyLists
	}

	for query.Len() > 0 {
		v := query.Back().Value.(orderData)
		if v.price < order.price { // 不能买卖，需要挂在链表的最末端
			
		}

	}
	return math.MaxInt

}

func appendLast()
