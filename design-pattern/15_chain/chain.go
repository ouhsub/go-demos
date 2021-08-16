package chain

type SensitiveWordFilter interface {
	Filter(content string) bool
}

type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

func (chain *SensitiveWordFilterChain) AddFilter(filter SensitiveWordFilter) {
	chain.filters = append(chain.filters, filter)
}

func (chain *SensitiveWordFilterChain) Filter(content string) bool {
	for _, filter := range chain.filters {
		if filter.Filter(content) {
			return true
		}
	}
	return false
}

type AdSensitiveFilter struct{}

func (filter *AdSensitiveFilter) Filter(content string) bool {
	return false
}

type PoliticalSensitiveFilter struct{}

func (filter *PoliticalSensitiveFilter) Filter(content string) bool {
	return false
}
