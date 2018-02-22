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
