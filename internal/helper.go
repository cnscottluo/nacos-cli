package internal

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"reflect"
)

var Verbose bool

func Log(format string, args ...interface{}) {
	if Verbose {
		_, _ = fmt.Fprintf(os.Stdout, format+"\n", args...)
	}
}

func Info(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, format+"\n", args...)
}

func Error(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

func CheckErr(err error) {
	cobra.CheckErr(err)
}

func Struct2StringMap(s interface{}) map[string]string {
	result := make(map[string]string)
	v := reflect.ValueOf(s)

	// 检查是否为结构体类型
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			value := v.Field(i)
			tagName := field.Tag.Get("json")
			if tagName == "" {
				tagName = field.Name // 如果没有标签，则使用字段名称
			}
			result[tagName] = fmt.Sprintf("%v", value.Interface())
		}
	}
	return result
}

func TableShow(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.AppendBulk(data)
	table.Render()
}

func GenData[T any](data *[]T, trans func(T) []string) [][]string {
	var result [][]string
	for _, item := range *data {
		result = append(result, trans(item))
	}
	return result
}
