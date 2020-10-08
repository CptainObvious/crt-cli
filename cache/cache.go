package cache

var emptyValue = struct{}{}
type Cache interface {
	Exist(value string) bool
	CacheValue(value string)
	Remove(value string)
}

type DomainCache struct {
	DomainsNames map[string]struct{}
}

func (c DomainCache) Exist(value string) bool  {
	if _, k := c.DomainsNames[value]; k {
		return true
	}
	return false
}

func (c DomainCache) CacheValue(value string) {
	c.DomainsNames[value] = emptyValue
}

func (c DomainCache) Remove(value string) {
	if c.Exist(value) {
		delete(c.DomainsNames, value)
	}
}

func NewDomainCache() Cache {
	return &DomainCache{
		DomainsNames: make(map[string]struct{}),
	}
}
