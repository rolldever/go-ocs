// Package ocs is a simple implementation of OCS client.
//
// 其中提供编码/解码任何 Golang 类型数据，包括自定义类型数据的方法
// GetObj, SetObj/SetObjE
//
// 编码/解码对象时选择使用 BSON 而 Marshal 和 Unmarshal。
// 见 Gist
//
//     https://gist.github.com/everbeen/07c38dcbba8d8a9d3c6f#file-benchmar
//
// 提供了 Gob vs. BSON vs. MessagePack 三种编码/解码方式的效率对比，
// 并从中选择了编码和解码速度更均衡和更高效的 BSON 方式。
//
// 对于 MessagePack，其编码速度稍优于 BSON，但解码速度低于 BSON，
// 而且没有提供 Golang 内置类型的直接支持。
//
// ocs 并不访问包外的任何对象，故其拥有独立的配置空间。源代码见 vars.go。
// 在使用 ocs 前应该配置 (设置) 这些公共变量。
//
// ocs 中也提供了一步式访问函数，当不需要维持 OCS 连接时可以使用这些一步式函数
// 操作 OCS 中的数据。源代码见 one_step.go
package ocs
