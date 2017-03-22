.PHONY: clean

OUT = salat
PKG = github.com/aldidana/go-salat
VERSION = 1.0.0
BIN_DIR = /usr/local/bin/

# target #

default: clean build

build: 
	go build -i -v -o ${OUT} -ldflags="-X main.version=${VERSION}" ${PKG}; \
	mv ${OUT} $(BIN_DIR)

clean:
	rm -rf ${OUT}