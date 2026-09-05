package validation

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

var engine = validator.New()

// StringValidator 是不可变的文本校验器，可复用并并发校验。
// 链式方法返回新实例，调用方需要使用返回值。
type StringValidator struct {
	field, label, value string
	rules               []string
	configErr           error
}

// String 创建文本校验器，不修改原始值。长度按 Unicode 字符计算。
func String(field, label, value string) *StringValidator {
	return &StringValidator{field: field, label: label, value: value}
}

func (v *StringValidator) clone() *StringValidator {
	next := *v
	next.rules = append([]string(nil), v.rules...)
	return &next
}

// WithValue 为已有规则绑定新文本，不改变原实例。
func (v *StringValidator) WithValue(value string) *StringValidator {
	next := v.clone()
	next.value = value
	return next
}

// Required 要求文本包含非空白字符。
func (v *StringValidator) Required() *StringValidator {
	next := v.clone()
	next.rules = append(next.rules, "required")
	return next
}

func (v *StringValidator) Min(length int) *StringValidator {
	return v.lengthRule("min", length)
}

func (v *StringValidator) Max(length int) *StringValidator {
	return v.lengthRule("max", length)
}

func (v *StringValidator) lengthRule(rule string, length int) *StringValidator {
	next := v.clone()
	if length < 0 {
		if next.configErr == nil {
			next.configErr = fmt.Errorf("validation: %s length must be non-negative", rule)
		}
		return next
	}
	next.rules = append(next.rules, rule+"="+strconv.Itoa(length))
	return next
}

// Validate 按声明顺序校验并返回首个错误。长度包含首尾空白；
// 如需去除空白，应由调用方先 TrimSpace。配置错误不转为字段校验错误。
func (v *StringValidator) Validate() error {
	if v.configErr != nil {
		return v.configErr
	}
	for _, rule := range v.rules {
		value := v.value
		if rule == "required" {
			value = strings.TrimSpace(value)
		}
		if err := engine.Var(value, rule); err != nil {
			return Wrap(v.field, v.label, err)
		}
	}
	return nil
}
