# GO

MAKE=make

WORKDIR=/home/wangjian/my_go/src/github.com/wangjian890523
DEBUG_CASE=$(WORKDIR)/go-case/debug_case


all: debug 

debug:
	cd $(DEBUG_CASE) && $(MAKE)

clean:
