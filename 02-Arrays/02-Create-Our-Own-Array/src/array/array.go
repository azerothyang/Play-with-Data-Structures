package array

import "strconv"

type Array struct {
	data []int
	size int
}

// 构造一个数组
func NewArray(capacity int) *Array {
	return &Array{
		data: make([]int, capacity),
		size: 0,
	}
}

// 获取数组元素个数
func (arr *Array) GetSize() int {
	return arr.size
}

// 获取数组的容量
func (arr *Array) GetCapacity() int {
	return len(arr.data)
}

// 数组末尾加入新的元素
func (arr *Array) AddLast(e int) {
	arr.Add(arr.size, e)
}

// 数组最前面加入新的元素
func (arr *Array) AddFirst(e int) {
	arr.Add(0, e)
}

// 在数组指定位置插入元素
func (arr *Array) Add(index int, ele int) {
	if arr.size == arr.GetCapacity() {
		panic("add element fail, capacity is full!")
	}
	if index >= arr.GetCapacity() {
		panic("index is larger than capacity!")
	}
	// 其他元素为新元素腾出位置
	for i := arr.size - 1; i >= index; i-- {
		arr.data[i+1] = arr.data[i]
	}

	arr.data[index] = ele
	arr.size++
}

// get data
func (arr *Array) GetData() []int {
	return arr.data
}

// 获取index索引位置
func (arr *Array) Get(index int) int {
	if index < 0 || index >= arr.size {
		panic("Get fail! Index is illegal.")
	}
	return arr.data[index]
}

// 获取index索引位置
func (arr *Array) Set(index int, e int) {
	if index < 0 || index >= arr.size {
		panic("Get fail! Index is illegal.")
	}
	arr.data[index] = e
}

// ToString is to return array
func (arr *Array) ToString() string {
	str := "Array: size = " + strconv.Itoa(arr.size) + "  , capacity = " + strconv.Itoa(arr.GetCapacity()) + "\n["
	for i := 0; i < arr.size; i++ {
		str += strconv.Itoa(arr.data[i]) + ", "
	}
	str = str[:len(str)-2]
	str += "]"
	return str
}
