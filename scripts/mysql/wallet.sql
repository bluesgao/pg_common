CREATE TABLE `payee` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `owner_id` varchar(64) NOT NULL COMMENT '卡号所属者ID',
    `owner_type` int NOT NULL COMMENT '卡号所属者类型 1-平台 2-商家 3-代理 4-会员',
    `payee_type` int NOT NULL COMMENT '收款账号类型 0-银行卡 1-USDT',
    `payee_name` varchar(64) DEFAULT NULL COMMENT '收款人名称',
    `payee_no` varchar(128) DEFAULT NULL COMMENT '收款人账号',
    `payee_org_name` varchar(128) DEFAULT NULL COMMENT '机构名称（银行）',
    `payee_org_no` varchar(128) DEFAULT NULL COMMENT '机构代码（银行代码）',
    `create_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `update_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_payee` (
        `owner_id`,
        `owner_type`,
        `payee_type`,
        `payee_no`
    )
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '收款账号表';

CREATE TABLE `wallet` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `owner_id` varchar(64) NOT NULL COMMENT '钱包所属者ID',
    `owner_type` int NOT NULL COMMENT '钱包所属者类型 1-平台 2-商家 3-代理 4-会员',
    `wallet_id` varchar(64) NOT NULL COMMENT '钱包ID',
    `name` varchar(64) DEFAULT NULL COMMENT '钱包名称',
    `wallet_type` int NOT NULL COMMENT '钱包类型 1储蓄(不可透支) 2信用(可透支) 3手续费 4奖励 5冻结 6充值中间账户 7提现中间账户',
    `status` int NOT NULL DEFAULT 0 COMMENT '钱包状态 0-初始化  1-激活 2-异常 3-关闭',
    `currency` varchar(8) NOT NULL COMMENT '币种',
    `balance` DECIMAL(38, 18) NOT NULL DEFAULT 0 COMMENT '可用余额',
    `frozen_balance` DECIMAL(38, 18) NOT NULL DEFAULT 0 COMMENT '冻结余额',
    `version` bigint NOT NULL DEFAULT 0 COMMENT '版本',
    `create_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `update_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_wallet` (
        `owner_id`,
        `owner_type`,
        `wallet_type`,
        `currency`
    )
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '钱包表';

CREATE TABLE `wallet_record` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `record_id` varchar(64) DEFAULT NULL COMMENT '明细号',
    `owner_id` varchar(64) NOT NULL COMMENT '钱包所属者ID',
    `owner_type` int NOT NULL COMMENT '钱包所属者类型 1-平台 2-商家 3-代理 4-会员',
    `wallet_type` int NOT NULL COMMENT '钱包类型 1储蓄(不可透支) 2信用(可透支) 3手续费 4奖励 5冻结 6充值中间账户 7提现中间账户',
    `wallet_id` varchar(64) NOT NULL COMMENT '钱包ID',
    `wallet_version` bigint NOT NULL DEFAULT 0 COMMENT '钱包版本',
    `currency` varchar(8) NOT NULL COMMENT '币种',
    `amount` DECIMAL(38, 18) NOT NULL DEFAULT 0 COMMENT '交易金额',
    `pre_balance` DECIMAL(38, 18) NOT NULL DEFAULT 0 COMMENT '交易发生前余额',
    `post_balance` DECIMAL(38, 18) NOT NULL DEFAULT 0 COMMENT '交易发生后余额',
    `tx_type` int NOT NULL DEFAULT 0 COMMENT '交易类型 0-未知 1-充值 2-提款 3-转账 4-冻结 5-解冻 6-退款 7-调账',
    `tx_order_id` varchar(64) NOT NULL COMMENT '交易单号',
    `tx_time` datetime NOT NULL COMMENT '交易时间',
    `tx_detail` varchar(1024) NOT NULL COMMENT '交易详情',
    `remark` varchar(64) DEFAULT NULL COMMENT '备注',
    `created_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_record` (`record_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '钱包明细表';