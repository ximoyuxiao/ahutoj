shutdown(){
    pid=`pgrep $1`
    if ! [ -z "$pid" ]
    then
        kill -9 $pid
        echo 已结束进程$1,PID为$pid
    fi
}
for exec in gateway ahutoj originJudge persistence judged;
do
    if test -x $i;then
        shutdown $exec
    fi;
done