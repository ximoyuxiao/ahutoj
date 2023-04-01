CONFIGPATH=./config
TARGETPATH=./tmp
TARGETCONFIGPATH=${TARGETPATH}/config
TARGETBINPATH=${TARGETPATH}/bin
COREPATH=./core/
SERVICETARGET=./web/service/
INSTALLPATH=/usr/bin/ahutoj/
BUILDBINS=ahutoj originJudge originproblem persistence useranalytics gatway

all:init build

init:
	go mod tidy

judged:
	cd ${COREPATH} && make

$(foreach BIN,$(BUILDBINS),$(eval $(BIN): $(BIN).go))

${TARGETBINPATH}:
	if [ ! -d ${TARGETBINPATH} ]; then \
		mkdir -p ${TARGETBINPATH}; \
	fi

%.go:
	go build ${SERVICETARGET}$*/$@
	mv $* ${TARGETBINPATH}/$*

build:${TARGETBINPATH} ${BUILDBINS} judged
	cp -r ${CONFIGPATH} ${TARGETCONFIGPATH}

install:
	if [ ! -d ${INSTALLPATH} ]; then \
		mkdir -p ${INSTALLPATH}; \
	fi
	
	cp -r ${TARGETBINPATH}/* ${INSTALLPATH}
	cp -r ${CONFIGPATH}/* ${INSTALLPATH}
clean:
	cd ${COREPATH} && make clean
	rm -rf ${TARGETPATH}