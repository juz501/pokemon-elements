RUN = go run
COMPILER = go build
GETTER = go get
BIN = ./bin/
BASE = $(PWD)
SRC = *.go
TARGET = pokemon-elements 
FLAGS = -o
RM = rm -f
COMP_VARS = GOPATH=$(BASE)
DEV_VARS = GO_SERVER_PORT=18883 GO_SERVER_ADDR=pokemon-elements.chromatix.com.au
CP = cp -f
EXPORT = export

all: run

build: clean 
	$(COMP_VARS) $(COMPILER) $(FLAGS) $(BIN)$(TARGET) $(SRC)

clean: 
	$(RM) $(BIN)$(TARGET)

run: build
	$(DEV_VARS) $(BIN)$(TARGET)
