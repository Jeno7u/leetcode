package main

import (
	"math/rand"
)


type RandomizedSet struct {
	keys []int
    values map[int]int
}


func Constructor() RandomizedSet {
	randomizedSet := RandomizedSet{keys: make([]int, 0), values: make(map[int]int)}
    return randomizedSet 
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, present := this.values[val]; !present {
		this.keys = append(this.keys, val)
		this.values[val] = len(this.keys) - 1
		return true
	}
	return false
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, present := this.values[val]; present {
		idxToDelete := this.values[val]
		this.keys[idxToDelete] = this.keys[len(this.keys) - 1]
		this.values[this.keys[len(this.keys) - 1]] = idxToDelete
		this.keys = this.keys[:len(this.keys)-1]
		delete(this.values, val)
		return true
	}
	return false
}


func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.keys))
	return this.keys[idx]
}

// основная проблема что нужно сделать getRandom за O(1). Если использовать HashSet, то чтобы 
// получить значения и потом выбрать случайное нужно O(n), а это плохо. Значит нужно сохранять
// эти ключи. Нам нужна упорядоченная структура чтобы мы могли по случайному индексу возвращать
// getRandom. Значит используем slice в котором val будут храниться. Но теперь проблема в удалении
// за O(n). Нам нужно найи элемент который мы будем удалять хотя бы. Так давайте использовать
// HashMap вместо HashSet и хранить как значение индекс в списке. Осталось как удалять за O(1).
// Это можно сделать просто заменяя элемент который нужно удалить на последний элемент в списке.
// Нам на порядок плевать, но важно исправить индексы в HashMap на новые
