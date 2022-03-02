package testmock

func GetUser(m MyFunc) string {
	user := m.GetInfo()
	return user
}
