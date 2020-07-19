package cmd

import (
	"log"
	"strings"

	"github.com/damingerdai/tour/internal/word"

	"github.com/spf13/cobra"
)

const (
	MODE_UPPER                         = iota + 1 // 全部单词转为大写
	MODE_LOWER                                    // 全部单词转为小写
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE            // 下划线单词转为大写驼峰单词
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE            // 下划线单词转为小写驼峰单词
	MODE_CAMELCASE_TO_UNDERSCORE                  // 驼峰单词转为下划线单词
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式， 模式如下：",
	"1: 全部单词转为大写",
	"2: 全部单词转为小写",
	"3: 下划线单词转为大写驼峰单词",
	"4: 下划线单词转为小写驼峰单词",
	"5: 驼峰单词转为下划线单词",
}, "\n")

var str string
var mode int8

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		log.Printf("模式 %d", mode)
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = word.UndersocreToUpperCamelCase(str)
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = word.UndersoreToLowerCamelCase(str)
		default:
			log.Fatalf("暂时不支持该转换模式，请执行help word 查看帮助文档")
		}
		log.Printf("输出结果： %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的格式")
}
