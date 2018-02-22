# gmoonscript
> Moonscript compiler for the Gopher Lua VM

*Very experimental... use with caution.*
This module provides methods for compiling Moonscript code to Lua from Go or Lua (Gopher-Lua) programs. Additionally this repository contains an amalgamation bundle of the Moonscript Lua library (modified to use only pure-Lua modules) that can be used in Lua or LuaJIT programs without any external dependencies (see examples/pure_example.lua for usage).

See the [godoc](https://godoc.org/github.com/rucuriousyet/gmoonscript) for documentation...

## Installation
`go get -v -u github.com/rucuriousyet/gmoonscript`

## Usage

Example usage from Gopher-Lua:
```go
package main

import (
	"github.com/rucuriousyet/gmoonscript"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	state := lua.NewState()
	state.PreloadModule("moonc", gmoonscript.Loader)

	err := state.DoString(`
local moonscript_code = [[
class Thing
  name: "unknown"

class Person extends Thing
  say_name: => print "Hello, I am #{@name}!"

with Person!
  .name = "MoonScript"
  \say_name!
]]

local moonc = require("moonc")

lua_code, err = moonc.compile(moonscript_code)
if err ~= nil then
	print(err)
else
	loadstring(lua_code)()
end
	`)
	if err != nil {
		panic(err)
	}
}
```

Example usage from Go
```go
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
```
