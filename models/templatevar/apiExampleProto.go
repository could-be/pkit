package templatevar

const (
	ApiExampleProtoTemplate = `
syntax = "proto3";

// package 名字 service 名字 和 proto 文件名字,三者一致
package {{ToLower .ProjectName}};

// service 名字应该与 package 报名保持一致
// 使用 options go_package = "" 会导致问题
// 全局一个 service, 不支持多个 service 接口
service {{FirstToUpper .ProjectName}} {
    // eg...
    rpc Upper(UpperRequest) returns (UpperResponse) {}
}

// eg...
message UpperRequest{
    string str = 1;
}

message UpperResponse{
    string new_str = 1;
}

`
)
