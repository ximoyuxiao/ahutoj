package dao

type User struct {
	Uid     string `db:"uid"`
	Uname   string `db:"uname"`
	Pass    string `db:"pass"`
	School  string `db:"school"`
	Classes string `db:"classes"`
	Major   string `db:"major"`
	Vjid    string `db:"vjid"`
	Vjpwd   string `db:"vjpwd"`
	Email   string `db:"email"`
}
