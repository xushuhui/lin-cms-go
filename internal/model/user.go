package model

type User struct {
	Id            uint   `gorm:"column:id;type:int unsigned" json:"id"`                                       // 账号id
	Email         string `gorm:"column:email;type:varchar(30);default:''" json:"email"`                       // 邮箱
	Phone         string `gorm:"column:phone;type:varchar(15);default:''" json:"phone"`                       // 手机号
	Password      string `gorm:"column:password;type:char(32);default:''" json:"password"`                    // 密码
	CreateAt      int    `gorm:"column:create_at;type:int;default:'0'" json:"create_at"`                      // 创建时间
	CreateIpAt    string `gorm:"column:create_ip_at;type:varchar(12);default:''" json:"create_ip_at"`         // 创建ip
	LastLoginAt   int    `gorm:"column:last_login_at;type:int;default:'0'" json:"last_login_at"`              // 最后一次登录时间
	LastLoginIpAt string `gorm:"column:last_login_ip_at;type:varchar(12);default:''" json:"last_login_ip_at"` // 最后一次登录ip
	LoginTimes    int    `gorm:"column:login_times;type:int;default:'0'" json:"login_times"`                  // 登录次数
	Status        int8   `gorm:"column:status;type:tinyint(1);default:'0'" json:"status"`                     // 状态 1:enable, 0:disable, -1:deleted

}

func GetAccountUserOne(where string, args ...interface{}) (model User, err error) {
	err = DB.First(&model, where, args).Error
	return

}
