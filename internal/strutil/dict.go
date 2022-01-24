package strutil

import (
	"sort"
	"strings"
)

// Dict is a simple wrapper for `map[string]struct{}`.
// Useful for keeping unique values.
type Dict map[string]struct{}

// NewDict returns a new Dict populated with the specified items.
func NewDict(items ...string) Dict {
	dict := Dict{}
	for _, item := range items {
		dict[item] = struct{}{}
	}
	return dict
}

// Add add items into the Dict. If the Dict has not been initialized,
// it will stop and items won't be added to Dict.
func (d Dict) Add(items ...string) {
	if d == nil {
		return
	}

	for _, item := range items {
		d[item] = struct{}{}
	}
}

// Remove removes items from the Dict.
func (d Dict) Remove(items ...string) {
	if d == nil {
		return
	}

	for _, item := range items {
		delete(d, item)
	}
}

// Contain checks if the specified item is registered inside the Dict.
func (d Dict) Contain(item string) bool {
	_, exist := d[item]
	return exist
}

// Items returns the sorted items that registered in the Dict.
func (d Dict) Items() []string {
	var items []string
	for item := range d {
		items = append(items, item)
	}
	sort.Strings(items)
	return items
}

// String returns the string representation of this Dict.
func (d Dict) String() string {
	return strings.Join(d.Items(), ", ")
}
