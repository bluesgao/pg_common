CREATE TABLE `user` (
    `id` bigint NOT NULL auto_increment,
    `merchant_id` varchar(64) DEFAULT '' NOT NULL comment '商家ID',
    `register_type` int NOT NULL DEFAULT '0' COMMENT ' 0账号注册 1邮箱注册 2手机号注册',
    `register_no` varchar(64) DEFAULT '' NOT NULL comment '用户注册号码（账号，邮箱，手机号）',
    `user_id` varchar(64) DEFAULT '' NOT NULL comment '用户ID',
    `agent_id` varchar(64) DEFAULT '' NOT NULL comment '直属代理ID',
    `user_type` int NOT NULL DEFAULT '0' COMMENT ' 0试玩 1正式',
    `status` int NOT NULL DEFAULT '0' COMMENT ' 0初始化 1启用 2异常 3停用',
    `login_pwd` varchar(256) DEFAULT '' COMMENT '登录密码',
    `pay_pwd` varchar(256) DEFAULT '' COMMENT '支付密码',
    `created_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE `uniq_user` (
        `merchant_id`,
        `register_type`,
        `register_no`
    )
) engine = InnoDB charset = utf8mb4 row_format = DYNAMIC comment '用户主表';

CREATE TABLE `user_info` (
    `id` bigint NOT NULL auto_increment,
    `merchant_id` varchar(64) DEFAULT '' NOT NULL comment '商家ID',
    `user_id` varchar(64) DEFAULT '' NOT NULL comment '用户ID',
    `name` varchar(64) DEFAULT '' NOT NULL comment '用户名称',
    `nickname` varchar(64) DEFAULT '' NOT NULL comment '用户昵称',
    `avatar` varchar(128) DEFAULT '' COMMENT '头像',
    `country` varchar(64) DEFAULT '' NOT NULL comment '国家',
    `region` varchar(64) DEFAULT '' NOT NULL comment '地区',
    `address` varchar(128) DEFAULT '' NOT NULL comment '地址',
    `birthday` DATE DEFAULT NULL comment '生日',
    `gender` varchar(32) NOT NULL DEFAULT 'Unknown' COMMENT ' Unknown未知 Male男 Female女',
    `created_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE `uniq_user` (`user_id`)
) engine = InnoDB charset = utf8mb4 row_format = DYNAMIC comment '用户信息表';