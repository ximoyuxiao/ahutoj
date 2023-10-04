root_dir=`pwd`
bin_dir=$root_dir/tmp/bin
core_dir=$root_dir/core
config_dir=$root_dir/config
cd $bin_dir
shutdown(){
    pid=`pgrep $1`
    if ! [ -z "$pid" ]
    then
        kill -9 $pid
        echo 已结束进程$1,PID为$pid
    fi
    rm -rf $1_dir
}
for exec in gatway ahutoj originJudge persistence judged oss forum;
do
    if test -x $i;then
        shutdown $exec
    fi;
done