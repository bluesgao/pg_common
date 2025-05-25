CREATE TABLE `merchant` (
    `id` bigint NOT NULL auto_increment,
    `merchant_id` varchar(64) DEFAULT '' NOT NULL comment '商家ID',
    `name` varchar(64) DEFAULT '' NOT NULL comment '商家名称',
    `merchant_type` int NOT NULL DEFAULT '0' COMMENT ' 商家类型 0自营 1加盟',
    `status` int NOT NULL DEFAULT '0' COMMENT ' 0初始化 1启用 2停用',
    `avatar` varchar(128) DEFAULT '' COMMENT '头像',
    `phone` varchar(32) DEFAULT '' COMMENT '负责人手机号',
    `email` varchar(64) DEFAULT '' COMMENT '负责人邮箱',
    `created_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE `uniq_merchant_name` (`name`)
) engine = InnoDB charset = utf8mb4 row_format = DYNAMIC comment '商户基本信息表';

CREATE TABLE `merchant_base_settings` (
    `id` bigint NOT NULL auto_increment,
    `merchant_id` varchar(64) DEFAULT '' NOT NULL comment '商家ID',
    `support_currency` json DEFAULT(json_array()) COMMENT '支持币种',
    `support_region` json DEFAULT(json_array()) COMMENT '支持地域',
    `support_language` json DEFAULT(json_array()) COMMENT '支持语言',
    `default_language` varchar(32) DEFAULT '' COMMENT '默认语言',
    `default_currency` varchar(32) DEFAULT '' comment '默认币种',
    `default_region` varchar(32) DEFAULT '' comment '默认地域',
    `created_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE `uniq_merchant_id` (`merchant_id`)
) engine = InnoDB charset = utf8mb4 row_format = DYNAMIC comment '商家基础配置息表';

CREATE TABLE `merchant_platform_settings` (
    `id` bigint NOT NULL auto_increment,
    `merchant_id` varchar(64) DEFAULT '' NOT NULL comment '商家ID',
    `platform_type` int NOT NULL DEFAULT '0' COMMENT '平台类型 类型 0-web 1-h5 2-app',
    `app_id` bigint NOT NULL DEFAULT '0' COMMENT 'APPID',
    `app_secret` varchar(64) NOT NULL DEFAULT '0' COMMENT 'APPK秘钥',
    `status` int NOT NULL DEFAULT '0' COMMENT ' 0初始化 1启用 2停用',
    `created_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE `uniq_merchant_platform_type_app_id` (
        `merchant_id`,
        `platform_type`,
        `app_id`
    )
) engine = InnoDB charset = utf8mb4 row_format = DYNAMIC comment '商家平台配置表';

CREATE TABLE `merchant_domain_settings` (
    `id` bigint NOT NULL auto_increment,
    `merchant_id` varchar(64) DEFAULT '' NOT NULL comment '商家ID',
    `host` varchar(128) NOT NULL DEFAULT 'localhost' COMMENT '访问域名',
    `port` int NOT NULL DEFAULT '80' COMMENT '端口',
    `user_type` int NOT NULL DEFAULT '0' COMMENT '用户类型 类型 0会员 1代理 2商家',
    `platform_type` int NOT NULL DEFAULT '0' COMMENT '平台类型 类型 0-web 1-h5 2-app',
    `priority` int NOT NULL DEFAULT '0' COMMENT '优先级（0-100，优先级越大，使用几率越大）',
    `status` int NOT NULL DEFAULT '0' COMMENT ' 0初始化 1启用 2停用',
    `created_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE `uniq_host_port` (`host`, `port`)
) engine = InnoDB charset = utf8mb4 row_format = DYNAMIC comment '商家域名配置表';

CREATE TABLE `merchant_security_settings` (
    `id` bigint NOT NULL auto_increment,
    `merchant_id` varchar(64) DEFAULT '' NOT NULL comment '商家ID',
    `user_type` int NOT NULL DEFAULT '0' COMMENT '用户类型 0-系统用户，1-会员',
    `max_login_attempt` int NOT NULL DEFAULT '3' COMMENT '最大登录尝试次数(用于限制密码错误登录尝试次数)',
    `lock_duration` int NOT NULL DEFAULT '600' COMMENT '账户锁定时长（秒）锁定后需要等待多久才能再次尝试登录',
    `password_min_length` int NOT NULL DEFAULT '8' COMMENT '密码最小长度要求',
    `password_complexity` int NOT NULL DEFAULT '0' COMMENT '密码强度（低：无特殊字符，中：包含字母和数字，高：包含特殊字符、大小写字母、数字等）',
    `enable_captcha` int NOT NULL DEFAULT '0' COMMENT '是否启用验证码（0：不启用，1：启用）',
    `captcha_activation_threshold` int NOT NULL DEFAULT '0' COMMENT '验证码触发阈值，登录失败的次数达到该限制时，启动验证码',
    `enable_two_factor` int NOT NULL DEFAULT '0' COMMENT '是否启用双因素认证（0：不启用，1：启用）',
    `status` int NOT NULL DEFAULT '0' COMMENT ' 0初始化 1启用 2停用',
    `created_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE `uniq_merchant_user_type` (`merchant_id`, `user_type`)
) engine = InnoDB charset = utf8mb4 row_format = DYNAMIC comment '商家安全配置表';

CREATE TABLE `merchant_access_settings` (
    `id` bigint NOT NULL auto_increment,
    `merchant_id` varchar(64) DEFAULT '' NOT NULL comment '商家ID',
    `user_type` int NOT NULL DEFAULT '0' COMMENT '用户类型 0-系统用户，1-会员',
    `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户id',
    `member_groups` json DEFAULT(json_array()) COMMENT '会员组列表',
    `member_tags` json DEFAULT(json_array()) COMMENT '会员标签列表',
    `ip_address` varchar(64) DEFAULT NULL COMMENT '限制的IP地址',
    `device_type` varchar(64) DEFAULT NULL COMMENT '设备类型，限制某种设备的访问',
    `platform_type` int NOT NULL DEFAULT '0' COMMENT '限制平台类型 类型 0-web 1-h5 2-app',
    `restriction_type` int NOT NULL DEFAULT '0' COMMENT '限制类型 0-访问限制，1-访问阻止',
    `reason` varchar(128) DEFAULT NULL COMMENT '限制的原因',
    `start_time` datetime(3) DEFAULT NULL COMMENT '限制开始时间',
    `end_time` datetime(3) DEFAULT NULL COMMENT '限制结束时间',
    `status` int NOT NULL DEFAULT '0' COMMENT ' 0初始化 1启用 2停用',
    `created_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`)
) engine = InnoDB charset = utf8mb4 row_format = DYNAMIC comment '商家访问控制表';

CREATE TABLE `merchant_product_settings` (
    `id` bigint NOT NULL auto_increment,
    `merchant_id` varchar(64) DEFAULT '' NOT NULL comment '商家ID',
    `game_category_id` bigint NOT NULL DEFAULT '0' COMMENT '游戏分类id',
    `support_games` json DEFAULT(json_array()) COMMENT '支持游戏列表(详细信息)',
    `default_game` varchar(32) DEFAULT '' COMMENT '默认游戏',
    `status` int NOT NULL DEFAULT '0' COMMENT ' 0初始化 1启用 2停用',
    `created_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE `uniq_merchant_category_id` (
        `merchant_id`,
        `game_category_id`
    )
) engine = InnoDB charset = utf8mb4 row_format = DYNAMIC comment '商家游戏配置息表';