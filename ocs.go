package ocs

import (
	"github.com/dustin/gomemcached/client"
	"gopkg.in/mgo.v2/bson"
)

const (
	memcachedProtocol = "tcp"
)

// Client 用来封装 memcached Client 对象（阿里云 OCS 兼容 memcached）。
type Client struct {
	c *memcached.Client
}

// Connect 创建一个 ocs 客户端连接。如果需要认证，则连接时会向服务器提供认证信息。
func Connect() (*Client, error) {
	c, err := memcached.Connect(memcachedProtocol, Server)
	if err != nil {
		return nil, err
	}

	if needAuth {
		_, err := c.Auth(Auth, Password)
		if err != nil {
			return nil, err
		}
	}

	return &Client{c}, nil
}

// Close 关闭 OCS 连接
func (c *Client) Close() {
	c.c.Close()
}

// Set 存储一个原始值（数据块）。
func (c *Client) Set(k string, v []byte) error {
	return c.SetE(k, v, 0)
}

// SetE 存储一个带有过期时间控制的原始值。
func (c *Client) SetE(k string, v []byte, exp int) error {
	_, err := c.c.Set(0, keyPrefix+k, 0, exp, v)
	return err
}

// SetObj 存储任何兼容 BSON Marshaler 的对象。
func (c *Client) SetObj(k string, v interface{}) error {
	return c.SetObjE(k, v, 0)
}

// SetObjE 存储任何兼容 BSON Marshaler 的对象，并带有超时控制。
func (c *Client) SetObjE(k string, v interface{}, exp int) error {
	b, err := bson.Marshal(v) // 把任意对象编码成 BSON 格式
	if err != nil {
		return err
	}

	// 把 resultBytes 存入 OCS
	return c.SetE(k, b, exp)
}

// Get 获取一个原始值。
func (c *Client) Get(k string) ([]byte, error) {
	resp, err := c.c.Get(0, keyPrefix+k)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

// GetObj 获取一个对象。
func (c *Client) GetObj(k string, v interface{}) error {
	b, err := c.Get(k)
	if err != nil {
		return err
	}
	return bson.Unmarshal(b, v)
}

// GetAll 获取全部给定的 key 所对应的原始值。
func (c *Client) GetAll(ks []string) (map[string][]byte, error) {
	for i, k := range ks {
		ks[i] = keyPrefix + k
	}
	resps, err := c.c.GetBulk(0, ks)
	if err != nil {
		return nil, err
	}

	result := make(map[string][]byte, len(resps))
	for k, resp := range resps {
		result[k] = resp.Body
	}
	return result, nil
}

// Del 删除一个键值对。
func (c *Client) Del(k string) error {
	_, err := c.c.Del(0, keyPrefix+k)
	return err
}
