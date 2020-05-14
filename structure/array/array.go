package array

import (
	"fmt"
)

// Array 数组
type Array struct {
	data   []interface{}
	length int
}

// NewArray 初始化数组
func NewArray(capacity int) *Array {
	if capacity > 0 {
		return &Array{
			data:   make([]interface{}, capacity, capacity),
			length: 0,
		}
	}
	return nil
}

// Len 返回数组长度
func (a *Array) Len() int {
	return a.length
}

func (a *Array) outOfRange(index int) bool {
	if index >= cap(a.data) || index < 0 {
		return true
	}
	return false
}

// Find 查找指定位置的元素
func (a *Array) Find(index int) interface{} {
	if a.outOfRange(index) {
		return nil
	}
	return a.data[index]
}

// Insert 将元素插入到指定位置
func (a *Array) Insert(index int, value interface{}) bool {
	if a.Len() == cap(a.data) {
		return false
	}
	if a.outOfRange(index) {
		return false
	}
	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = value
	a.length++
	return true
}

// InsertToTail 在尾部插入一个元素
func (a *Array) InsertToTail(value interface{}) bool {
	return a.Insert(a.length, value)
}

// InsertToHead 在头部插入一个元素
func (a *Array) InsertToHead(value interface{}) bool {
	return a.Insert(0, value)
}

// Delete 删除指定位置的元素
func (a *Array) Delete(index int) bool {
	if a.outOfRange(index) {
		return false
	}
	for i := index; i < a.length-1; i++ {
		a.data[index] = a.data[index+1]
	}
	a.length--
	return true
}

func (a *Array) String() string {
	if a.length == 0 {
		return fmt.Sprint([]int{})
	}
	return fmt.Sprint(a.data[0:a.length])

}
