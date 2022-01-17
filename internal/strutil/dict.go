package strutil

import (
	"sort"
	"strings"
)

type Dict map[string]struct{}

func NewDict(items ...string) Dict {
	dict := Dict{}
	for _, item := range items {
		dict[item] = struct{}{}
	}
	return dict
}

func (d Dict) Add(items ...string) {
	if d == nil {
		return
	}

	for _, item := range items {
		d[item] = struct{}{}
	}
}

func (d Dict) Remove(items ...string) {
	if d == nil {
		return
	}

	for _, item := range items {
		delete(d, item)
	}
}

func (d Dict) Contain(item string) bool {
	_, exist := d[item]
	return exist
}

func (d Dict) String() string {
	var keys []string
	for key := range d {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	return strings.Join(keys, ", ")
}
