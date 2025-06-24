package model

type Rule struct {
	Type       RuleType `json:"type"`
	Expression string   `json:"expression"`
}

type RuleType string

const (
	RuleTypeJavascript RuleType = "javascript"
)
