package domain

import "yy-im/domain/model"

// UsersDB 用户数据库
var UsersDB = []model.User{
	{
		ID:       1,
		Name:     "admin",
		Password: "admin",
	},
	{
		ID:       2,
		Name:     "张三",
		Password: "123456",
	},
	{
		ID:       3,
		Name:     "李四",
		Password: "123456",
	},
	{
		ID:       4,
		Name:     "王五",
		Password: "123456",
	},
	{
		ID:       5,
		Name:     "赵六",
		Password: "123456",
	},
	{
		ID:       6,
		Name:     "钱七",
		Password: "123456",
	},
	{
		ID:       7,
		Name:     "孙八",
		Password: "123456",
	},
}

// GetUserByName 根据用户名获取用户信息
func GetUserByName(username string) *model.User {
	for _, user := range UsersDB {
		if username == user.Name {
			tmp := model.User{
				ID:       user.ID,
				Name:     user.Name,
				Password: user.Password,
			}
			return &tmp
		}
	}
	return nil
}

// GetAllUser 获取所有用户信息
func GetAllUser() []model.User {
	temp := make([]model.User, len(UsersDB))
	// 防止外部修改，采用copy
	copy(temp, UsersDB)
	return temp
}
