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
    b.elems = make([]Elem, n)
    b.held = make([]bool, n)
    b.bucket = make(chan int, n)
    for i := 0; i < n; i++ {
        b.bucket<-i
    }
    return b
}

func (b *Bucket) Init(f func(int) interface{}) {
    for _, p := range b.held {
        if p {
            panic("held")
        }
    }
    for i, _ := range b.elems {
        b.elems[i] = Elem{i, f(i)}
    }
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
    b.bucket <- i
    b.held[i] = false
}
