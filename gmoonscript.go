package gmoonscript

import (
	lua "github.com/yuin/gopher-lua"
)

// Compile compiles Moonscript code into Lua code using a Gopher-Lua
// VM. It is recommended that this method be called using an isolated
// Lua state, as it creates several globals that you may want to keep
// away from your program.
func Compile(state *lua.LState, moonscriptCode string) (string, error) {
	moonbundle, err := Asset("moon-bundle.lua")
	if err != nil {
		return "", err
	}

	state.SetGlobal("_moonbundle_code", lua.LString(moonbundle))
	state.SetGlobal("__moonscript_code", lua.LString(moonscriptCode))

	err = state.DoString(`
    package.loaded.moonscript = loadstring(_moonbundle_code)()

    local moonparse = require("moonscript.parse")
    local mooncompile = require("moonscript.compile")

    local tree, err = moonparse.string(__moonscript_code)
    if not tree then
      print("gmoonscript error: unable to parse moonscript, check formatting!")
    else
      __output_lua_code_, err = mooncompile.tree(tree)
    end

    -- remove all created modules and vars
    package.loaded.moonscript = nil
    moonparse = nil
    mooncompile = nil

    _moonbundle_code = nil
    __moonscript_code = nil
    collectgarbage()
  `)
	if err != nil {
		return "", err
	}

	luaOutput := state.GetGlobal("__output_lua_code_")
	state.SetGlobal("__output_lua_code", lua.LNil)

	return luaOutput.String(), nil
}

// Loader is the default stub used by the Gopher-Lua
// PreloadModule function to import a Go module.
func Loader(state *lua.LState) int {
	mod := state.SetFuncs(state.NewTable(), map[string]lua.LGFunction{
		"compile": func(L *lua.LState) int {
			code := L.CheckString(1)

			luaCode, err := Compile(L, code)
			if err != nil {
				state.Push(lua.LNil)
				state.Push(lua.LString(err.Error()))

				return 2
			}

			L.Push(lua.LString(luaCode))
			return 1
		},
	})

	// returns the module
	state.Push(mod)
	return 1
}
