package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

const luaFile = "./lua/test.lua"

func main() {
	l := lua.NewState()
	defer l.Close()

	if err := l.DoFile(luaFile); err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter: ")
		text, _ := reader.ReadString('\n')
		s := strings.TrimSpace(text)
		switch s {
		case "exit", "quit":
			return
		case "reload": // Update code in lua file and execute `reload`
			if err := l.DoFile(luaFile); err != nil {
				panic(err)
			}

		default: // Enter a number and call lua function `GetNumber` by this number
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err)
				break
			}
			if err := l.CallByParam(lua.P{
				Fn:      l.GetGlobal("GetNumber"),
				NRet:    1,
				Protect: true,
			}, lua.LNumber(n)); err != nil {
				fmt.Println(err)
				break
			}
			ret := l.Get(-1)
			l.Pop(-1)
			if res, ok := ret.(lua.LNumber); ok {
				fmt.Printf("GetNumber: %d\n", int(res))
			} else {
				fmt.Println("unexpected result")
			}
		}
	}
}
