package validation

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// Error 是与 HTTP 无关的字段校验错误。
type Error struct {
	Field   string
	Rule    string
	Param   string
	Message string
	cause   error
}

func (e *Error) Error() string { return e.Message }
func (e *Error) Unwrap() error { return e.cause }

// Wrap 用于单个字段的校验，返回首个失败规则；非校验错误原样返回。
func Wrap(field, label string, err error) error {
	if err == nil {
		return nil
	}
	var failures validator.ValidationErrors
	if !errors.As(err, &failures) || len(failures) == 0 {
		return err
	}
	failure := failures[0]
	message := label + "格式不正确"
	switch failure.Tag() {
	case "required":
		message = label + "不能为空"
	case "min", "max", "len", "gt", "gte", "lt", "lte":
		message = comparisonMessage(label, failure)
	case "eq":
		message = fmt.Sprintf("%s必须等于%s", label, failure.Param())
		if isCollection(failure.Kind()) {
			message = fmt.Sprintf("%s元素数量必须等于%s", label, failure.Param())
		}
	case "ne":
		message = fmt.Sprintf("%s不能等于%s", label, failure.Param())
		if isCollection(failure.Kind()) {
			message = fmt.Sprintf("%s元素数量不能等于%s", label, failure.Param())
		}
	case "oneof":
		message = fmt.Sprintf("%s必须是以下选项之一：%s", label, failure.Param())
	case "email":
		message = label + "必须是有效的邮箱地址"
	case "url", "http_url":
		message = label + "必须是有效的网址"
	case "ip", "ipv4", "ipv6":
		message = label + "必须是有效的 " + failure.Tag() + " 地址"
	case "uuid", "uuid3", "uuid4", "uuid5":
		message = label + "必须是有效的 " + failure.Tag()
	case "datetime":
		message = fmt.Sprintf("%s必须符合日期时间格式：%s", label, failure.Param())
	case "numeric":
		message = label + "必须是有效的数字"
	case "number":
		message = label + "只能包含数字字符"
	case "alpha":
		message = label + "只能包含英文字母"
	case "alphanum":
		message = label + "只能包含英文字母和数字"
	case "lowercase":
		message = label + "必须为小写形式"
	case "uppercase":
		message = label + "必须为大写形式"
	case "contains":
		message = fmt.Sprintf("%s必须包含%q", label, failure.Param())
	case "excludes":
		message = fmt.Sprintf("%s不能包含%q", label, failure.Param())
	case "startswith":
		message = fmt.Sprintf("%s必须以%q开头", label, failure.Param())
	case "endswith":
		message = fmt.Sprintf("%s必须以%q结尾", label, failure.Param())
	case "eqfield":
		message = fmt.Sprintf("%s必须与字段 %s 一致", label, failure.Param())
	case "nefield":
		message = fmt.Sprintf("%s不能与字段 %s 相同", label, failure.Param())
	case "unique":
		message = label + "不能包含重复项"
	case "json":
		message = label + "必须是有效的 JSON"
	}
	return &Error{
		Field: field, Rule: failure.Tag(), Param: failure.Param(),
		Message: message, cause: err,
	}
}

// comparisonMessage 区分文本长度、集合数量与数值，避免将长度描述为数值范围。
func comparisonMessage(label string, failure validator.FieldError) string {
	subject := label
	switch failure.Kind() {
	case reflect.String:
		subject += "字数"
	case reflect.Slice, reflect.Array, reflect.Map:
		subject += "元素数量"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64:
	default:
		// 时间等类型的比较语义不同，不能将空参数描述为数值边界。
		return label + "不符合范围要求"
	}
	operator := map[string]string{
		"min": "不能小于", "max": "不能大于", "len": "必须等于",
		"gt": "必须大于", "gte": "不能小于", "lt": "必须小于", "lte": "不能大于",
	}[failure.Tag()]
	if failure.Kind() == reflect.String || isCollection(failure.Kind()) {
		if failure.Tag() == "min" || failure.Tag() == "gte" {
			operator = "不能少于"
		} else if failure.Tag() == "max" || failure.Tag() == "lte" {
			operator = "不能超过"
		}
	}
	return subject + operator + failure.Param()
}

func isCollection(kind reflect.Kind) bool {
	return kind == reflect.Slice || kind == reflect.Array || kind == reflect.Map
}
