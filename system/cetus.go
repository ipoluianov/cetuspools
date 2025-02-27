package system

func (c *System) CetusGetName() string {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	return c.Name
}
