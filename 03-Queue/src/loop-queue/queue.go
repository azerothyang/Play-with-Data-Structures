package lqueue

import (
	"fmt"
	"strconv"
)

type Queue struct {
	Data    []interface{} // data array
	Front   int           // front position 队列最前面的一个元素索引
	Tail    int           // tail position 队列尾部下一个空的存储空间
	EleNums int           // 实际包含的元素个数
}

func (q *Queue) GetIndex() int {
	return q.Tail
}

func (q *Queue) Enqueue(ele interface{}) {
	//判断是否达到容积最大空间， 如果是则扩容

	if (q.Tail+1)%cap(q.Data) == q.Front {
		q.resize(cap(q.Data) * 2)
	}

	//写入队列
	q.Data[q.Tail] = ele
	q.Tail = (q.Tail + 1) % cap(q.Data)
	q.EleNums++
}

func (q *Queue) Dequeue() interface{} {
	if q.Front == q.Tail {
		//如果队列为空，返回nil
		return nil
	}
	if q.EleNums == cap(q.Data)/4 && cap(q.Data)/2 != 0 {
		//如果队列中的元素数量小于等于队列容积的1/4，则缩容到原容积的1/2并且不为0，懒加载
		q.resize(cap(q.Data) / 2)
	}
	ret := q.Data[q.Front]
	q.Data[q.Front] = nil
	q.Front = (q.Front + 1) % cap(q.Data)
	q.EleNums--
	return ret
}

func (q *Queue) IsEmpty() bool {
	return q.EleNums == 0
}

func (q *Queue) GetFront() interface{} {
	return q.Data[q.Front]
}

func (q *Queue) GetEleNums() int {
	return q.EleNums
}

func (q *Queue) resize(size int) {
	newData := make([]interface{}, size+1)
	length := len(q.Data)
	var i = 0
	for q.Tail != q.Front%length {
		newData[i] = q.Data[q.Front%length]
		i++
		q.Front++
	}
	q.Data = newData
	//重制指针位置
	q.Front = 0
	q.Tail = i
}

// ToString is to return array
func (q *Queue) ToString() string {
	str := "Queue: size = " + strconv.Itoa(q.EleNums) + "  , capacity = " + strconv.Itoa(cap(q.Data)) + "\nfront ["
	for i := 0; i < q.EleNums; i++ {
		str += fmt.Sprintf("%v", q.Data[i%cap(q.Data)]) + ", "
	}
	str = str[:len(str)-2]
	str += "] tail"
	return str
}

func New(length int) *Queue {
	return &Queue{
		Data:    make([]interface{}, length+1), // 留一个空位以区分front == tail时为empty 而front == tail+1为full
		Front:   0,
		Tail:    0,
		EleNums: 0,
	}
}
