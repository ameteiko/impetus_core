# Strip-out the additional make positional arguments.
args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`
argss = `arg="$(word 2, $(MAKECMDGOALS))" && echo $${arg}`

list:
	@ls .data | cut -d " " -f 1

create:
	# check if exists
	@echo Creating instance [$(args)]
	@echo Creating instances [$(argss)]

start_db: create
	@echo not implemented

mock:
	@go generate ./src/...

# A placeholder to make possible commands like:
#   make create sandbox
%:
	@: