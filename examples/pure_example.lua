-- this script is an example of moonscript compilation in a
-- pure-Lua program that serves to explain the methods broken
-- out into the gmoonscript module.

local moonscript_code = [[
class Thing
  name: "unknown"

class Person extends Thing
  say_name: => print "Hello, I am #{@name}!"

with Person!
  .name = "MoonScript"
  \say_name!
]]

local moonscript = require("moon-bundle")

local parser = require("moonscript.parse")
local compiler = require("moonscript.compile")

tree, err = parser.string(moonscript_code)
assert(tree)

lua_code, err = compiler.tree(tree)
assert(lua_code)

loadstring(lua_code)()
