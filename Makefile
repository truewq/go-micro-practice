# note: call scripts from /scripts

# usage:
#     make 
           #build all cmds
#     make cmd=xxx
           #build xxx cmd
all:
	./scripts/build.sh $(cmd)
