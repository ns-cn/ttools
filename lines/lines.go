package main

import (
	"bufio"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/ns-cn/ttools"
	"github.com/spf13/cobra"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	// LINE 占位符
	LINE               = "#number"
	REG_LINERANGE      = "\\A\\[?\\d*[:,]\\d*\\]?$"
	REG_LINERANGESTART = "\\d+[:,]"
	REG_LINERANGEEND   = "[:,]\\d+"
)

var (
	// 数据源
	FilePath      = ""
	FromClipboard = false // 是否从粘贴板读取字符数据，优先级第一
	ToClipboard   = false // 是否将结果写入到粘贴板而不是命令行标准输出
	lineRange     = ""    // 行范围原始数据绑定
	RangeStart    = -1    // 行范围开始行
	RangeEnd      = -1    // 行范围结束行
	RangeCutStart = false // 行范围是否截断开始
	RangeCutEnd   = false // 行范围是否截断结束
)

var root = &cobra.Command{
	Short: "针对文本的逐行读取小工具",
	Long: `
针对文件内文本的逐行读取及简易编辑操作
`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func main() {
	initPrefix()
	initReplace()
	RegisterLineCommand(CmdPrefix)                           // 前缀
	RegisterLineCommand(CmdSuffix)                           // 后缀
	RegisterLineCommand(CmdReplace)                          // 替换操作
	RegisterLineCommand(CmdTrim)                             // 去除先后的空白字符
	RegisterLineCommand(CmdTrimLeft)                         // 去除左侧的空白字符
	RegisterLineCommand(CmdTrimRight)                        // 去除右侧的空白字符
	RegisterLineCommand(CmdSkipEmpty)                        // 跳过空白字符行
	root.AddCommand(ttools.VersionCommand("lines", VERSION)) // 打印版本号
	// 数据源
	_ = root.Execute()
}

func RegisterLineCommand(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&FilePath, "file", "F", "", "目标文件, 不指定则从管道中读取")
	cmd.Flags().BoolVarP(&FromClipboard, "fromClipboard", "C", false, "是否从粘贴板读取数据作为格式化数据的数据源，如果为true，则不需要指定file参数")
	cmd.Flags().BoolVarP(&ToClipboard, "toClipboard", "c", false, "是否将处理结果粘贴到粘贴板（默认输出到标准输出）")
	cmd.Flags().StringVarP(&lineRange, "range", "R", "", "具体操作的行范围，例如10:20，可使用[或]分别截断开始和结束，例如[10:20]截断并只保留10到20行的内容")
	root.AddCommand(cmd)
}

func parseRange() bool {
	if lineRange == "" {
		return true
	}
	reg := regexp.MustCompile(REG_LINERANGE)
	regStart := regexp.MustCompile(REG_LINERANGESTART)
	regEnd := regexp.MustCompile(REG_LINERANGEEND)
	if !reg.MatchString(lineRange) {
		handleErrString("错误的行范围格式", "[S:E]，其中[或]截断符可选")
		return false
	}
	RangeCutStart = strings.HasPrefix(lineRange, "[")
	RangeCutEnd = strings.HasSuffix(lineRange, "]")
	lineRangeStartCut := regStart.FindString(lineRange)
	lineRangeEndCut := regEnd.FindString(lineRange)
	if lineRangeStartCut != "" {
		RangeStart, _ = strconv.Atoi(lineRangeStartCut[:len(lineRangeStartCut)-1])
	}
	if lineRangeEndCut != "" {
		RangeEnd, _ = strconv.Atoi(lineRangeEndCut[1:])
	}
	return true
}

type Line struct {
	number int
	value  string
}

func InputAction(cmd *cobra.Command, action func(*bufio.Reader)) {
	parseWithoutErr := parseRange()
	if !parseWithoutErr {
		_ = cmd.Help()
		return
	}
	var reader *bufio.Reader
	if FromClipboard {
		clipboardString, err := clipboard.ReadAll()
		if err != nil {
			handleErrWithTips("读取粘贴板失败", err)
			return
		}
		reader = bufio.NewReader(strings.NewReader(clipboardString))
	} else if FilePath == "" {
		reader = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.OpenFile(FilePath, os.O_RDWR, 0666)
		if err != nil {
			handleErrWithTips("open file error!", err)
			return
		}
		defer file.Close()
		if err != nil {
			panic(err)
		}
		reader = bufio.NewReader(file)
	}
	action(reader)
}

/*
LineAction
行处理函数，需要传递单行文本处理函数供统筹处理
*/
func LineAction(cmd *cobra.Command, action func(line Line) string) {
	InputAction(cmd, func(reader *bufio.Reader) {
		result := strings.Builder{}
		var lineNumber = 1
		var handleLine = func(number int, line string) {
			inRange := (RangeStart == -1 || lineNumber >= RangeStart) && (RangeEnd == -1 || lineNumber <= RangeEnd)
			if inRange {
				actionResult := action(Line{number: lineNumber, value: line})
				result.WriteString(actionResult)
				//result.WriteString(fmt.Sprintf("%d%s", lineNumber, actionResult))
			} else {
				if RangeStart != -1 && lineNumber < RangeStart && !RangeCutStart {
					result.WriteString(line)
				}
				if RangeEnd != -1 && lineNumber > RangeEnd && !RangeCutEnd {
					result.WriteString(line)
				}
			}
		}
		var needToBackSpace = false
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					if line != "" {
						handleLine(lineNumber, line)
					} else {
						needToBackSpace = true
					}
					break
				} else {
					handleErrWithTips("Read file error!", err)
					return
				}
			}
			handleLine(lineNumber, line)
			lineNumber++
		}
		resultString := result.String()
		resultSize := len(resultString)
		if needToBackSpace && resultSize > 1 {
			resultString = resultString[:resultSize-1]
		}
		if ToClipboard {
			err := clipboard.WriteAll(resultString)
			if err != nil {
				_, _ = os.Stderr.WriteString(err.Error())
			}
		} else {
			fmt.Println(resultString)
		}
	})
}
