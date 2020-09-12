package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//要求to必须已经分配好空间
func Uint32ToBytes(from uint32) (to []byte) {
	to = make([]byte, 4)
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, from)
	copy(to, buffer.Bytes()[0:4])
	return
}

func BytesToUint32(from []byte) (to uint32) {
	buffer := bytes.NewBuffer(from)
	binary.Read(buffer, binary.BigEndian, &to)
	return
}
func StringToInt(v string) (d int, err error) {
	tmp, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return
	}
	return int(tmp), err
}
func StringToUint(v string) (d uint, err error) {
	tmp, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		return
	}
	return uint(tmp), err
}
func StringToUint32(v string) (d uint32, err error) {
	tmp, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		return
	}
	return uint32(tmp), err
}
func StringToUint8(v string) (d uint8, err error) {
	tmp, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		return
	}
	return uint8(tmp), err
}
func StringToUint64(v string) (d uint64, err error) {
	d, err = strconv.ParseUint(v, 10, 64)
	return
}

func Uint64ToString(from uint64) (to string) {
	to = strconv.FormatUint(from, 10)
	return
}

func Float64ToString(from float64) (to string) {
	to = strconv.FormatFloat(from, 'f', -1, 64)
	return
}
func StringToFloat(v string) (d float32, err error) {
	tmp, err := strconv.ParseFloat(v, 10)
	d = float32(tmp)
	return
}
func StringToFloat64(v string) (d float64, err error) {
	d, err = strconv.ParseFloat(v, 10)
	return
}

func StringToInt64(v string) (d int64, err error) {
	d, err = strconv.ParseInt(v, 10, 64)
	return
}

func Int64ToString(from int64) (to string) {
	to = strconv.FormatInt(from, 10)
	return
}
func IntToString(from int) (to string) {
	to = strconv.FormatInt(int64(from), 10)
	return
}
func Uint32ToString(from uint32) (to string) {
	to = strconv.FormatInt(int64(from), 10)
	return
}

//把interface类型转换成string类型
func ToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func ToBool(v interface{}) (bool, error) {
	switch value := v.(type) {
	case bool:
		return value, nil
	case string:
		switch value {
		case "true", "True":
			return true, nil
		case "false", "False":
			return false, nil
		default:
			return false, errors.New("cannot convert " + value + " to bool")
		}
	case float32, float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return value != 0, nil

	default:
		return false, errors.New(fmt.Sprintf("cannot convert %v(%v) to bool", v, reflect.TypeOf(v)))
	}
}
func ToFloat64(v interface{}) (float64, error) {
	switch value := v.(type) {
	case string:
		return StringToFloat64(v.(string))
	case float32:
		return float64(value), nil
	case float64:
		return value, nil
	case int8:
		return float64(value), nil
	case int16:
		return float64(value), nil
	case int32:
		return float64(value), nil
	case int:
		return float64(value), nil
	case int64:
		return float64(value), nil
	case uint8:
		return float64(value), nil
	case uint16:
		return float64(value), nil
	case uint32:
		return float64(value), nil
	case uint:
		return float64(value), nil
	case uint64:
		return float64(value), nil
	case json.Number:
		return value.Float64()
	default:
		return 0, errors.New(fmt.Sprintf("cannot convert %v(%v) to float64", v, reflect.TypeOf(v)))
	}
}

func ToInt64(v interface{}) (int64, error) {
	switch value := v.(type) {
	case string:
		return StringToInt64(v.(string))
	case float32:
		return int64(value), nil
	case float64:
		return int64(value), nil
	case int8:
		return int64(value), nil
	case int16:
		return int64(value), nil
	case int32:
		return int64(value), nil
	case int:
		return int64(value), nil
	case int64:
		return value, nil
	case uint8:
		return int64(value), nil
	case uint16:
		return int64(value), nil
	case uint32:
		return int64(value), nil
	case uint:
		return int64(value), nil
	case uint64:
		return int64(value), nil
	case json.Number:
		return value.Int64()
	default:
		return 0, errors.New(fmt.Sprintf("cannot convert %v(%v) to float64", v, reflect.TypeOf(v)))
	}
}

func ToUint8(v interface{}) (uint8, error) {
	i, e := ToInt64(v)
	return uint8(i), e
}
func ToUint16(v interface{}) (uint16, error) {
	i, e := ToInt64(v)
	return uint16(i), e
}
func ToUint32(v interface{}) (uint32, error) {
	i, e := ToInt64(v)
	return uint32(i), e
}
func ToUint(v interface{}) (uint, error) {
	i, e := ToInt64(v)
	return uint(i), e
}
func ToUint64(v interface{}) (uint64, error) {
	i, e := ToInt64(v)
	return uint64(i), e
}
func ToInt8(v interface{}) (int8, error) {
	i, e := ToInt64(v)
	return int8(i), e
}
func ToInt16(v interface{}) (int16, error) {
	i, e := ToInt64(v)
	return int16(i), e
}
func ToInt(v interface{}) (int, error) {
	i, e := ToInt64(v)
	return int(i), e
}
func ToInt32(v interface{}) (int32, error) {
	i, e := ToInt64(v)
	return int32(i), e
}

func ToFloat32(v interface{}) (float32, error) {
	i, e := ToFloat64(v)
	return float32(i), e
}
func ToStringSlice(v interface{}) ([]string, error) {
	switch slice := v.(type) {
	case []string:
		return slice, nil
	case []interface{}:
		r := make([]string, 0, len(slice))
		for _, item := range slice {
			r = append(r, fmt.Sprintf("%v", item))
		}
		return r, nil
	case []uint64:
		r := make([]string, 0, len(slice))
		for _, item := range slice {
			r = append(r, fmt.Sprintf("%v", item))
		}
		return r, nil
	default:
		return nil, errors.New(fmt.Sprintf("cannot convert %v(%v) to []string", v, reflect.TypeOf(v)))
	}
}

func ToUint32Slice(v interface{}) ([]uint32, error) {
	switch slice := v.(type) {
	case []uint32:
		return slice, nil
	case []float64:
		r := make([]uint32, 0, len(slice))
		for _, item := range slice {
			if i, e := ToUint32(item); e != nil {
				return nil, e
			} else {
				r = append(r, i)

			}
			// r = append(r, fmt.Sprintf("%v", item))
		}
		return r, nil
	case []string:
		r := make([]uint32, 0, len(slice))
		for _, item := range slice {
			if i, e := ToUint32(item); e != nil {
				return nil, e
			} else {
				r = append(r, i)
			}
			// r = append(r, fmt.Sprintf("%v", item))
		}
		return r, nil
	case []interface{}:
		r := make([]uint32, 0, len(slice))
		for _, item := range slice {
			if i, e := ToUint32(item); e != nil {
				return nil, e
			} else {
				r = append(r, i)
			}
			// r = append(r, fmt.Sprintf("%v", item))
		}
		return r, nil
	default:
		return nil, errors.New(fmt.Sprintf("cannot convert %v(%v) to []string", v, reflect.TypeOf(v)))
	}
}

func ToUint64Slice(v interface{}) ([]uint64, error) {
	switch slice := v.(type) {
	case []uint64:
		return slice, nil
	case []float64:
		r := make([]uint64, 0, len(slice))
		for _, item := range slice {
			if i, e := ToUint64(item); e != nil {
				return nil, e
			} else {
				r = append(r, i)

			}
			// r = append(r, fmt.Sprintf("%v", item))
		}
		return r, nil
	case []string:
		r := make([]uint64, 0, len(slice))
		for _, item := range slice {
			if i, e := ToUint64(item); e != nil {
				return nil, e
			} else {
				r = append(r, i)
			}
		}
		return r, nil
	case []interface{}:
		r := make([]uint64, 0, len(slice))
		for _, item := range slice {
			if i, e := ToUint64(item); e != nil {
				return nil, e
			} else {
				r = append(r, i)
			}
		}
		return r, nil
	default:
		return nil, errors.New(fmt.Sprintf("cannot convert %v(%v) to []string", v, reflect.TypeOf(v)))
	}
}

func ToInt64Slice(v interface{}) ([]int64, error) {
	switch slice := v.(type) {
	case []int64:
		return slice, nil
	case []float64:
		r := make([]int64, 0, len(slice))
		for _, item := range slice {
			if i, e := ToInt64(item); e != nil {
				return nil, e
			} else {
				r = append(r, i)

			}
		}
		return r, nil
	case []string:
		r := make([]int64, 0, len(slice))
		for _, item := range slice {
			if i, e := ToInt64(item); e != nil {
				return nil, e
			} else {
				r = append(r, i)
			}
		}
		return r, nil
	case []interface{}:
		r := make([]int64, 0, len(slice))
		for _, item := range slice {
			if i, e := ToInt64(item); e != nil {
				return nil, e
			} else {
				r = append(r, i)
			}
		}
		return r, nil
	default:
		return nil, errors.New(fmt.Sprintf("cannot convert %v(%v) to []string", v, reflect.TypeOf(v)))
	}
}

func Join(v interface{}, sep string) (string, error) {
	switch slice := v.(type) {
	case []string:
		return strings.Join(slice, sep), nil
	case []uint32:
		if len(slice) == 0 {
			return "", nil
		}
		var buf bytes.Buffer
		_, _ = fmt.Fprintf(&buf, "%v", slice[0])
		for i := 1; i < len(slice); i++ {
			_, _ = fmt.Fprintf(&buf, "%s%v", sep, slice[i])
		}
		return buf.String(), nil
	case []uint64:
		if len(slice) == 0 {
			return "", nil
		}
		var buf bytes.Buffer
		_, _ = fmt.Fprintf(&buf, "%v", slice[0])
		for i := 1; i < len(slice); i++ {
			_, _ = fmt.Fprintf(&buf, "%s%v", sep, slice[i])
		}
		return buf.String(), nil
	case []interface{}:
		if len(slice) == 0 {
			return "", nil
		}
		var buf bytes.Buffer
		_, _ = fmt.Fprintf(&buf, "%v", slice[0])
		for i := 1; i < len(slice); i++ {
			_, _ = fmt.Fprintf(&buf, "%s%v", sep, slice[i])
		}
		return buf.String(), nil
	default:
		return "", errors.New(fmt.Sprintf("cannot convert %v(%v) to Slice", v, reflect.TypeOf(v)))
	}
}

//结构->MAP
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		return nil
	}
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := f.Tag.Get("json")
		if name == "" {
			name = f.Name
		}
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func Struct2Map2(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		return nil
	}
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := f.Tag.Get("json")
		if name == "" {
			name = f.Name
		}
		data[t.Field(i).Tag.Get("json")] = v.Field(i).Interface()
	}
	return data
}

func Bool2Int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func MakeKey(keys ...interface{}) string {
	if len(keys) == 0 {
		return ""
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%v", keys[0]))
	for i := 1; i < len(keys); i++ {
		buf.WriteString(fmt.Sprintf("_%v", keys[i]))
	}
	return buf.String()
}

func MakeKey2(keys ...interface{}) string {
	if len(keys) == 0 {
		return ""
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%v", keys[0]))
	for i := 1; i < len(keys); i++ {
		buf.WriteString(fmt.Sprintf("|%v", keys[i]))
	}
	return buf.String()
}

func SplitKey(key string) []string {
	return strings.Split(key, "_")
}

func SplitKey2(key string) []string {
	return strings.Split(key, "|")
}

func JSONMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
	b, err := json.Marshal(v)
	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}

func XmlMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
	b, err := xml.Marshal(v)
	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("&amp;"), []byte("&"), -1)
	}
	return b, err
}
