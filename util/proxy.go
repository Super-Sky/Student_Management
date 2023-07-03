package util

//github api接口代理
func GetProxy(useFirst bool) string {
	s := []string{"http://WinBdhE:pqet87Nm@43.154.80.231:12128", "http://WinBdhE:pqet87Nm@43.134.210.13:12128"}
	if useFirst {
		return s[0]
	}
	return s[1]
}
