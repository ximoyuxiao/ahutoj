package constanct

type VerfiyLevel uint8

// map所需要的的权限 第x为为1表示  如果用户权限为第X位则具备权限
const (
	UNLOGIN       VerfiyLevel = 0xff // 1111 1111
	CommomUser    VerfiyLevel = 0xfe // 1111 1110
	Administrator VerfiyLevel = 0xfc // 1111 1100
	ProblemAdmin  VerfiyLevel = 0x88 // 1000 1000
	ContestAdmin  VerfiyLevel = 0x90 // 1001 0000
	SourceBorwser VerfiyLevel = 0xa0 // 1010 0000
	ListAdmin     VerfiyLevel = 0xc0 // 1100 0000
	SuperAdmin    VerfiyLevel = 0x80 // 1000 0000
)
