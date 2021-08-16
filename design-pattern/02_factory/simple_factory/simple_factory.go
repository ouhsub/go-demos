package factory

// IRuleConfigParser define interface for config parser
type IRuleConfigParser interface {
	Parse([]byte)
}

type jsonRuleConfigParser struct {
}

// Parse implement IRuleConfigParser with json data
func (j jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type yamlRuleConfigParser struct {
}

// Parse implement IRuleConfigParser with yaml data
func (y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// NewRuleConfigParse return IRuleConfigParser instance
func NewRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}
