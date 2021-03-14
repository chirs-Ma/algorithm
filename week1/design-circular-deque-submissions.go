//采用数组+头尾指针解决
type MyCircularDeque struct {
	que    []int
	nums   int
	length int
	head   int
	foot   int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		que:    make([]int, k),
		head:   1,
		length: k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.nums == this.length {
		return false
	}
	this.head--
	if this.head == -1 {
		this.head = this.length - 1
	}
	this.que[this.head] = value
	this.nums++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.nums == this.length {
		return false
	}
	this.foot++
	if this.foot == this.length {
		this.foot = 0
	}
	this.que[this.foot] = value
	this.nums++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.nums == 0 {
		return false
	}
	this.head++
	if this.head == this.length {
		this.head = 0
	}
	this.nums--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.nums == 0 {
		return false
	}
	this.foot--
	if this.foot == -1 {
		this.foot = this.length - 1
	}
	this.nums--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.nums == 0 {
		return -1
	}
	return this.que[this.head]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.nums == 0 {
		return -1
	}
	return this.que[this.foot]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.nums == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.nums == this.length
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */