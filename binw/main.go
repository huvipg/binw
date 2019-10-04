//
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/huvipg/binw/color"
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
	app.Name = "binw"               //项目名称
	app.Author = "huvip"                   //作者名称
	app.Version = "1.0.0 2019.10.3"     //版本号
	app.Copyright = "@Copyright 2019"   //版权保护
	app.Usage = "快速使用go命令行代码模版" //说明
	cli.HelpFlag = cli.BoolFlag{        //修改系统默认
		Name:  "help",
		Usage: "显示命令帮助",
	}
	cli.VersionFlag = cli.BoolFlag{ //修改系统默认
		Name:  "version, v",
		Usage: "显示版本号",
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "h", Value: "localhost", Usage: "数据库地址"},
		cli.IntFlag{Name: "P", Value: 3306, Usage: "端口号"},
		cli.StringFlag{Name: "u", Value: "root", Usage: "数据库用户名称"},
		cli.StringFlag{Name: "p", Value: "", Usage: "数据库密码"},
		cli.StringFlag{Name: "c", Value: "utf8mb4", Usage: "编码格式"},
		cli.StringFlag{Name: "d", Value: "", Usage: "*数据库名称"},
	}
	app.Action = func(c *cli.Context) error {
		if c.NumFlags() == 0 {
			cli.ShowAppHelp(c)
			return nil
		}
		//数据库地址
		host := c.String("h")
		if strings.EqualFold(host, "") {
			host = "localhost"
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
