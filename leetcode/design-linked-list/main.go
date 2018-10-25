package main

import (
    "container/list"
    "github.com/davecgh/go-spew/spew"
)

type MyLinkedList struct {
    list.List
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    return MyLinkedList{}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    f := this.List.Front()
    for i:=0; f != nil && i<index; i++ {
        f = f.Next()
    }
    if f == nil {
        return -1
    }
    return  f.Value.(int)
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
    this.List.PushFront(val)
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
    this.List.PushBack(val)
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index == this.Len() {
        this.PushBack(val)
        return
    }
    f := this.List.Front()
    for i:=0; f != nil && i<index; i++ {
        f = f.Next()
    }
    if f != nil {
        this.List.InsertBefore(val, f)
    }
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    f := this.List.Front()
    for i:=0; f != nil && i<index; i++ {
        f = f.Next()
    }
    if f != nil {
        this.List.Remove(f)
    }
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

 func main() {
     obj := Constructor()
     obj.AddAtHead(1)
     obj.AddAtIndex(1, 2)
     spew.Dump(obj.Get(1),
     obj.Get(0),
     obj.Get(2))
 }