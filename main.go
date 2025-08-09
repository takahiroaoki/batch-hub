package main

import (
	"fmt"

	"github.com/takahiroaoki/batch-hub/cmd"
	"github.com/takahiroaoki/batch-hub/util"
)

func main() {
	if err := cmd.NewBundle().Execute(); err != nil {
		util.FatalLog(fmt.Sprintf("Failed to execute the command. Error: %v", err))
	}
}
