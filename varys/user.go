package varys

type UserInfo struct {
	Name string
}

// GetInfoByUID 获取用户uid
func GetInfoByUID(int64) (*UserInfo, error) {
	return &UserInfo{Name: "RandySun"}, nil
}
