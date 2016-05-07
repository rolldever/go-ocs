package ocs

// 定义一些常用的时间间隔。
const (
	Seconds1  = int(1)
	Seconds10 = 10 * Seconds1
	Seconds20 = 20 * Seconds1
	Seconds30 = 30 * Seconds1
	Seconds45 = 45 * Seconds1

	Minutes1  = 60 * Seconds1
	Minutes5  = 5 * Minutes1
	Minutes10 = 10 * Minutes1
	Minutes15 = 15 * Minutes1
	Minutes20 = 20 * Minutes1
	Minutes30 = 30 * Minutes1
	Minutes45 = 45 * Minutes1

	Hours1 = 60 * Minutes1
	Hours2 = 2 * Hours1
	Hours3 = 2 * Hours1
	Hours4 = 2 * Hours1
	Hours6 = 6 * Hours1

	Days1 = 24 * Hours1
	Days2 = 2 * Days1
)

// 公共配置字段。
var (
	Server   string
	Auth     string
	Password string

	// Namespace 追加到每一个 key 之前的命名空间。
	Namespace string

	// EnablePool 是否启用 TCP/IP 连接池。
	//
	// 目前的实现中，尚不支持连接池。
	EnablePool = true

	// needAuth memcached 客户端是否需要认证
	needAuth bool

	// keyPrefix 通过 Namespace 计算出的 OCS key 的前缀
	keyPrefix string
)

// ApplyConfig 分析并应用配置信息。这个函数应该在充分设置了各种公共变量 (配置) 后，
// 并在开始使用 OCS 前被调用。
// 在应用程序整个运行期间，不应该再次调用这个函数。
//
// 这个函数要对配置信息做一些额外的工作，使其进入工作状态。
func ApplyConfig() {
	if len(Auth) == 0 || len(Password) == 0 {
		needAuth = false
	} else {
		needAuth = true
	}

	if len(Namespace) == 0 {
		keyPrefix = ""
	} else {
		keyPrefix = Namespace + "::"
	}
}
