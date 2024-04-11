create database restful;

use restful;

create table account (
    id int primary key auto_increment comment '主键',
    nickname varchar(20) comment '昵称',
    email varchar(50) comment '邮箱',
    deleted tinyint(1) default 0 comment '是否删除',
    created_at datetime comment '创建时间',
    modified_at datetime comment '修改时间'
) comment '账户表';

insert into account (id, nickname, email, deleted, created_at, modified_at) values (null, '张三', 'test@gmail.com', false, now(), now());
