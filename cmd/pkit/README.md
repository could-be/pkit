## pkit

`一个根据 proto 文件生成微服务的工具`

## 安装, 安装到$`GOPATH/bin` 目录$

`make install`

## 生成微服务框架

`pkit init [project]`

## 更新微服务框架, 需要再当前工程目录下,可以指定也可以不指定proto 文件

`pkit update`
`pkit update path1/path2/path3/api/a.proto`

或者当前目录执行

`go generate -x ./...`

## 测试环境下可选环境变量

| ---               | --                               |                                              |
| ----------------- | -------------------------------- | -------------------------------------------- |
| 可选变量名         | 含义                               | 例子                                           |
| GIT               | util的git 仓库名, eg: github; gitlab 等    | GIT="github.com"                             |
| GIT_USER          | git 仓库账号                         | GIT_USER="xiaoming"_                         |
| GIT_LOCAL_FLAG    | 非空, go mod 使用 replace 替换本地包      | GIT_LOCAL="true"_                            |
| GIT_UTIL_PATH     | 配合 GIT_LOCAL替换远程仓库为本地包, 本地包的相对路径 | GIT_LOCAL_ADDRESS="/Users/admin/myPproject_" |

## 正式环境下

`GIT="github.com" GIT_USER="xiaoming" pkit update`
