GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

EXEC=gc


all: clean 

build:
	$(GOBUILD)
	@echo "$@ $(EXEC) completed!"
test:
	$(GOTEST) -v
	@echo "$@ $(EXEC) completed!"

clean:
	$(GOCLEAN)
	@echo "$@ $(EXEC) completed!"

run:
	$(GOBUILD)
    
