package lqueue

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
		q.resize(len(q.Data) * 2)
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
	ret := q.Data[q.Front]
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
	newData := make([]interface{}, size)
	length := len(q.Data)
	var i = 0
	for q.Tail%length != q.Front%length {
		newData[i] = q.Data[q.Front%length]
		i++
		q.Front++
	}
	q.Data = newData
	//重制指针位置
	q.Front = 0
	q.Tail = i
}

func New(length int) *Queue {
	return &Queue{
		Data:    make([]interface{}, length+1), // 留一个空位以区分front == tail时为empty 而front == tail+1为full
		Front:   0,
		Tail:    0,
		EleNums: 0,
	}
}
