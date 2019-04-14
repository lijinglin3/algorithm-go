package structure

import (
	"fmt"
)

type Array struct {
	data   []interface{}
	length uint
}

func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	} else {
		return &Array{
			data:   make([]interface{}, capacity, capacity),
			length: 0,
		}
	}
}

func (a *Array) Len() uint {
	return a.length
}

func (a *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(a.data)) {
		return true
	}
	return false
}

func (a *Array) Find(index uint) interface{} {
	if a.isIndexOutOfRange(index) {
		return nil
	}
	return a.data[index]
}

func (a *Array) Insert(index uint, value int) bool {
	if a.Len() == uint(cap(a.data)) {
		return false
	}
	if a.isIndexOutOfRange(index) {
		return false
	}
	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = value
	a.length++
	return true
}

func (a *Array) InsertToTail(value int) bool {
	return a.Insert(a.length, value)
}

func (a *Array) InsertToHead(value int) bool {
	return a.Insert(0, value)
}

func (a *Array) Delete(index uint) bool {
	if a.isIndexOutOfRange(index) {
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
		return fmt.Sprintln([]int{})
	} else {
		return fmt.Sprintln(a.data[0 : a.length-1])
	}
}
