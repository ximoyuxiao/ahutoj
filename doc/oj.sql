#1.创建表空间
create database ahutoj;
#2.创建用户
CREATE USER 'AHUTOnlinejudge'@'localhost' IDENTIFIED BY '2019ahut';
#3.授予用户表空间的权限
grant all privileges on ahutoj.* to 'AHUTOnlinejudge'@'localhost';
#4.创建表
use ahutoj
create table User(
    uid varchar(20)   primary key,
    uname  varchar(20),
    pass   varchar(128),
    school varchar(128),
    classes  varchar(30),
    major  varchar(30),
    vjid   varchar(20),
    vjpwd  varchar(128),
    email varchar(20)
)DEFAULT CHARSET=utf8mb4;
create table Permission(
    uid varchar(20),
    administrator varchar(2) check (administrator   in ('N','Y')),
    problem_edit varchar(2) check (problem_edit    in ('N','Y')),
    source_browser varchar(2) check (source_browser  in ('N','Y')),
    contest_creator varchar(2) check (contest_creator in ('N','Y'))
)DEFAULT CHARSET=utf8mb4;

create table Problem(
    pid int primary key AUTO_INCREMENT,
    title Text not null,
    Description	 Text not null,
    input	 Text,	
    output	 Text,	
    sample_input	 Text,	
    sample_output	 Text,
    limtTime	     int,	
    limithitMemory	     int,
    hit 	Text
)DEFAULT CHARSET=utf8mb4;
ALTER TABLE Problem AUTO_INCREMENT = 1000;
create table List(
    lid int primary key AUTO_INCREMENT,
    uid varchar(20),
    title Text,
    stime datetime,
    constraint fk_lst_uid FOREIGN KEY (uid)
    references User(uid) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;
ALTER TABLE List AUTO_INCREMENT = 1000;

create table ListProblem(
    lid int,
    pid int,
    constraint pk_lpt primary key(lid,pid),
   
    constraint fk_lpt_pid FOREIGN KEY (pid)
    references Problem(pid) ON UPDATE CASCADE ON DELETE CASCADE,
    
    constraint fk_lpt_lid FOREIGN KEY (lid)
    references List(lid) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;

create table ListUser(
    lid int,
    uid varchar(20),
    submit_num int,
    ac_num int,
    constraint pk_lst primary key(uid,lid),
    
    constraint fk_lut_uid FOREIGN KEY (uid)
    references User(uid) ON UPDATE CASCADE ON DELETE CASCADE,
    
    constraint fk_lut_lid FOREIGN KEY (lid)
    references List(lid) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;
ALTER TABLE List AUTO_INCREMENT = 1000;

create table Contest(
    cid int primary key AUTO_INCREMENT,
    uid varchar(20),
    title Text,
    description Text,
    begin_time datetime,
    end_time datetime,	
    ctype varchar(10) check (ctype in('ACM','OI')),
    ispublic varchar(10) check (ispublic in('private','public')),
    pass varchar(128),
    constraint fk_ct_uid FOREIGN KEY (uid)
    references User(uid) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;
ALTER TABLE Contest AUTO_INCREMENT = 1000;

create table ConPro(
    cid int,
    pid int,
    submit_num int,
    ac_num int,
    constraint pk_CPT primary key(cid,pid),

    constraint fk_cpt_cid FOREIGN KEY (cid)
    references Contest(cid) ON UPDATE CASCADE ON DELETE CASCADE,

    constraint fk_cpt_pid FOREIGN KEY (pid)
    references Problem(pid) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;



CREATE table Submit(
    sid int primary key AUTO_INCREMENT,
    pid int,
    uid varchar(20),
    cid int,
    judgeid int	,
    source Text,
    lang varchar(10) check ( lang in('C++11','JAVA','Python3','C99')),
    result varchar(10),
    usetime int,
    memory int,
    submitTime datetime,
    constraint fk_st_pids FOREIGN KEY (pid)
    references Problem(pid) ON UPDATE CASCADE ON DELETE CASCADE,

    constraint fk_st_uids FOREIGN KEY (uid)
    references User(uid) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;
ALTER TABLE Submit AUTO_INCREMENT = 1000;
#5.添加数据
insert into User values('admin','墨羽','21de184f26d37d33d5581d923ae52c17','AHUT','软191','软件工程',null,null,'a2571717957@163.com');
#此处对于密码199094212              
insert into Permission values('admin','Y','Y','Y','Y');
insert into Problem values(null,'A+B问题','输入一个数字A和一个数字B要求输出A和B的和','分别输入两个整数A和B','输出A和B的和','1 2','3','1','128','');
insert into Contest values(null,'admin','测试比赛1','用于测试','2021-12-15 16:32:22','2021-12-16 16:32:22','ACM','public',null);
insert into Contest values(null,'admin','测试比赛2','用于测试','2021-12-15 16:32:22','2021-12-16 20:32:22','ACM','public',null);
insert into Contest values(null,'admin','测试比赛3','用于测试','2021-12-15 16:32:22','2021-12-16 22:32:22','ACM','public',null);
insert into List values(null,'admin','测试','2021-12-15 16:32:22');
