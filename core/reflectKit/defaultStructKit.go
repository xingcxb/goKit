// Package reflectKit Default将默认值应用于零值的struct字段。
// 本文件引用zzjcool/goutils项目，引用地址为：https://github.com/zzjcool/goutils/blob/main/defaults/defauls.go
package reflectKit

import (
	"fmt"
	"reflect"
	"strconv"
)

// ErrNotAStructPointer 一个指向struct的指针，
type ErrNotAStructPointer string

func newErrNotAStructPointer(v interface{}) ErrNotAStructPointer {
	return ErrNotAStructPointer(fmt.Sprintf("%t", v))
}

// Error 实现错误接口
func (e ErrNotAStructPointer) Error() string {
	return fmt.Sprintf("expected a struct, instead got a %T", string(e))
}

// ErrorUnsettable 当不能设置字段时使用
type ErrorUnsettable string

// Error 实现错误接口
func (e ErrorUnsettable) Error() string {
	return fmt.Sprintf("can't set field %s", string(e))
}

// ErrorUnsupportedType 指示结构字段的类型还不是此包中的支持
type ErrorUnsupportedType struct {
	t reflect.Type
}

// Error 实现错误接口
func (e ErrorUnsupportedType) Error() string {
	return fmt.Sprintf("unsupported type %v", e.t)
}

// StructDefault 给结构体增加默认值
/**
 * @param t 结构体
 * @return 返回错误信息
 */
func StructDefault(t interface{}) error {
	val := reflect.ValueOf(t)
	if val.Kind() != reflect.Ptr {
		return newErrNotAStructPointer(t)
	}

	// 确保指针指向一个结构体
	ref := val.Elem()
	if ref.Kind() != reflect.Struct {
		return newErrNotAStructPointer(t)
	}

	return parseFields(ref)
}

func parseFields(v reflect.Value) error {
	for i := 0; i < v.NumField(); i++ {
		err := parseField(v.Field(i), v.Type().Field(i))
		if err != nil {
			return err
		}
	}
	return nil
}

func parseField(value reflect.Value, field reflect.StructField) error {
	tagVal := field.Tag.Get("default")

	isStruct := value.Kind() == reflect.Struct
	isStructPointer := value.Kind() == reflect.Ptr && value.Type().Elem().Kind() == reflect.Struct

	if (tagVal == "" || tagVal == "-") && !(isStruct || isStructPointer) {
		return nil
	}

	if !value.CanSet() {
		return nil
	}

	if !value.IsZero() {
		// 在这个字段上设置了一个值，所以不需要设置默认值
		return nil
	}

	switch value.Kind() {
	case reflect.String:
		value.SetString(tagVal)
		return nil
	case reflect.Bool:
		b, err := strconv.ParseBool(tagVal)
		if err != nil {
			return err
		}
		value.SetBool(b)
		return nil
	case reflect.Int:
		i, err := strconv.ParseInt(tagVal, 10, 32)
		if err != nil {
			return err
		}
		value.SetInt(i)
		return nil
	case reflect.Int8:
		i, err := strconv.ParseInt(tagVal, 10, 8)
		if err != nil {
			return err
		}
		value.SetInt(i)
		return nil
	case reflect.Int16:
		i, err := strconv.ParseInt(tagVal, 10, 16)
		if err != nil {
			return err
		}
		value.SetInt(i)
		return nil
	// 注意: int32 也可能是 rune
	case reflect.Int32:
		i, err := parseInt32(tagVal)
		if err != nil {
			return err
		}
		value.SetInt(int64(i))
		return nil
	case reflect.Int64:
		i, err := strconv.ParseInt(tagVal, 10, 64)
		if err != nil {
			return err
		}
		value.SetInt(i)
		return nil
	case reflect.Uint:
		i, err := strconv.ParseInt(tagVal, 10, 32)
		if err != nil {
			return err
		}
		value.SetUint(uint64(i))
		return nil
	case reflect.Uint8:
		i, err := strconv.ParseInt(tagVal, 10, 8)
		if err != nil {
			return err
		}
		value.SetUint(uint64(i))
		return nil
	case reflect.Uint16:
		i, err := strconv.ParseInt(tagVal, 10, 16)
		if err != nil {
			return err
		}
		value.SetUint(uint64(i))
		return nil
	case reflect.Uint32:
		i, err := strconv.ParseInt(tagVal, 10, 32)
		if err != nil {
			return err
		}
		value.SetUint(uint64(i))
		return nil
	case reflect.Uint64:
		i, err := strconv.ParseInt(tagVal, 10, 64)
		if err != nil {
			return err
		}
		value.SetUint(uint64(i))
		return nil
	case reflect.Float32:
		f, err := strconv.ParseFloat(tagVal, 32)
		if err != nil {
			return err
		}
		value.SetFloat(f)
		return nil
	case reflect.Float64:
		f, err := strconv.ParseFloat(tagVal, 64)
		if err != nil {
			return err
		}
		value.SetFloat(f)
		return nil
	case reflect.Slice:
		switch value.Type().Elem().Kind() {
		// 把[]uint8作为[]byte的别名
		case reflect.Uint8:
			value.SetBytes([]byte(tagVal))
			return nil
		default:
			return ErrorUnsupportedType{value.Type()}
		}
	case reflect.Struct:
		if value.NumField() == 0 {
			return nil
		}
		return parseFields(value) // 递归
	case reflect.Ptr:
		ref := value.Type().Elem()
		switch ref.Kind() {
		case reflect.String:
			value.Set(reflect.ValueOf(&tagVal))
			return nil
		case reflect.Bool:
			b, err := strconv.ParseBool(tagVal)
			if err != nil {
				return err
			}
			value.Set(reflect.ValueOf(&b))
			return nil
		case reflect.Int:
			n, err := strconv.ParseInt(tagVal, 10, 32)
			if err != nil {
				return err
			}
			i := int(n)
			value.Set(reflect.ValueOf(&i))
			return nil
		case reflect.Int8:
			n, err := strconv.ParseInt(tagVal, 10, 8)
			if err != nil {
				return err
			}
			i := int8(n)
			value.Set(reflect.ValueOf(&i))
			return nil
		case reflect.Int16:
			n, err := strconv.ParseInt(tagVal, 10, 16)
			if err != nil {
				return err
			}
			i := int16(n)
			value.Set(reflect.ValueOf(&i))
			return nil
		case reflect.Int32:
			// NB: *int32 is an alias for a *rune
			i, err := parseInt32(tagVal)
			if err != nil {
				return err
			}
			value.Set(reflect.ValueOf(&i))
			return nil
		case reflect.Int64:
			i, err := strconv.ParseInt(tagVal, 10, 64)
			if err != nil {
				return err
			}
			value.Set(reflect.ValueOf(&i))
			return nil
		case reflect.Uint:
			n, err := strconv.ParseInt(tagVal, 10, 32)
			if err != nil {
				return err
			}
			u := uint(n)
			value.Set(reflect.ValueOf(&u))
			return nil
		case reflect.Uint8:
			n, err := strconv.ParseInt(tagVal, 10, 8)
			if err != nil {
				return err
			}
			u := uint8(n)
			value.Set(reflect.ValueOf(&u))
			return nil
		case reflect.Uint16:
			n, err := strconv.ParseInt(tagVal, 10, 16)
			if err != nil {
				return err
			}
			u := uint16(n)
			value.Set(reflect.ValueOf(&u))
			return nil
		case reflect.Uint32:
			n, err := strconv.ParseInt(tagVal, 10, 32)
			if err != nil {
				return err
			}
			u := uint32(n)
			value.Set(reflect.ValueOf(&u))
			return nil
		case reflect.Uint64:
			n, err := strconv.ParseInt(tagVal, 10, 64)
			if err != nil {
				return err
			}
			u := uint64(n)
			value.Set(reflect.ValueOf(&u))
			return nil
		case reflect.Float32:
			f, err := strconv.ParseFloat(tagVal, 32)
			if err != nil {
				return err
			}
			f32 := float32(f)
			value.Set(reflect.ValueOf(&f32))
			return nil
		case reflect.Float64:
			f, err := strconv.ParseFloat(tagVal, 64)
			if err != nil {
				return err
			}
			value.Set(reflect.ValueOf(&f))
			return nil
		case reflect.Slice:
			switch ref.Elem().Kind() {
			// *[]uint作为*[]byte的别名
			case reflect.Uint8:
				b := []byte(tagVal)
				value.Set(reflect.ValueOf(&b))
				return nil
			default:
				return ErrorUnsupportedType{value.Type()}
			}
		case reflect.Struct:
			if ref.NumField() == 0 {
				return nil
			}
			// 如果是nil，就设为它的默认值，这样我们就可以设置
			// 如果需要子节点的话
			if value.IsNil() {
				value.Set(reflect.New(ref))
			}
			return parseFields(value.Elem()) // recurse
		default:
			return ErrorUnsupportedType{value.Type()}
		}
	default:
		return ErrorUnsupportedType{value.Type()}
	}
}

// 尝试将字符串解析为int32，如果失败，就使用rune
func parseInt32(s string) (int32, error) {
	// Try parsing it as an int.
	i, err := strconv.ParseInt(s, 10, 32)
	if err == nil {
		return int32(i), nil
	}

	// 我们不能把它解析为int型, 也许是rune.
	runes := []rune(s)
	if len(runes) == 1 {
		return runes[0], nil
	} else {
		return 0, err
	}
}
