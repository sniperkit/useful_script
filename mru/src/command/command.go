package command

type Command struct {
	Delay  int
	Mode   string
	CMD    string `json:"cli"`
	End    string
	Result string
}
