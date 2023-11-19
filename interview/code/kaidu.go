package code


/**

====

num_bin: 使用的内存大小应为O(num_bin)级别
window：移动平均的窗口长度

timestamp是时间戳，浮点数，数据间隔不固定，可能一秒来几十个数据，也可能几小时来一个数据


Interface
=========

（以python为例）
class MovingAverage:
    def __init__(self, num_bin: int, window: float):
        pass

    def Get(self) -> float:
        "返回当前计算的平均值"
        raise NotImplementedError
        
    def Update(self, timestamp: float, value: float):
        "新数据到达，更新状态"
        raise NotImplementedError

每次数据到达时，会依次调用Update和Get方法
**/

type MovingAverage struct{}

func Init(numBin int, window float64) *MovingAverage {
	return &MovingAverage{}
}

