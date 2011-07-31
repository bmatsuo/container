package bucket
/*
 *  Filename:    bucket.go
 *  Package:     bucket
 *  Author:      Bryan Matsuo <bmatsuo@soe.ucsc.edu>
 *  Created:     Sat Jul 30 19:48:17 PDT 2011
 *  Description: 
 */
import ()

type Elem struct {
    id    int
    Value interface{}
}

type Bucket struct {
    held   []bool
    elems  []Elem
    bucket chan int
}

func New(n int) *Bucket {
    b := new(Bucket)
    b.bucket = make(chan int, n)
    b.elems = make([]Elem, n)
    b.held = make([]bool, n)
    return b
}

func (b *Bucket) Size() int {
    return len(b.elems)
}

func (b *Bucket) Retain() (int, interface{}) {
    if i, ok := <-b.bucket; ok {
        if b.held[i] {
            panic("held")
        }
        b.held[i] = true
        return i, b.elems[i].Value
    }
    panic("closed")
}

func (b *Bucket) Release(i int) {
    if !b.held[i] {
        panic("free")
    }
    b.bucket<-i
    b.held[i] = false
}
