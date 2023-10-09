root_dir=`pwd`
bin_dir=$root_dir/tmp/bin
core_dir=$root_dir/core
config_dir=$root_dir/config
cd $bin_dir

run(){
    if [ ! -d $1_dir ]; then \
		mkdir -p $1_dir; \
	fi
    cd $1_dir
    nohup $bin_dir/$1 $config_dir/config.yaml &
    cd ..
}

for exec in gatway ahutoj originJudge persistence oss forum;
do
    if test -x $i;then
        run $exec
    fi;
done
cd $core_dir
./judged
cd $root_dir