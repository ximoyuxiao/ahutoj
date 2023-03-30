#!/bin/bash
### BEGIN INIT INFO
# Provides:          ahutoj        //judged是自己创建的脚本名称
# Required-Start:    $local_fs $network $remote_fs $syslog $named
# Required-Stop:     $local_fs $network $remote_fs $syslog $named
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: starts the ahutoj daemon
# Description:       starts ahutoj using start-stop-daemon
### END INIT INFO

case "$1" in
ahutoj_server)
    DAEMON=/etc/bin/ahutoj_server
    PID=/run/ahutoj_server.pid
    ;;
judged)
    DAEMON=/etc/bin/judged
    PID=/run/judged.pid
    ;;

gatway)
    DAEMON=/etc/bin/gatway
    PID=/run/gatway.pid
    ;;

originjudged)
    DAEMON=/etc/bin/originjudged
    PID=/run/originjudged.pid
    ;;

*)
    echo "Usage: $NAME {ahutoj | judged | gatway | originjudged}" >&2
    exit 3
    ;;
esac

PATH=/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin
NAME=$1
DESC=$1
#DAEMON_OPTS=/etc/ahutoj/config.conf

test -x $DAEMON || exit 0 //-x检查文件名是否有执行权限

. /lib/init/vars.sh
. /lib/lsb/init-functions

start() {
    start-stop-daemon --start --quiet --make-pidfile \
        --pidfile $PID --exec $DAEMON >/dev/null ||
        return 1
    #	start-stop-daemon --start --quiet --make-pidfile \
    #         --pidfile $PID --exec $DAEMON -- $DAEMON_OPTS 2>/dev/null \
    #	|| return 2
}

#test_config() {
#	$DAEMON -t $DAEMON_OPTS >/dev/null 2>&1
#}

stop() {
    start-stop-daemon --stop --quiet --pidfile $PID \
        --exec $DAEMON
    RETVAL="$?"
    return "$RETVAL"
}

case "$2" in
start)
    log_daemon_msg "Starting $DESC" "$NAME"
    start
    echo "$(($(cat $PID) + 1))" >$PID
    case "$?" in
    0 | 1) log_end_msg 0 ;;
        #				2)   log_end_msg 1 ;;
    esac
    ;;
stop)
    log_daemon_msg "Stopping $DESC" "$NAME"
    stop
    case "$?" in
    0 | 1) log_end_msg 0 ;;
    2) log_end_msg 1 ;;
    esac
    ;;
restart)
    # Check configuration before stopping nginx
    #		if ! test_config; then
    #			log_end_msg 1 # Configuration error
    #			exit $?
    #		fi
    log_daemon_msg "Restarting $DESC" "$NAME"
    stop
    case "$?" in
    0 | 1)
        start
        case "$?" in
        0) log_end_msg 0 ;;
        1) log_end_msg 1 ;; # Old process is still running
        *) log_end_msg 1 ;; # Failed to start
        esac
        ;;
    *)
        # Failed to stop
        log_end_msg 1
        ;;
    esac
    ;;
*)
    echo "Usage: $NAME {start | stop | restart}" >&2
    exit 3
    ;;
esac

#case "$0" in
#        judged)
#        DAEMON=DAEMON_judged
#        NAME=NAME_judged
#        PID=PID_judged
#        ;;
#        ahutoj_server)
#        DAEMON=DAEMON_ahutoj_server
#        NAME=NAME_ahutoj_server
#        PID=PID_ahutoj_server
#        ;;
#esac
#ps -ef | grep judged | awk '{print $2}' | head -n1 > /run/judged.pid
#STOP_SCHEDULE="${STOP_SCHEDULE:-QUIT/5/TERM/5/KILL/5}"

#touch /run/judgeClient_test.pid
#ps -ef | grep judgeClient_test | awk '{print $2}' | head -n1 > /run/judged.pid
#ps -ef | grep -v grep | grep YOUR_PROCESS_NAME | awk '{ print $2 }' > pidfile

#upgrade_judged(){
#	if start-stop-daemon --stop --signal USR2 --quiet --pidfile \
#	$PID --name $NAME; then
#		# Wait for both old and new master to write their pid file
#		while [ ! -s "${PID}.oldbin" ] || [ ! -s "${PID}" ]; do
#			cnt=`expr $cnt + 1`
#			if [ $cnt -gt 10 ]; then
#				return 2
#			fi
#			sleep 1
#		done
#		# Everything is ready, gracefully stop the old master
#		if start-stop-daemon --stop --signal QUIT --quiet \
#		--pidfile "${PID}.oldbin" --name $NAME; then
#			return 0
#		else
#			return 3
#		fi
#	else
#		return 1
#	fi
#}
#upgrade)
#		log_daemon_msg "Upgrading binary" "$NAME"
#		upgrade_judged
#		log_end_msg $?
#	;;
