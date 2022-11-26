package runner

import ( 

	"read/input"
        "read/output"
)


type Runner struct {
	Input input.Input

}

func New() *Runner {
	input := input.ReadArgs() // build in /input/flags.go
	return &Runner{Input: input}
}

func (r *runner) Execute() {
	var (
		commandProvided = false
	}


}






