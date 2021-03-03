package coreStruct

import (
	"fmt"
	"hash/crc32"
	"sort"
	"sync"
)

type units []uint32

func (u units) Len() int {
	return len(u)
}

func (u units) Less (x, j int) bool {
	return u[x] < u[j]
}

func (u units) Swap (x, j int) {
	u[x], u[j] = u[j], u[x]
}


type HashCircular struct {
	circle map[uint32]string
	sortedHashes  units
	VirtualNodeNum int
	sync.Mutex
}

// 构造函数
func NewHashCircular () *HashCircular {
	return &HashCircular{
		circle: make(map[uint32]string),
		VirtualNodeNum: 20,
	}
}

// 生成虚拟节点的key
func (h *HashCircular) generateKey(element string, index int) string {
	return fmt.Sprintf("%s%d", element, index)
}

// 根据key生成hash值
func (h *HashCircular) hashKey (key string) uint32 {
	if len(key) < 64{
		var srcatch [64]byte
		copy(srcatch[:], key)
		return crc32.ChecksumIEEE(srcatch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

// 更新SortedHashes
func (h *HashCircular) updateSortedHashes() {
	hashes := h.sortedHashes[:0]
	standardCap := (h.VirtualNodeNum*4) *len(h.circle)
	if len(h.circle) > standardCap{
		hashes = make([]uint32, len(h.circle))
	}
	for key := range h.circle{
		hashes = append(hashes, key)
	}
	sort.Sort(hashes)
	h.sortedHashes = hashes
}

// 添加节点
func (h *HashCircular) add (element string){
	for i:=0;i<h.VirtualNodeNum;i++{
		key := h.generateKey(element, i)
		hashKey := h.hashKey(key)
		h.circle[hashKey] = element
		h.updateSortedHashes()
	}
}

// 添加外部接口
func (h *HashCircular) Add (element string) {
	h.Lock()
	defer h.Unlock()
	h.add(element)
}

// 删除节点
func (h *HashCircular) remove (element string) {
	for i:=0;i<h.VirtualNodeNum;i++{
		key := h.generateKey(element, i)
		kayHash := h.hashKey(key)
		delete(h.circle, kayHash)
	}
	h.updateSortedHashes()
}

// 删除节点
func (h *HashCircular) Remove (element string){
	h.Lock()
	defer h.Unlock()
	h.remove(element)
}

// 查找最近节点
func (h *HashCircular) search (ind uint32) int {

	f := func(x int) bool {
		return h.sortedHashes[x] > ind
	}
	i := sort.Search(len(h.sortedHashes), f)
	if i >= len(h.sortedHashes){
		i = 0
	}
	return i
}


// 查找值
func (h *HashCircular) Get (element string) (string, error) {
	if len(h.circle) == 0 {
		return "", nil
	}
	keyHash := h.hashKey(element)
	keyInd := h.search(keyHash)
	return h.circle[h.sortedHashes[keyInd]], nil
}

