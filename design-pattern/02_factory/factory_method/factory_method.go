package factory_method

type IRuleConfigParser interface {
	Parse([]byte)
}

type jsonRuleConfigParser struct{}

func (j jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type yamlRuleConfigParser struct{}

func (y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

type jsonRuleConfigParserFactory struct{}

func (j jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

type yamlRuleConfigParserFactory struct{}

func (y yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	}
	return nil
}
