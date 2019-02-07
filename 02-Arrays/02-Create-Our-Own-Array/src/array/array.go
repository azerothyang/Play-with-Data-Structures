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
		arr.resize(arr.GetCapacity() * 2)
		//panic("add element fail, capacity is full!")
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

// 删除index位置元素，返回删除的元素
func (arr *Array) Remove(index int) int {
	if index >= arr.size || index < 0 {
		panic("index is illegal")
	}
	removeEle := arr.data[index]
	//当前元素后面的元素集体左移一位
	for i := index + 1; i < arr.size; i++ {
		arr.data[i-1] = arr.data[i]
	}
	arr.size--
	if arr.size == arr.GetCapacity()/2 {
		// 减少一半容积
		arr.resize(arr.size)
	}
	return removeEle
}

func (arr *Array) RemoveFirst() int {
	return arr.Remove(0)
}

func (arr *Array) RemoveLast() int {
	return arr.Remove(arr.size - 1)
}

// remove exist element
func (arr *Array) RemoveElement(e int) {
	index := arr.Find(e)
	if index != -1 {
		arr.Remove(index)
	}
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

// 查找数组中是否有元素e
func (arr *Array) Contains(e int) bool {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return true
		}
	}
	return false
}

// 查找元素e对应的索引，如果没有返回-1
func (arr *Array) Find(e int) int {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return i
		}
	}
	return -1
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

// 翻倍容量
func (arr *Array) resize(newCapacity int) {
	newData := make([]int, newCapacity)
	// copy arr.data to new Data
	for i := 0; i < arr.size; i++ {
		newData[i] = arr.data[i]
	}
	arr.data = newData
}
