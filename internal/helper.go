package internal

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

const totalLength = 80

var Verbose bool

func Log(format string, args ...any) {
	if Verbose {
		fmt.Println(format, args)
	}
}

func LogReq(req *resty.Request) {
	if Verbose {
		fmt.Println(strings.Repeat(">", totalLength))
		fmt.Println("URL: ", req.URL)
		fmt.Println("Method: ", req.Method)
		fmt.Println("Query Params: ", req.QueryParam)
		fmt.Println("FormData: ", req.FormData)
		fmt.Println("Body: ", req.Body)
		fmt.Println()
	}
}

func LogRes(res *resty.Response) {
	if Verbose {
		fmt.Println(strings.Repeat("<", totalLength))
		fmt.Println("URL: ", res.Request.URL)
		fmt.Println("Res: ", string(res.Body()))
		fmt.Println()
	}
}

func Info(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stdout, format+"\n", args...)
}

func Error(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

func CheckErr(err error) {
	cobra.CheckErr(err)
}

func ShowConfig(dataId string, content string) {
	const totalLength = 80
	paddingLength := (totalLength - len(dataId)) / 2
	fmt.Println(strings.Repeat("=", paddingLength) + dataId + strings.Repeat("=", totalLength-len(dataId)-paddingLength))
	fmt.Println(content)
	fmt.Println(strings.Repeat("=", totalLength))
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

func SaveConfig(dataId string, result string) {
	_ = os.WriteFile(dataId, []byte(result), 0644)
}

func Bool2String(success bool) string {
	if success {
		return "success"
	} else {
		return "fail"
	}
}

func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func Md5String(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5Bytes(bytes []byte) string {
	h := md5.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}
