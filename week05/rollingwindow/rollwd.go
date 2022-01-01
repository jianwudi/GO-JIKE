package rollingwindow

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"unsafe"
)

type Bucket struct {
	requestNum uint64
	LastTime   time.Time
}
type Window struct {
	requestNum uint64
}
type Windows struct {
	buckets   []*Bucket
	duration  time.Duration
	bucketnum uint64
	idx       uint64
	//	wds       []*Window
	wdMutex sync.Mutex
}

func NewWindows(windua time.Duration, bucketdua time.Duration) (*Windows, error) {
	if windua%bucketdua != 0 {
		return nil, errors.New("window duration%bucket duration != 0")
	}
	var bucketnum uint64
	buckets := []*Bucket{}
	num := windua / bucketdua
	bucketnum = *(*uint64)(unsafe.Pointer(&num))
	var i uint64
	for i = 0; i < bucketnum; i++ {
		buckets = append(buckets, &Bucket{
			LastTime:   time.Now(),
			requestNum: 0,
		})
	}
	return &Windows{
		buckets:   buckets,
		duration:  bucketdua,
		bucketnum: bucketnum,
		idx:       0,
		//	wds:       []*Window{},
	}, nil
}

func (w *Windows) AddRequest() {
	w.wdMutex.Lock()
	defer w.wdMutex.Unlock()
	w.buckets[w.idx%w.bucketnum].requestNum++
}

func (w *Windows) WindowCountTimer() {
	ticker := time.NewTicker(w.duration)
	for {
		select {
		case <-ticker.C:
			w.wdMutex.Lock()
			if w.idx >= w.bucketnum-1 {
				wd := &Window{requestNum: 0}
				for _, b := range w.buckets {
					wd.requestNum += b.requestNum
				}
				fmt.Printf("request NUM :%d \n", wd.requestNum)
				w.buckets[(w.idx+1)%(w.bucketnum)].requestNum = 0
				w.idx++
			} else {
				fmt.Printf("request idx :%d \n", w.idx)
				w.idx++
			}
			w.wdMutex.Unlock()
		}
	}
}
