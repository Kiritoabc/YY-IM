package domain

// User 用户信息
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UsersDB 用户数据库
var UsersDB = []User{
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
