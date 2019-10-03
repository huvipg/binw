binw 这个是一个控制台工具go代码生成工具

更新（2019年10月3日 13:20:50）
安装
go get -u -v github.com/huvip/binw/...
文档
DOC
演示实例使用说明
参数说明
-h value       数据库地址 (default: "localhost")
-P value       端口号 (default: 3306)
-u value       数据库用户名称 (default: "root")
-p value       数据库密码 (default: "123456")
-c value       编码格式 (default: "utf8mb4")
-d value       *数据库名称
--help         显示命令帮助
--version, -v  显示版本号
使用说明
输入界面上不同的命令进行操作即可
binw  -p root -d dbname
