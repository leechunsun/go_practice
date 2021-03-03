package set

import (
	"strings"
	"sync"
)

type StringSet struct {
	core map[string]bool
	sync.RWMutex
}

func (s *StringSet) init(){
	if s.core == nil{
		s.core = make(map[string]bool, 10)
	}
}

func (s *StringSet) Len() int {
	return len(s.core)
}


func (s *StringSet) Add (a string){
	s.Lock()
	defer s.Unlock()
	s.init()
	s.core[a] = false
}

func (s *StringSet) Contains (a string) bool {
	s.RLock()
	defer s.RUnlock()
	s.init()
	_, ok := s.core[a]
	return ok
}

func (s *StringSet) ToString () string {
	s.Lock()
	defer s.Unlock()
	s.init()
	sb := &strings.Builder{}
	sb.WriteString("{")
	for key, _ := range s.core{
		sb.WriteString(key + ", ")
	}
	subs := sb.String()
	subs = subs[:strings.LastIndex(subs, ",")] + "}"
	return subs
}

func (s *StringSet) Differ (s2 *StringSet)  *StringSet {
	s.Lock()
	defer s.Unlock()
	res := &StringSet{}
	res.init()
	for key, _ := range s.core{
		if ! s2.Contains(key) {
			res.Add(key)
		}
	}
	return res
}
