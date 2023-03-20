INSTALL_DIR = /usr/bin
SOURCE_DIR  = ./
SERVERNAME = ahutoj_server
BINDIR = ./tmp/bin

ahutoj_server:a_s_build a_s_install

a_s_build:
	go build -o ${BINDIR}/${SERVERNAME} ./web/main.go
a_s_install:build
	cp ${BINDIR}/${SERVERNAME} ${INSTALL_DIR}/${SERVERNAME}