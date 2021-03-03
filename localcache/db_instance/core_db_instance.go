package db_instance

import (
	"sync"
	"time"
)

var CacheDbSingle *CacheDB

type CacheDB struct {
	coreMap map[string]interface{}
	deadLineMap map[string]int64
	sync.Mutex
}

func (c *CacheDB) Save (key string, val interface{}, deadLine int64) {
	// 先处理一下deadline
	dline := time.Now().Unix()
	if deadLine != 0 {
		dline += deadLine
	}
	c.Lock()
	defer c.Unlock()
	// 然后修改值
	c.coreMap[key] = val
	c.deadLineMap[key] = dline
}

func (c *CacheDB) Get (key string) interface{} {
	if dline, ok := c.deadLineMap[key]; (!ok) || dline < time.Now().Unix() {
		if ok {
			c.Lock()
			delete(c.deadLineMap, key)
			delete(c.coreMap, key)
			c.Unlock()
		}
		return nil
	}
	if val, ok := c.coreMap[key]; ok{
		return val
	}
	return nil
}

func (c *CacheDB) CheckOverline(){
	overdue := make([]string, 0)
	for key, dline := range c.deadLineMap{
		if dline < time.Now().Unix() {
			overdue = append(overdue, key)
		}
	}
	if len(overdue) > 0 {
		c.Lock()
		defer c.Unlock()
		for _, key := range overdue{
			delete(c.deadLineMap, key)
			delete(c.coreMap, key)
		}
	}
}

func (c *CacheDB) Clear () {
	c.Lock()
	defer c.Unlock()
	c.deadLineMap = make(map[string]int64, 100)
	c.coreMap = make(map[string]interface{}, 100)
}

func GetCacheDBSingle() *CacheDB {
	if CacheDbSingle == nil{
		CacheDbSingle = &CacheDB{
			coreMap:     make(map[string]interface{}, 100),
			deadLineMap: make(map[string]int64, 100),
			Mutex:       sync.Mutex{},
		}
	}
	return CacheDbSingle
}

