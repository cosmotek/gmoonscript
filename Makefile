stat:
	# https://github.com/jteeuwen/go-bindata
	go-bindata -o moonbundle.go -pkg gmoonscript moon-bundle.lua

amalg:
	# https://github.com/siffiejoe/lua-amalg
	amalg -o ../moonscriptbundle.lua -d moonscript moonscript.init moonscript.base \
	moonscript.compile moonscript.parse moonscript.util moonscript.transform moonscript.dump \
	moonscript.types moonscript.data moonscript.line_tables moonscript.parse.env \
	moonscript.parse.literals moonscript.parse.util moonscript.transform.statement \
	moonscript.transform.transformer moonscript.transform.names moonscript.transform.statements \
	moonscript.transform.comprehension moonscript.transform.destructure moonscript.errors \
	moonscript.transform.class moonscript.transform.value moonscript.transform.accumulator \
	moonscript.compile.statement moonscript.compile.value
