package command

type Command struct {
	Delay  int
	Mode   string `json:"mode"`
	CMD    string `json:"cli"`
	End    string
	Result string
}
