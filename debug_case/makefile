DIRAPP	=addr  dlv  endian  esacape  gc retval inline createobj quote 
MAKE 	= make

all:app

app:
	for i in $(DIRAPP); do\
		(cd $$i && $(MAKE)) || exit 1;\
	done
