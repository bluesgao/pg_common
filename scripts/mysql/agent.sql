CREATE TABLE `agent` (
    `id` bigint NOT NULL auto_increment,
    `merchant_id` varchar(64) DEFAULT '' NOT NULL comment '商家ID',
    `agent_id` varchar(64) DEFAULT '' NOT NULL comment '代理ID',
    `name` varchar(64) DEFAULT '' NOT NULL comment '代理名称',
    `agent_type` int NOT NULL DEFAULT '0' COMMENT ' 0商家默认代理 1商家直属代理 2商家下级代理',
    `status` int NOT NULL DEFAULT '0' COMMENT ' 0初始化 1启用 2异常 3停用',
    `level` int NOT NULL DEFAULT '0' COMMENT '代理层级',
    `avatar` varchar(128) DEFAULT '' COMMENT '头像',
    `phone` varchar(32) DEFAULT '' COMMENT '手机号',
    `email` varchar(64) DEFAULT '' COMMENT '邮箱',
    `created_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE `uniq_merchant_agent_id` (`merchant_id`, `agent_id`)
) engine = InnoDB charset = utf8mb4 row_format = DYNAMIC comment '代理基本信息表';

CREATE TABLE `agent_hierarchy` (
    `id` bigint NOT NULL auto_increment,
    `merchant_id` varchar(64) DEFAULT '' NOT NULL comment '商家ID',
    `ancestor_id` varchar(64) DEFAULT '' NOT NULL comment '上级代理 ID',
    `descendant_id` varchar(64) DEFAULT '' NOT NULL comment '下级代理 ID',
    `depth` int DEFAULT 0 NOT NULL comment '层级相对深度（0 表示直接上下级）',
    `created_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE `uniq_merchant_ancestor_descendant_id` (
        `merchant_id`,
        `ancestor_id`,
        `descendant_id`
    )
) engine = InnoDB charset = utf8mb4 row_format = DYNAMIC comment '代理层级关系表';

CREATE TABLE `agent_dividend` (
    `id` bigint NOT NULL auto_increment,
    `merchant_id` varchar(64) DEFAULT '' NOT NULL comment '商家ID',
    `agent_id` varchar(64) DEFAULT '' NOT NULL comment '代理ID',
    `dividend_mode` int NOT NULL DEFAULT '0' COMMENT '代理分红模式（佣金，分成）',
    `dividend_rate` DECIMAL(20, 8) NOT NULL COMMENT '分红比例',
    `created_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE `uniq_merchant_agent_id` (`merchant_id`, `agent_id`)
) engine = InnoDB charset = utf8mb4 row_format = DYNAMIC comment '代理分红模式表';

CREATE TABLE `agent_settlement` (
    `id` bigint NOT NULL auto_increment,
    `merchant_id` varchar(64) DEFAULT '' NOT NULL comment '商家ID',
    `agent_id` varchar(64) DEFAULT '' NOT NULL comment '代理ID',
    `settlement_mode` int NOT NULL DEFAULT '0' COMMENT '结算模式(实时,每天,每周,每月)',
    `created_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE `uniq_merchant_agent_id` (
        `merchant_id`,
        `agent_id`,
        `settlement_mode`
    )
) engine = InnoDB charset = utf8mb4 row_format = DYNAMIC comment '代理结算模式表';

CREATE TABLE `agent_member_migration_record` (
    `id` bigint NOT NULL auto_increment,
    `merchant_id` varchar(64) DEFAULT '' NOT NULL comment '商家ID',
    `agent_id` varchar(64) DEFAULT '' NOT NULL comment '代理ID',
    `member_id` varchar(64) DEFAULT '' NOT NULL comment '会员ID',
    `target_agent_id` varchar(64) DEFAULT '' NOT NULL comment '目标代理ID',
    `remark` varchar(128) DEFAULT NULL COMMENT '备注',
    `created_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`)
) engine = InnoDB charset = utf8mb4 row_format = DYNAMIC comment '代理会员迁移记录表';