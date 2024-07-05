DROP TABLE IF EXISTS teacher;

CREATE TABLE
    teacher (
        id bigint (20) unsigned NOT NULL AUTO_INCREMENT,
        name varchar(10) NOT NULL DEFAULT '' COMMENT '姓名',
        nickname varchar(30) NOT NULL DEFAULT '' COMMENT '昵称',
        introduce varchar(200) NOT NULL DEFAULT '' COMMENT '简介',
        avatar varchar(100) NOT NULL DEFAULT '' COMMENT '头像',
        domain varchar(10) NOT NULL DEFAULT '' COMMENT '擅长领域',
        area varchar(10) NOT NULL DEFAULT '' COMMENT '地区',
        remark varchar(10) NOT NULL DEFAULT '' COMMENT '备注',
        class_hour varchar(20) NOT NULL DEFAULT '' COMMENT '上课时间',
        phone varchar(15) NOT NULL DEFAULT '' COMMENT '手机号',
        created_at TIMESTAMP DEFAULT NULL,
        updated_at TIMESTAMP DEFAULT NULL,
        deleted_at TIMESTAMP DEFAULT NULL,
        PRIMARY KEY (id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS user;

CREATE TABLE
    user (
        id bigint (20) NOT NULL AUTO_INCREMENT,
        phone varchar(15) NOT NULL DEFAULT '' COMMENT '手机号',
        `openid` varchar(50) DEFAULT NULL,
        `nickname` varchar(60) DEFAULT NULL,
        `wx_profile` json DEFAULT NULL,
        created_at TIMESTAMP DEFAULT NULL,
        updated_at TIMESTAMP DEFAULT NULL,
        deleted_at TIMESTAMP DEFAULT NULL,
        PRIMARY KEY (id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS lesson;

CREATE TABLE
    lesson (
        id bigint (20) unsigned NOT NULL AUTO_INCREMENT,
        teacher_id bigint (20) unsigned NOT NULL,
        title varchar(10) NOT NULL DEFAULT '' COMMENT '标题',
        introduce varchar(200) NOT NULL DEFAULT '' COMMENT '简介',
        `start_at` datetime NOT NULL DEFAULT '' COMMENT '上课时间',
        `end_at` datetime NOT NULL DEFAULT '' COMMENT '上课时间',
        `lesson_status` tinyint (1) NOT NULL DEFAULT 0 COMMENT '状态 0 未开始 1 正在上课 2 已结束',
        `is_paid` tinyint (1) NOT NULL DEFAULT 0 COMMENT '类型 0 免费 1 收费',
        `price` decimal(8, 2) NOT NULL DEFAULT 0 COMMENT '价格',
        province varchar(10) NOT NULL DEFAULT '' COMMENT '省',
        city varchar(10) NOT NULL DEFAULT '' COMMENT '市',
        area varchar(10) NOT NULL DEFAULT '' COMMENT '区',
        address varchar(10) NOT NULL DEFAULT '' COMMENT '地址',
        created_at TIMESTAMP DEFAULT NULL,
        updated_at TIMESTAMP DEFAULT NULL,
        deleted_at TIMESTAMP DEFAULT NULL,
        PRIMARY KEY (id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS lesson_signup;

CREATE TABLE
    lesson_signup (
        id bigint (20) unsigned NOT NULL AUTO_INCREMENT,
        user_id bigint (20) unsigned NOT NULL,
        lesson_id bigint (20) unsigned NOT NULL,
        created_at TIMESTAMP DEFAULT NULL,
        updated_at TIMESTAMP DEFAULT NULL,
        deleted_at TIMESTAMP DEFAULT NULL,
        PRIMARY KEY (id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '课程报名';

DROP TABLE IF EXISTS lesson_signin;

CREATE TABLE
    `lesson_signin` (
        id bigint (20) unsigned NOT NULL AUTO_INCREMENT,
        user_id bigint (20) unsigned NOT NULL,
        lesson_id bigint (20) unsigned NOT NULL,
        created_at TIMESTAMP DEFAULT NULL,
        updated_at TIMESTAMP DEFAULT NULL,
        deleted_at TIMESTAMP DEFAULT NULL,
        PRIMARY KEY (id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '课程签到';

CREATE TABLE
    `lesson_comment` (
        id bigint (20) unsigned NOT NULL AUTO_INCREMENT,
        user_id bigint (20) unsigned NOT NULL,
        teacher_id bigint (20) unsigned NOT NULL,
        lesson_id bigint (20) unsigned NOT NULL,
        created_at TIMESTAMP DEFAULT NULL,
        updated_at TIMESTAMP DEFAULT NULL,
        deleted_at TIMESTAMP DEFAULT NULL,
        PRIMARY KEY (id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '课程评论';