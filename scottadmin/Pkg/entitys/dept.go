package entitys

type Dept struct {
	Deptno int `json:"deptno"`
	Dname  string `json:"dname"`
	Loc    string `json:"loc"`
}

type DeptList []Dept

type Error struct {
	ValidationMsg []string
	Flg          bool
}

type Flash struct {
	FlashMsg string
	Flg          bool
}
