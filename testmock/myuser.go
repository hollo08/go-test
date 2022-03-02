package testmock

func getUser(m MyFunc) string {
	user := m.GetInfo()
	return user
}
