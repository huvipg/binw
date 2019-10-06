//
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/huvipg/binw/color"
	
	_ "github.com/huvipg/binw/dir"
	"github.com/urfave/cli"
)

var p = fmt.Println

func Showhelp() {

	color.Yellow("序号:0 ")
	color.Green("设置指定生成目录\n")
	color.Yellow("序号:1 ")
	color.Green("生成表markdown文档\n")
	color.Yellow("序号:2 ")
	color.Green("生成表结构实体\n")
	color.Yellow("序号:3 ")
	color.Green("生成CURD增删改查\n")
	color.Yellow("序号:4 ")
	color.Green("设置结构体的映射名称\n")
	color.Yellow("序号:5, e, q, exit")
	color.Green(" 退出\n")
}
func scant() {
	Showhelp()
	for {
		var err error

		color.Blue("\n请输入以上命令序号:")
		input, _, err := bufio.NewReader(os.Stdin).ReadLine()
		if err != nil {
			os.Exit(9)
		}
		switch string(input) {
		case "0":
			p("设置生成目录中...")

		case "1":
			p("正在生成markdown文档...")
		case "2":
			p("生成表结构实体成功...")
		case "3":
			p("生成CURD增删改查成功...")
		case "4":
			p("设置结构体的映射名称成功...")
		case "5", "e", "q", "exit": //退出

			p("退出成功!")
			os.Exit(0)
		default:
			p("命令输入有错误!!!")
		}
	}
}
func main() {
	app := cli.NewApp()
	app.Name = "binw"                 //项目名称
	app.Author = "huvip"              //作者名称
	app.Version = "1.0.0 2019.10.3"   //版本号
	app.Copyright = "@Copyright 2019" //版权保护
	app.Usage = "快速使用go命令行文档翻译"       //说明
	cli.HelpFlag = cli.BoolFlag{      //修改系统默认
		Name:  "help",
		Usage: `“ af”：“南非荷兰语”`,
	}
	cli.VersionFlag = cli.BoolFlag{ //修改系统默认
		Name:  "version, v",
		Usage: "显示版本号",
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "u", Value: "root", Usage: "数据库用户名称"},
		cli.StringFlag{Name: "p", Value: "", Usage: "数据库密码"},
		cli.StringFlag{Name: "hh", Value: "", Usage: "显示帮助文档"},
	}
	app.Action = func(c *cli.Context) error {
		if c.NumFlags() == 0 {
			cli.ShowAppHelp(c)
			return nil
		}
		//数据库地址
		host := c.String("hh")
		if strings.EqualFold(host, "") {
			p(`“ af”：“南非荷兰语”，
“ sq”：“阿尔巴尼亚语”，
“ ar”：“阿拉伯语”，
“ hy”：“亚美尼亚语”，
“ az”：“阿塞拜疆”，
“ eu”：“巴斯克语”，
“ be”：“白俄罗斯语”，
“ bn”：“孟加拉语”，
“ bs”：“波斯尼亚语”，
“ bg”：“保加利亚语”，
“ ca”：“加泰罗尼亚语”，
“ ceb”：“ Cebuano”，
“ ny”：“ Chichewa”，
“ zh-cn”：“简体中文”，
“ zh-tw”：“繁体中文”，
“ co”：“ Corsican”，
“ hr”：“克罗地亚”，
“ cs”：“捷克”，
“ da”：“丹麦语”，
“ nl”：“荷兰语”，
“ en”：“英语”，
“ eo”：“世界语”，
“ et”：“爱沙尼亚语”，
“ tl”：“菲律宾”，
“ fi”：“芬兰语”，
“ fr”：“法语”，
“ fy”：“弗里斯兰语”，
“ gl”：“加利西亚语”，
“ ka”：“乔治亚”，
“ de”：“德语”，
“ el”：“希腊语”，
“ gu”：“古吉拉特语”，
“ ht”：“ Haitian Creole”，
“ ha”：“豪萨”，
“ haw”：“ Hawaiian”，
“ iw”：“希伯来语”，
“ hi”：“印地语”，
“ hmn”：“苗族”，
“ hu”：“匈牙利人”，
“是”：“冰岛语”，
“ ig”：“ Igbo”，
“ id”：“印尼语”，
“ ga”：“爱尔兰”，
“ it”：“意大利语”，
“ ja”：“日语”，
“ jw”：“ Javanese”，
“ kn”：“ Kannada”，
“ kk”：“哈萨克”，
“ km”：“高棉”，
“ ko”：“韩文”，
“ ku”：“库尔德语（Kurmanji）”，
“ ky”：“吉尔吉斯”，
“ lo”：“老挝”，
“ la”：“拉丁”，
“ lv”：“拉脱维亚”，
“ lt”：“立陶宛语”，
“ lb”：“卢森堡”，
“ mk”：“马其顿语”，
“ mg”：“恶意”，
“ ms”：“马来语”，
“ ml”：“马拉雅拉姆语”，
“ mt”：“马耳他语”，
“ mi”：“毛利人”，
“先生”：“马拉地”，
“ mn”：“蒙古语”，
“ my”：“缅甸（缅甸）”，
“ ne”：“尼泊尔”，
“ no”：“挪威语”，
“ ps”：“ Pashto”，
“ fa”：“波斯语”，
“ pl”：“波兰语”，
“ pt”：“葡萄牙语”，
“ ma”：“旁遮普语”，
“ ro”：“罗马尼亚语”，
“ ru”：“俄语”，
“ sm”：“萨摩亚”，
“ gd”：“苏格兰盖尔语”，
“ sr”：“塞尔维亚”，
“ st”：“ Sesotho”，
“ sn”：“ Shona”，
“ sd”：“ Sindhi”，
“ si”：“僧伽罗”，
“ sk”：“斯洛伐克”，
“ sl”：“斯洛文尼亚语”，
“ so”：“索马里”，
“ es”：“西班牙语”，
“ su”：“苏丹语”，
“ sw”：“斯瓦希里语”，
“ sv”：“瑞典语”，
“ tg”：“塔吉克”，
“ ta”：“泰米尔语”，
“ te”：“泰卢固语”，
“ th”：“泰国”，
“ tr”：“土耳其语”，
“ uk”：“乌克兰语”，
“ ur”：“ Urdu”，
“ uz”：“乌兹别克”，
“ vi”：“越南语”，
“ cy”：“威尔士语”，
“ xh”：“ Xhosa”，
“ yi”：“ Yiddish”，
“ yo”：“约鲁巴语”，
“zu”：“祖鲁”`)
		}
		//端口号
		port := c.Int("P")
		if port == 0 {
			port = 3306
		}
		//数据库用户名称
		user := c.String("u")
		if user == "" {
			user = "root"
		}
		//数据库密码
		pass := c.String("p")
		if pass == "" {
			color.Blue("输入密码:")
		}
		//编码格式
		charset := c.String("c")
		if charset == "" {
			charset = "utf8mb4"
		}
		//数据库名称
		dbname := c.String("d")
		if dbname == "" {
			return cli.NewExitError("数据库名称为空, 请使用 -d dbname", 9)
		}
		/*
			p("数据库端口号", port)
			p("数据库用户名称", user)
			p("数据库密码", pass)
			p("编码格式", charset)
			p("数据库地址", host)
			p("数据库名称", dbname)
		*/
		if dbname == "dbname" {
			color.Green("数据库登入成功\n")

			scant()
		} else {
			color.Green("数据库登入失败\n")

		}

		return nil
	}

	var err error
	err = app.Run(os.Args)
	if err != nil {
		//	p(err)
	}
}
