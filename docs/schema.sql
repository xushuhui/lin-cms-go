SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- 文件表
-- ----------------------------
DROP TABLE IF EXISTS lin_file;
CREATE TABLE lin_file
(
   id        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    path        varchar(500)     NOT NULL,
    type       tinyint(1)      NOT NULL DEFAULT 1 COMMENT '1 LOCAL 本地，2 REMOTE 远程',
    name        varchar(100)     NOT NULL,
    extension   varchar(50)               DEFAULT NULL,
    size        int(11)                   DEFAULT NULL,
    md5         varchar(40)               DEFAULT NULL COMMENT 'md5值，防止上传重复文件',
    created_at TIMESTAMP      DEFAULT NULL,
    updated_at TIMESTAMP       DEFAULT NULL,
    deleted_at TIMESTAMP               DEFAULT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY md5_del (md5, deleted_at)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;
ALTER TABLE lin_file MODIFY created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE lin_file MODIFY updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP  ON UPDATE CURRENT_TIMESTAMP;
-- ----------------------------
-- 日志表
-- ----------------------------
DROP TABLE IF EXISTS lin_log;
CREATE TABLE lin_log
(
   id        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    message     varchar(450)              DEFAULT NULL,
    user_id     int(10) unsigned NOT NULL,
    username    varchar(24)               DEFAULT NULL,
    status_code bigint(20)                 DEFAULT NULL,
    method      varchar(20)               DEFAULT NULL,
    path        varchar(50)               DEFAULT NULL,
    permission  varchar(100)              DEFAULT NULL,
    created_at TIMESTAMP      DEFAULT NULL,
    updated_at TIMESTAMP       DEFAULT NULL,
    deleted_at TIMESTAMP               DEFAULT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;
ALTER TABLE lin_log MODIFY created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE lin_log MODIFY updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
-- ----------------------------
-- 权限表
-- ----------------------------
DROP TABLE IF EXISTS lin_permission;
CREATE TABLE lin_permission
(
   id        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    name        varchar(60)      NOT NULL COMMENT '权限名称，例如：访问首页',
    module      varchar(50)      NOT NULL COMMENT '权限所属模块，例如：人员管理',
    mount       tinyint(1)       NOT NULL DEFAULT 1 COMMENT '0：关闭 1：开启',
    created_at TIMESTAMP      DEFAULT NULL,
    updated_at TIMESTAMP       DEFAULT NULL,
    deleted_at TIMESTAMP               DEFAULT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;
ALTER TABLE lin_permission MODIFY created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE lin_permission MODIFY updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
-- ----------------------------
-- 分组表
-- ----------------------------
DROP TABLE IF EXISTS lin_group;
CREATE TABLE lin_group
(
   id        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    name        varchar(60)      NOT NULL COMMENT '分组名称，例如：搬砖者',
    info        varchar(255)              DEFAULT NULL COMMENT '分组信息：例如：搬砖的人',
    level       tinyint(2)       NOT NULL DEFAULT 3 COMMENT '分组级别 1：root 2：guest 3：user（root、guest分组只能存在一个)',
    created_at TIMESTAMP      DEFAULT NULL,
    updated_at TIMESTAMP       DEFAULT NULL,
    deleted_at TIMESTAMP               DEFAULT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY name_del (name, deleted_at)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;
ALTER TABLE lin_group MODIFY created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE lin_group MODIFY updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
-- ----------------------------
-- 分组-权限表
-- ----------------------------
DROP TABLE IF EXISTS lin_group_permission;
CREATE TABLE lin_group_permission
(
    id           bigint(20)unsigned NOT NULL AUTO_INCREMENT,
    group_id     bigint(20) unsigned NOT NULL COMMENT '分组id',
    permission_id bigint(20)unsigned NOT NULL COMMENT '权限id',
    PRIMARY KEY (id),
    KEY group_id_permission_id (group_id, permission_id) USING BTREE COMMENT '联合索引'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

-- ----------------------------
-- 用户基本信息表
-- ----------------------------
DROP TABLE IF EXISTS lin_user;
CREATE TABLE lin_user
(
   id        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    username    varchar(24)      NOT NULL COMMENT '用户名，唯一',
    nickname    varchar(24)               DEFAULT NULL COMMENT '用户昵称',
    avatar      varchar(500)              DEFAULT NULL COMMENT '头像url',
    phone       char(20)              DEFAULT NULL COMMENT '手机',
    created_at TIMESTAMP      DEFAULT NULL,
    updated_at TIMESTAMP       DEFAULT NULL,
    deleted_at TIMESTAMP               DEFAULT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY username_del (username, deleted_at),
    UNIQUE KEY email_del (email, deleted_at)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;
ALTER TABLE lin_user MODIFY created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE lin_user MODIFY updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
-- ----------------------------
-- ----------------------------
-- 用户授权信息表
# id
# user_id
# identity_type 登录类型（手机号 邮箱 用户名）或第三方应用名称（微信 微博等）
# identifier 标识（手机号 邮箱 用户名或第三方应用的唯一标识）
# credential 密码凭证（站内的保存密码，站外的不保存或保存token）
-- ----------------------------
DROP TABLE IF EXISTS lin_user_identiy;
CREATE TABLE lin_user_identiy
(
    id           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    user_id       bigint(20) unsigned NOT NULL COMMENT '用户id',
    identity_type varchar(100)     NOT NULL,
    identifier    varchar(100),
    credential    varchar(100),
    created_at   TIMESTAMP      DEFAULT NULL,
    updated_at   TIMESTAMP       DEFAULT NULL,
    deleted_at   TIMESTAMP               DEFAULT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;
ALTER TABLE lin_user_identiy MODIFY created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE lin_user_identiy MODIFY updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
DROP TABLE IF EXISTS book;
CREATE TABLE book
(
    id         bigint(20)  NOT NULL AUTO_INCREMENT,
    title       varchar(50) NOT NULL,
    author      varchar(30)          DEFAULT NULL,
    summary     varchar(1000)        DEFAULT NULL,
    image       varchar(100)         DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP  DEFAULT NULL,
    deleted_at TIMESTAMP          DEFAULT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;
ALTER TABLE book MODIFY created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE book MODIFY updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;

-- ----------------------------
-- 用户-分组表
-- ----------------------------
DROP TABLE IF EXISTS lin_user_group;
CREATE TABLE lin_user_group
(
    id      bigint(20)unsigned NOT NULL AUTO_INCREMENT,
    user_id  bigint(20) unsigned NOT NULL COMMENT '用户id',
    group_id bigint(20)unsigned NOT NULL COMMENT '分组id',
    PRIMARY KEY (id),
    KEY user_id_group_id (user_id, group_id) USING BTREE COMMENT '联合索引'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;

-- ----------------------------
-- 插入超级管理员
-- 插入root分组
-- ----------------------------
BEGIN;
INSERT INTO lin_user(id, username, nickname)
VALUES (1, 'root', 'root');

INSERT INTO lin_user_identiy (id, user_id, identity_type, identifier, credential)

VALUES (1, 1, 'USERNAME_PASSWORD', 'root',
        '$2a$10$Subt/mCLem8axyivSN4BBeSLUtdPPEhDJJTWpGQVWJst1aVa9TzQq');

INSERT INTO lin_group(id, name, info, level)
VALUES (1, 'root', '超级用户组', 1);

INSERT INTO lin_group(id, name, info, level)
VALUES (2, 'guest', '游客组', 2);

INSERT INTO lin_user_group(id, user_id, group_id)
VALUES (1, 1, 1);
INSERT INTO `lin_permission` VALUES (1, '查看lin的信息', '信息',1, '2020-04-23 09:11:16', '2020-04-23 09:11:16', NULL);
INSERT INTO `lin_permission` VALUES (2, '查询自己信息', '用户',1, '2020-04-23 09:11:16.531', '2020-04-23 09:11:16.531', NULL);
INSERT INTO `lin_permission` VALUES (3, '查询自己拥有的权限', '用户',1, '2020-04-23 09:11:16.544', '2020-04-23 09:11:16.544', NULL);
INSERT INTO `lin_permission` VALUES (4, '查询日志记录的用户', '日志',1, '2020-04-23 09:11:16.554', '2020-04-23 09:11:16.554', NULL);
INSERT INTO `lin_permission` VALUES (5, '删除图书', '图书',1, '2020-04-23 09:11:16.562', '2020-04-23 09:11:16.562', NULL);
INSERT INTO `lin_permission` VALUES (6, '查询所有日志', '日志',1, '2020-04-23 09:11:16.571', '2020-04-23 09:11:16.571', NULL);
INSERT INTO `lin_permission` VALUES (7, '测试日志记录', '信息',1, '2020-04-23 09:11:16.580', '2020-04-23 09:11:16.580', NULL);
INSERT INTO `lin_permission` VALUES (8, '搜索日志', '日志',1, '2020-04-23 09:11:16.590', '2020-04-23 09:11:16.590', NULL);

COMMIT;