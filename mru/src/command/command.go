package command

type Command struct {
	Delay  int
	Mode   string
	CMD    string `json:"Command"`
	End    string
	Result string
}
