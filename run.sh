cd ./tmp/bin

run(){
    if [ ! -d $1 ]; then \
		mkdir -p $1; \
	fi
    cd $1
    nohup ../$1 ../../config/config.yaml &
    cd ..
}

for exec in `gatway ahutoj originJudge persistence`;
do
    if test -x $i;then
        run exec
    fi;
done
cd core
./jugded