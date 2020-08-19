package param

// usually use struct in models package as param but sometimes param is complex and need other struct
// some param has its own verify method and trans method, if there is no param package and as usual
// param struct would be in router package, and other packages can not use these struct because of cycle import
type Pager struct {
	Page   int    `json:"page"`
	LastID string `json:"last_id"`
	Limit  int    `json:"limit"`
}
