package main

import lua "github.com/yuin/gopher-lua"
import "github.com/rucuriousyet/gmoonscript"

func main() {
	state := lua.NewState()
	moonCode := `
class Thing
  name: "unknown"

class Person extends Thing
  say_name: => print "Hello, I am #{@name}!"

with Person!
  .name = "MoonScript"
  \say_name!
  `

	luaCode, err := gmoonscript.Compile(state, moonCode)
	if err != nil {
		panic(err)
	}

	err = state.DoString(luaCode)
	if err != nil {
		panic(err)
	}
}
