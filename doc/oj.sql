#1.创建表空间
create database ahutoj;
#2.创建用户
CREATE USER 'AHUTOnlinejudge'@'localhost' IDENTIFIED BY '2019ahut';
#3.授予用户表空间的权限
grant all privileges on ahutoj.* to 'AHUTOnlinejudge'@'localhost';
#4.创建表
use ahutoj
create table User(
    UID varchar(20)   primary key,
    UserName  varchar(20),
    Pass   varchar(128),
    School varchar(128),
    Classes  varchar(30),
    Major  varchar(30),
    Adept  varchar(128),
    Vjid   varchar(20),
    Vjpwd  varchar(128),
    Email varchar(20),
    HeadUrl Text,
)DEFAULT CHARSET=utf8mb4;

create table Permission(
    UID varchar(20),
    SuperAdmin varchar(2) check (SuperAdmin   in ('N','Y')),
    ProblemAdmin varchar(2) check (ProblemAdmin    in ('N','Y')),
    ListAdmin   varchar(2) check (ListAdmin in('N','Y')),
    SourceAdmin varchar(2) check (SourceAdmin  in ('N','Y')),
    ContestAdmin varchar(2) check (ContestAdmin in ('N','Y'))
)DEFAULT CHARSET=utf8mb4;

create table Problem(
    PID int primary key AUTO_INCREMENT,
    Title Text not null,
    Description	 Text not null,
    Input	 Text,	
    Output	 Text,	
    SampleInput	 Text,	
    SampleOutput	 Text,
    LimitTime	     int,	
    LimitMemory	 int,
    Hit 	Text,
    Label            Text
)DEFAULT CHARSET=utf8mb4;
ALTER TABLE Problem AUTO_INCREMENT = 1000;

create table List(
    LID int primary key AUTO_INCREMENT,
    UID varchar(20),
    Title Text,
    StartTime long,
    constraint fk_lst_UID FOREIGN KEY (UID)
    references User(UID) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;
ALTER TABLE List AUTO_INCREMENT = 1000;

create table ListProblem(
    LID int,
    PID int,
    Title Text,
    constraint pk_lpt primary key(LID,PID),
   
    constraint fk_lpt_PID FOREIGN KEY (PID)
    references Problem(PID) ON UPDATE CASCADE ON DELETE CASCADE,
    
    constraint fk_lpt_LID FOREIGN KEY (LID)
    references List(LID) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;

create table ListUser(
    LID int,
    UID varchar(20),
    SubmitNum int,
    ACNum int,
    constraint pk_lst primary key(UID,LID),
    
    constraint fk_lut_UID FOREIGN KEY (UID)
    references User(UID) ON UPDATE CASCADE ON DELETE CASCADE,
    
    constraint fk_lut_LID FOREIGN KEY (LID)
    references List(LID) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;
ALTER TABLE List AUTO_INCREMENT = 1000;

create table Contest(
    CID int primary key AUTO_INCREMENT,
    UID varchar(20),
    Title Text,
    Description Text,
    BeginTime long,
    EndTime long,
    # 2 acm 1 oi	
    Type int,  
    IsPublic int,
    Problems Text,
    Pass varchar(128),
    constraint fk_ct_UID FOREIGN KEY (UID)
    references User(UID) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;
ALTER TABLE Contest AUTO_INCREMENT = 1000;

create table ConPro(
    CID int,
    PID int,
    Title Text,
    SubmitNum int,
    ACNum int,
    constraint pk_CPT primary key(CID,PID),

    constraint fk_cpt_CID FOREIGN KEY (CID)
    references Contest(CID) ON UPDATE CASCADE ON DELETE CASCADE,

    constraint fk_cpt_PID FOREIGN KEY (PID)
    references Problem(PID) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;

CREATE table Submit(
    SID int primary key AUTO_INCREMENT,
    PID int,
    UID varchar(20),
    CID int,
    JudgeID int	,
    Source Text,
    Lang int,
    Result varchar(30),
    UseTime int,
    UseMemory int,
    SubmitTime long,
    constraint fk_st_PIDs FOREIGN KEY (PID)
    references Problem(PID) ON UPDATE CASCADE ON DELETE CASCADE,

    constraint fk_st_UIDs FOREIGN KEY (UID)
    references User(UID) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;

create table CEINFO(
    SID int,
    Info Text
)DEFAULT CHARSET=utf8mb4;

ALTER TABLE Submit AUTO_INCREMENT = 1000;
#5.添加数据
insert into User values('admin','墨羽','21de184f26d37d33d5581d923ae52c17','AHUT','软191','软件工程',null,null,'a2571717957@163.com','动态规划');
#此处对于密码199094212              
insert into Permission values('admin','Y','Y','Y','Y','Y');
insert into Problem values(null,'A+B问题','输入一个数字A和一个数字B要求输出A和B的和','分别输入两个整数A和B','输出A和B的和','1 2','3','1','128','','基础');
insert into Contest values(null,'admin','测试比赛1','用于测试',1639559000000,1639599000000,1,1,"1000;",null);
insert into Contest values(null,'admin','测试比赛2','用于测试',1639559000000,1639599000000,1,1,"",null);
insert into Contest values(null,'admin','测试比赛3','用于测试',1639559000000,1639599000000,1,1,"",null);
insert into List values(null,'admin','测试',1639599000000);
insert into ConPro values(1000,1000,'A+B问题',0,0);