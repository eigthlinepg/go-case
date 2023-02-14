GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

EXEC=createobj


all: clean 

build:
	$(GOBUILD)
	@echo "$@ $(EXEC) completed!"
test:
	$(GOTEST) -v
	@echo "$@ $(EXEC) completed!"

clean:
	$(GOCLEAN)
	@rm -f $(EXEC).o
	@echo "$@ $(EXEC) completed!"

run:
	$(GOBUILD)
    
