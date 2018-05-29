package model

type User struct {
	ID       int    `xorm:"id"`
	Name     string `xorm:"name"`
	Password string `xorm:"password"`
}

func (u User) GetUser() *User {
	has, _ := engine.Get(&u)
	if has {
		return &u
	}
	return nil
}

func (u User) Insert() *User {
	_, err := engine.Insert(&u)
	if err == nil {
		return &u
	} else {
		return nil
	}
}
