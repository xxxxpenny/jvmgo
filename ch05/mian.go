package main

import (
	"jvmgo/ch05/classpath"
	"strings"
	"jvmgo/ch05/rtda/heap"
	"fmt"
)

func main() {
	cmd := parseCmd()
	startJVM(cmd)

}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classLoader := heap.NewClassLoader(cp, cmd.p)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	} else {
		interpret(mainMethod, cmd.p, cmd.args)
	}

}
