shutdown(){
    pid=`pgrep $1`
    if ! [ -z "$pid" ]
    then
        kill -9 $pid
    fi
}
for exec in `gatway ahutoj originJudge persistence judged`;
do
    if test -x $i;then
        shutdown exec
    fi;
done