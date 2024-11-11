package internal

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// TotalLength is the total length of the log
const TotalLength = 80

// Verbose is verbose mode
var Verbose bool

// VerboseLog print verbose log
func VerboseLog(format string, args ...any) {
	if Verbose {
		fmt.Printf(format+"\n", args)
	}
}

// VerboseLogReq print request
func VerboseLogReq(req *resty.Request) {
	if Verbose {
		fmt.Println(strings.Repeat(">", TotalLength))
		fmt.Printf("URL: %s\n", req.URL)
		fmt.Printf("Method: %s\n", req.Method)
		fmt.Printf("Query Params: %+v\n", req.QueryParam)
		fmt.Printf("FormData: %+v\n", req.FormData)
		fmt.Printf("Body: %v\n", req.Body)
		fmt.Println()
	}
}

// VerboseLogRes print response
func VerboseLogRes(res *resty.Response) {
	if Verbose {
		fmt.Println(strings.Repeat("<", TotalLength))
		fmt.Printf("URL: %s\n", res.Request.URL)
		fmt.Printf("Res: %s\n", string(res.Body()))
		fmt.Println()
	}
}

// Info print info
func Info(format string, args ...any) {
	fmt.Printf(format+"\n", args...)
}

// Error print error
func Error(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

// CheckErr check error
func CheckErr(err error) {
	cobra.CheckErr(err)
}

// ShowConfig show config
func ShowConfig(dataId string, content string) {
	paddingLength := (TotalLength - len(dataId)) / 2
	fmt.Println(strings.Repeat("=", paddingLength) + dataId + strings.Repeat("=", TotalLength-len(dataId)-paddingLength))
	fmt.Println(content)
	fmt.Println(strings.Repeat("=", TotalLength))
}

// TableShow show table
func TableShow(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.AppendBulk(data)
	table.Render()
}

// GenData generate data
func GenData[T any](data *[]T, trans func(T) []string) [][]string {
	var result [][]string
	for _, item := range *data {
		result = append(result, trans(item))
	}
	return result
}

// SaveConfig save config
func SaveConfig(dataId string, result string) {
	_ = os.WriteFile(dataId, []byte(result), 0644)
}

// Bool2String bool to string
func Bool2String(success bool) string {
	if success {
		return "success"
	} else {
		return "fail"
	}
}

// ReadFile read file
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

// GenStringMD5 md5 string
func GenStringMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// GenBytesMD5 md5 bytes
func GenBytesMD5(bytes []byte) string {
	h := md5.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}
