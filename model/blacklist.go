package model

import (
	"path"
)

type Blacklist struct {
	items []*BlacklistedItem
}

type BlacklistedItem struct {
	pattern string
}

func NewBlacklist(items []string) (*Blacklist, error) {
	var itemsBl []*BlacklistedItem
	for _, item := range items {
		_, err := path.Match(item, "")
		if err != nil {
			return nil, err
		}
		itemsBl = append(itemsBl, &BlacklistedItem{pattern: item})
	}
	return &Blacklist{
		items: itemsBl,
	}, nil
}

func (b *Blacklist) IsBlacklisted(name string) bool {
	for i := range b.items {
		if b.items[i].Match(name) {
			return true
		}
	}
	return false
}

func (bi *BlacklistedItem) Match(name string) bool {
	k, err := path.Match(bi.pattern, name)
	if err != nil {
		// should not happen since we test that the pattern is valid when adding the BlacklistedItem
		panic(err)
	}
	return k
}