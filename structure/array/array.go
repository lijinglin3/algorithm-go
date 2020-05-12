package array

import (
	"fmt"
)

type Array struct {
	data   []interface{}
	length int
}

func NewArray(capacity int) *Array {
	if capacity > 0 {
		return &Array{
			data:   make([]interface{}, capacity, capacity),
			length: 0,
		}
	}
	return nil
}

func (a *Array) Len() int {
	return a.length
}

func (a *Array) outOfRange(index int) bool {
	if index >= cap(a.data) || index < 0 {
		return true
	}
	return false
}

func (a *Array) Find(index int) interface{} {
	if a.outOfRange(index) {
		return nil
	}
	return a.data[index]
}

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

func (a *Array) InsertToTail(value interface{}) bool {
	return a.Insert(a.length, value)
}

func (a *Array) InsertToHead(value interface{}) bool {
	return a.Insert(0, value)
}

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
