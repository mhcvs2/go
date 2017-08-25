package test

import (
	"goCli/cmd/commands"
	"fmt"
	myLogger "goCli/logger"
	"goCli/config"
)

var CmdRun = &commands.Command{
	UsageLine: "test [name] [-main=...]",
	Short:     "test short",
	Long: `
test long........................

`,
	PreRun: func(cmd *commands.Command, args []string) { fmt.Println("I am pre run") },
	Run:    testRun,
}

var (
	main string

)

func init() {
	CmdRun.Flag.StringVar(&main, "main", "default_main","Specify main go files.")
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func testRun(cmd *commands.Command, args []string) int {
	//started <- true

	if len(args) > 0 {
		fmt.Println("arg[0]:", args[0])
	}
	fmt.Println("main:", main)
	//myLogger.Log.Fatal("fatal...............")   //will exit
	myLogger.Log.Errorf("error.............")
	myLogger.Log.Critical("Critical.............")
	myLogger.Log.Criticalf("Criticalf....")
	myLogger.Log.Errorf("Errorf....")
	myLogger.Log.Hint("hint...")
	myLogger.Log.Info("info")

	myLogger.Log.Infof("%d", config.Conf.Version)
	myLogger.Log.Debug("dsfsdffffffffffffffffffffff")


	return 0

}