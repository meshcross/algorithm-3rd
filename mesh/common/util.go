/*
 * @Description: 类型转换
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-15 11:25:40
 * @LastEditTime: 2020-03-06 13:01:34
 * @LastEditors:
 */
package Common

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func ToString(d interface{}, flen ...int) string {

	if d == nil {
		return ""
	}

	jingdu := 2
	if len(flen) > 0 {
		jingdu = flen[0]
	}
	if jingdu < 0 {
		jingdu = 0
	}
	if jingdu > 8 {
		jingdu = 8
	}
	jstr := fmt.Sprintf("%d", jingdu)

	switch v := d.(type) {
	case int:
		{
			i := int64(v)
			return fmt.Sprintf("%d", i)
		}
	case int8:
		{
			i := int64(v)
			return fmt.Sprintf("%d", i)
		}
	case int32:
		{
			i := int64(v)
			return fmt.Sprintf("%d", i)
		}
	case int64:
		{
			i := int64(v)
			return fmt.Sprintf("%d", i)
		}
	case float32:
		{
			i := float64(v)
			return fmt.Sprintf("%."+jstr+"f", i)
		}
	case float64:
		{
			i := float64(v)
			return fmt.Sprintf("%."+jstr+"f", i)
		}
	case string:
		{
			return d.(string)
		}
	case map[string]interface{}:
		{
			return JsonToString(d.(map[string]interface{}))
		}
	default:
	}
	return ""
}
func ToBool(d interface{}) bool {
	if d == nil {
		return false
	}
	switch v := d.(type) {
	case bool:
		{
			i := bool(v)
			return i
		}
	}
	return false
}
func ToTime(d interface{}) time.Time {
	if d == nil {
		return time.Time{}
	}
	switch v := d.(type) {
	case time.Time:
		{
			i := time.Time(v)
			return i
		}
	}
	return time.Time{}
}

func ToInt(d interface{}) int {
	if d == nil {
		return 0
	}

	//switch v := d.(type) {
	switch v := d.(type) {
	case int:
		{
			i := int(v)
			return i
		}
	case int8:
		{
			i := int(v)
			return i
		}
	case int32:
		{
			i := int(v)
			return i
		}
	case int64:
		{
			i := int(v)
			return i
		}
	case uint32:
		{
			i := int(v)
			return i
		}
	case uint64:
		{
			i := int(v)
			return i
		}
	case float32:
		{
			i := int(v)
			return i
		}
	case float64:
		{
			i := int(v)
			return i
		}
	case string:
		{
			i, err := strconv.Atoi(d.(string))
			if err != nil {
				return 0
			}
			return int(i)
		}
	default:
	}
	return 0

}
func ToStringSlice(v interface{}, defaultv ...[]string) []string {
	var d []string = nil
	if len(defaultv) > 0 {
		d = defaultv[0]
	}

	if v == nil {
		return d
	}

	switch t := v.(type) {
	case []string:
		{
			return []string(t)
		}
	case []int64:
		{
			arr := []int64(t)
			ret := []string{}
			for _, v := range arr {
				ret = append(ret, fmt.Sprintf("%d", v))
			}
			return ret
		}
	}
	return d
}
func ToInt64Slice(v interface{}, defaultv ...[]int64) []int64 {
	var d []int64 = nil
	if len(defaultv) > 0 {
		d = defaultv[0]
	}

	if v == nil {
		return d
	}

	if v, ok := v.([]int64); ok {
		return v
	}
	return d
}
func ToInt64(d interface{}) int64 {
	if d == nil {
		return 0
	}

	//switch v := d.(type) {
	switch v := d.(type) {
	case int:
		{
			i := int64(v)
			return i
		}
	case int8:
		{
			i := int64(v)
			return i
		}
	case int32:
		{
			i := int64(v)
			return i
		}
	case int64:
		{
			i := int64(v)
			return i
		}
	case uint32:
		{
			i := int64(v)
			return i
		}
	case uint64:
		{
			i := int64(v)
			return i
		}
	case float32:
		{
			i := int64(v)
			return i
		}
	case float64:
		{
			i := int64(v)
			return i
		}
	case string:
		{
			i, err := strconv.Atoi(d.(string))
			if err != nil {
				return 0
			}
			return int64(i)
		}
	default:
	}
	return 0

}

func ToFloat64(d interface{}) float64 {
	if d == nil {
		return 0
	}

	//switch v := d.(type) {
	switch v := d.(type) {
	case int:
		{
			i := float64(v)
			return i
		}
	case int8:
		{
			i := float64(v)
			return i
		}
	case int32:
		{
			i := float64(v)
			return i
		}
	case int64:
		{
			i := float64(v)
			return i
		}
	case float32:
		{
			i := float64(v)
			return i
		}
	case float64:
		{
			i := float64(v)
			return i
		}
	case string:
		{
			i, err := strconv.ParseFloat(d.(string), 64)
			if err != nil {
				return 0
			}
			return float64(i)
		}
	default:
	}
	return 0
}

func MergeArrayMap(arrs ...[]map[string]interface{}) []map[string]interface{} {
	arr := []map[string]interface{}{}
	for _, iarr := range arrs {
		for _, v := range iarr {
			arr = append(arr, v)
		}
	}
	return arr
}

func MergeArrayInt64(arrs ...[]int64) []int64 {
	arr := []int64{}
	for _, iarr := range arrs {
		for _, v := range iarr {
			arr = append(arr, v)
		}
	}
	return arr
}

func GetTypeOf(a interface{}) string {
	return reflect.TypeOf(a).String()
}

func MergeMap(arrs ...map[string]interface{}) map[string]interface{} {
	m := map[string]interface{}{}
	for _, iarr := range arrs {
		for k, v := range iarr {
			m[k] = v
		}
	}
	return m
}

func Strings2Int64s(strs []string) []int64 {
	ret := []int64{}
	for _, v := range strs {
		ret = append(ret, String2Int64(v))
	}
	return ret
}
func Ints2Strings(is []int64) []string {
	ret := []string{}
	for _, v := range is {
		ret = append(ret, Int2String(v))
	}
	return ret
}
func String2Int64(str string) int64 {
	b, err := strconv.ParseInt(str, 10, 0)
	//b,err := strconv.Atoi(str)
	if err == nil {
		return int64(b)
	}
	return 0
}
func Int2String(i int64) string {
	return strconv.FormatInt(int64(i), 10)
}

func StringToJson(str string) map[string]interface{} {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(str), &dat); err == nil {
		return dat
	}
	return nil
}

func JsonToString(m map[string]interface{}) string {
	s, err := json.Marshal(m)
	if err == nil {
		return string(s)
	}
	return ""
}

func GetXString(str string, l int) string {
	arr := []string{}
	for k := 0; k < l; k++ {
		arr = append(arr, str)
	}
	return strings.Join(arr, "")
}
