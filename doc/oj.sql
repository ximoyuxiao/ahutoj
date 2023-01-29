#1.创建表空间
drop database ahutoj if EXISTS ahutoj;
create database ahutoj;
#2.创建用户
CREATE USER 'AHUTOnlinejudge'@'localhost' IDENTIFIED BY '2019ahut';
#3.授予用户表空间的权限
grant all privileges on ahutoj.* to 'AHUTOnlinejudge'@'%';
#4.创建表
use ahutoj

# 此处存储用户的基本信息
create table User(
    UID varchar(20)   primary key comment '用户ID',
    UserName  varchar(20) comment '用户名', 
    Pass   varchar(128) comment '密码',
    School varchar(128) comment '学校',
    Classes  varchar(30) comment '班级',
    Major  varchar(30) comment '专业',
    Adept  varchar(128) comment '擅长',
    Vjid   varchar(20) comment 'vj账号',
    Vjpwd  varchar(128) comment 'vj密码',
    CodeForceUser Text comment 'cf用户',
    Email varchar(20) comment '邮箱',
    HeadUrl Text comment '头像地址',
    Rating int comment '用户分数',
    LoginIP varchar(20) comment '最近登录IP',
    RegisterTime long comment '注册时间',
    Submited int(11) comment '提交次数',
    Solved int(11) comment 'AC次数',
    Defaulted varchar(3) comment '删除用户标志'
)DEFAULT CHARSET=utf8mb4;

create table Permission(
    UID varchar(20) comment '用户ID',
    SuperAdmin varchar(2)  comment '超级管理员' check (SuperAdmin   in ('N','Y')),
    ProblemAdmin varchar(2) comment '题目权限' check (ProblemAdmin in ('N','Y')) ,
    ListAdmin   varchar(2)  comment '题单权限' check (ListAdmin in('N','Y')),
    SourceAdmin varchar(2) comment '代码查看权限' check (SourceAdmin  in ('N','Y')) ,
    ContestAdmin varchar(2)  comment '竞赛权限'check (ContestAdmin in ('N','Y'))
)DEFAULT CHARSET=utf8mb4;

create table Problem(
    PID     varchar(40) primary key comment '题目ID',
    PType   varchar(40) comment '题目平台(LOCAL、本地|CODEFORCES、CF|ATCODER atcoder|LUOGU 洛谷)',
    Title   Text not null comment '标题' ,
    Description	 Text not null comment '描述',
    Input	 Text comment '输入',	
    Output	 Text comment '输出',	
    SampleInput	 Text comment '样例输入',	
    SampleOutput	 Text comment '样例输出',
    LimitTime	     int comment '极限时间',	
    LimitMemory	 int comment '极限内存',
    Hit 	Text comment '提示',
    Label            Text comment '标签',
    Origin int comment '是否外部题目',
    OriginPID Text comment '对应外部题目的ID',
    ContentType int comment '表示类型',    
    Accepted    int comment '总AC数量',
    Submited   int comment      '总提交数量',
    Visible  int comment '题目是否可见' ,
    SpjJudege varchar(3) comment '是否开启特判(N:不开启|Y:开启)',
    Source Text  comment '题目信息'
)DEFAULT CHARSET=utf8mb4;
ALTER TABLE Problem AUTO_INCREMENT = 1000;

create table List(
    LID int primary key AUTO_INCREMENT comment '题单ID',
    UID varchar(20) comment '创建用户',
    Description Text comment '题单描述',
    Title Text comment '题单标题',
    StartTime long comment '开始时间',
    Problems Text comment '题单题目序列',
    constraint fk_lst_UID FOREIGN KEY (UID)
    references User(UID) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;
ALTER TABLE List AUTO_INCREMENT = 1000;

create table ListProblem(
    LID int comment '题单ID',
    PID varchar(40) comment '题单题目',
    Title Text comment '题目标题',
    constraint pk_lpt primary key(LID,PID),
   
    constraint fk_lpt_PID FOREIGN KEY (PID)
    references Problem(PID) ON UPDATE CASCADE ON DELETE CASCADE,
    
    constraint fk_lpt_LID FOREIGN KEY (LID)
    references List(LID) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;

create table ListUser(
    LID int comment '题单ID',
    UID varchar(20) comment '用户ID',
    Submited int comment '提交数',
    Solved int comment 'AC数',
    constraint pk_lst primary key(UID,LID),
    
    constraint fk_lut_UID FOREIGN KEY (UID)
    references User(UID) ON UPDATE CASCADE ON DELETE CASCADE,
    
    constraint fk_lut_LID FOREIGN KEY (LID)
    references List(LID) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;
ALTER TABLE List AUTO_INCREMENT = 1000;

create table Contest(
    CID int primary key AUTO_INCREMENT comment '竞赛ID',
    UID varchar(20) comment '创建用户ID',
    Title Text comment '标题',
    Description Text comment '描述',
    BeginTime long comment '开始时间',
    EndTime long comment '结束时间',
    # 2 acm 1 oi	
    Type int comment '竞赛类型',  
    IsPublic int comment '是否公开',
    Problems Text comment '题目 + 顺序',
    Pass varchar(128) comment '竞赛密码',
    LangMask varchar(30) comment '语言掩码',
    Defaulted varchar(3) comment '是否可见 Y|N',
    constraint fk_ct_UID FOREIGN KEY (UID)
    references User(UID) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;
ALTER TABLE Contest AUTO_INCREMENT = 1000;

create table ConPro(
    CID int comment '竞赛ID',
    PID varchar(40) comment '题目ID',
    Title Text comment '题目标题',
    Submited  int comment '提交数',
    Solved int comment 'AC数',
    constraint pk_CPT primary key(CID,PID),

    constraint fk_cpt_CID FOREIGN KEY (CID)
    references Contest(CID) ON UPDATE CASCADE ON DELETE CASCADE,

    constraint fk_cpt_PID FOREIGN KEY (PID)
    references Problem(PID) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;

CREATE table Submit(
    SID int primary key AUTO_INCREMENT comment '提交结果',
    PID varchar(40) comment '题目ID',
    UID varchar(20) comment '提交用户ID',
    CID int comment '提交竞赛ID，-1表示为提交',
    JudgeID int	 comment '判题机ID',
    Source Text comment '提交代码',
    Lang int comment '提交语言',
    ResultACM varchar(30) comment 'ACM判题结果',
    PassSample int(11) comment '分数: 样例通过 / 样例总数 *  100',
    Sim   int(5) comment '相似度检测结果（0 -100）',
    UseTime long comment '使用时间',
    UseMemory long comment '使用内存',
    SubmitTime long comment '提交时间',
    # 这一块主要用于做缓存
    IsOriginJudge boolean comment '是否外部平台',
    OriginPID  Text comment '外部平台的PID',
    OJPlatform int comment '属于哪个平台',
    constraint fk_st_PIDs FOREIGN KEY (PID)
    references Problem(PID) ON UPDATE CASCADE ON DELETE CASCADE,

    constraint fk_st_UIDs FOREIGN KEY (UID)
    references User(UID) ON UPDATE CASCADE ON DELETE CASCADE
)DEFAULT CHARSET=utf8mb4;

create table CEINFO(
    SID int comment '提交ID',
    Info Text comment '错误内容'
)DEFAULT CHARSET=utf8mb4;

ALTER TABLE Submit AUTO_INCREMENT = 1000;
#5.添加数据
insert into User values('199094212','admin','',)
insert into Permission values('admin','Y','Y','Y','Y','Y');

#6.注册账户
use mysql;
update user set host='%' where user='AHUTOnlinejudge';
flush privileges;