package ocs

// Set 存储一个原始值。
func Set(k string, v []byte) error {
	c, err := Connect()
	if err != nil {
		return err
	}
	defer c.Close()
	return c.Set(k, v)
}

// SetE 存储一个原始值，并带有超时控制。
func SetE(k string, v []byte, exp int) error {
	c, err := Connect()
	if err != nil {
		return err
	}
	defer c.Close()
	return c.SetE(k, v, exp)
}

// SetObj 存储一个兼容于 BSON marshaler 的对象。
func SetObj(k string, v interface{}) error {
	c, err := Connect()
	if err != nil {
		return err
	}
	defer c.Close()
	return c.SetObj(k, v)
}

// SetObjE 存储一个兼容于 BSON marshaler 的对象，并提供超时控制。
func SetObjE(k string, v interface{}, exp int) error {
	c, err := Connect()
	if err != nil {
		return err
	}
	defer c.Close()
	return c.SetObjE(k, v, exp)
}

// Get 获取一个原始值。
func Get(k string) ([]byte, error) {
	c, err := Connect()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.Get(k)
}

// GetObj 获取一个对象。
func GetObj(k string, v interface{}) error {
	c, err := Connect()
	if err != nil {
		return err
	}
	defer c.Close()
	return c.GetObj(k, v)
}

// GetAll 获取 ks (Keys) 对应的所有原始值。
func GetAll(ks []string) (map[string][]byte, error) {
	c, err := Connect()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.GetAll(ks)
}

// Del 删除一个键值对。
func Del(k string) error {
	c, err := Connect()
	if err != nil {
		return err
	}
	defer c.Close()
	return c.Del(k)
}
