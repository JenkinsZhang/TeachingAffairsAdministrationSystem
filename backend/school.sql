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

create table Term(
    term varchar(20) not null,
    isCurrent int not null,
    primary key(term)
);

create table Other (
    name  varchar(20) not null primary key,
    msg varchar(20) not null
);

create table Log(
    type varchar(20),
    table varchar(20),
    time date
);


DELIMITER $
create trigger log_trigger_insert_department after insert 
on Department
for each row
begin
insert into Log values("insert", "department", now());
end $

DELIMITER $
create trigger log_trigger_insert_course after insert 
on Course
for each row
begin
insert into Log values("insert", "course", now());
end $

DELIMITER $
create trigger log_trigger_insert_teacher after insert 
on Teacher
for each row
begin
insert into Log values("insert", "teacher", now());
end $

DELIMITER $
create trigger log_trigger_insert_student after insert 
on Student
for each row
begin
insert into Log values("insert", "student", now());
end $

DELIMITER $
create trigger log_trigger_insert_courseCalendar after insert 
on CourseCalendar
for each row
begin
insert into Log values("insert", "courseCalendar", now());
end $


DELIMITER $
create trigger log_trigger_insert_courseSchedule after insert 
on CourseSchedule
for each row
begin
insert into Log values("insert", "courseSchedule", now());
end $

DELIMITER $
create trigger log_trigger_insert_user after insert 
on User
for each row
begin
insert into Log values("insert", "user", now());
end $

DELIMITER $
create trigger log_trigger_insert_term after insert 
on Term
for each row
begin
insert into Log values("insert", "term", now());
end $

DELIMITER $
create trigger log_trigger_insert_other after insert 
on Other
for each row
begin
insert into Log values("insert", "other", now());
end $
