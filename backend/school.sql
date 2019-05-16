create database school;

create table Department(
    id varchar(10) not null,
    name varchar(20) ,
    address varchar(50),
    phone varchar(20),
    primary key(id)
);

create table Student(
    id varchar(10) not null,
    name varchar(20) ,
    gender varchar(10),
    birthday date,
    birthplace varchar(50),
    phone varchar(20),
    did varchar(10),
    primary key(id)
    foreign key (did) references Department(did)
);


create table Course(
    cid varchar(20) not null,
    cname varchar(20),
    credit varchar(20),
    did varchar(20),
    primary key(cid)
    foreign key (did) references Department(did)
);

create table Teacher(
    tid varchar(20) not null,
    did varchar(20),
    tname varchar(20),
    gender varchar(10),
    birthday date,
    education varchar(20),
    wage int,
    primary key(tid),
    foreign key (departmentId) references Department(id)
);

create table CourseCalendar(
    tid varchar(20) not null,
    cid varchar(20) not null,
    term varchar(20) not null,
    classTime varchar(20) not null,
    primary key(tid,cid,term,classTime),
    foreign key (cid) references Course(cid),
    foreign key (tid) references Teacher(tid)
);

create table CourseSelect(
    id varchar(20) not null,
    cid varchar(20) not null,
    tid varchar(20) not null,
    term varchar(20) not null,
    classTime varchar(20) not null,
    primary key(id,cid,tid,term)
);
    foreign key (studentId) references Student(id),
    foreign key (teacherId) references Teacher(id)
    foreign key (courseId) references Course(id),

create table User(
    password varchar(20) not null,
    identity varchar(20) not null,
    id varchar(20) not null,
    primary key(id)
);