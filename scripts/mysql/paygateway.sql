
CREATE TABLE `pay_channel` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `channel_id` varchar(64) NOT NULL COMMENT '支付渠道ID',
    `name` varchar(64) NOT NULL COMMENT '支付渠道名称',
    `avatar` varchar(128) DEFAULT '' COMMENT '支付渠道头像',
    `support_region` json DEFAULT(json_array()) COMMENT '支持地域',
    `support_currency` json DEFAULT(json_array()) COMMENT '支持币种',
    `support_pay_mode` json DEFAULT(json_array()) COMMENT '支持支付方式(储蓄卡，信用卡，收款码，TRC32)',
    `status` int DEFAULT '0' COMMENT '支付渠道状态 0-初始化，1-激活，2-停用',
    `phone` varchar(32) DEFAULT '' COMMENT '支付渠道手机号',
    `email` varchar(64) DEFAULT '' COMMENT '支付渠道邮箱',
    `remark` varchar(256) NOT NULL COMMENT '备注',
    `create_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `update_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT  CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    unique `uniq_channel_id` (`channel_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '支付渠道表';

CREATE TABLE `pay_channel_params` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `channel_id` varchar(64) NOT NULL COMMENT '支付渠道ID',
    `mch_id` varchar(64) NOT NULL COMMENT '支付渠道mch_id',
    `app_id` varchar(64) NOT NULL COMMENT '支付渠道app_id',
    `app_secret_key` varchar(128) NOT NULL COMMENT '支付渠道app_secret',
    `api_public_key` varchar(2048) NOT NULL COMMENT '支付渠道api证书公钥',
    `api_private_key` varchar(2048) NOT NULL COMMENT '支付渠道api证书私钥',
    `pay_url` varchar(512) NOT NULL COMMENT '支付链接',
    `callback_url` varchar(512) NOT NULL COMMENT '支付回调链接(path参数中附带商户id和渠道id和订单id)',
    `remark` varchar(512) NOT NULL COMMENT '备注',
    `create_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `update_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT  CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    unique `uniq_channel_app_id` (
        `channel_id`,
        `app_id`,
    )
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '支付渠道参数表';

CREATE TABLE `pay_channel_quota` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `channel_id` varchar(64) NOT NULL COMMENT '支付渠道ID',
    `biz_type` int default 0 not null comment '业务类型 0-充值，1-提款',
    `pay_mode` varchar(8) NOT NULL COMMENT '支付模式',
    `currency` varchar(8) NOT NULL COMMENT '币种',
    `quota_type` int DEFAULT '0' COMMENT '支付限制类型 0-不限制，1-金额，2-笔数',
    `quota_cycle` int default 0 not null comment '限制周期（0-不限制，1-每天，2-每月，3-每年）',
    `amount_limit` decimal(20, 8) NOT NULL COMMENT '限制周期内的限额',
    `fee_free_amount_limit` decimal(20, 8) NOT NULL COMMENT '限制周期内免手续费限额',
    `used_amount` decimal(20, 8) NOT NULL COMMENT '限制周期内已用额度',
    `count_limit` decimal(20, 8) NOT NULL COMMENT '限制周期内的限额笔数',
    `fee_free_count_limit` decimal(20, 8) NOT NULL COMMENT '限制周期内的免手续费笔数',
    `used_count` decimal(20, 8) NOT NULL COMMENT '限制周期内已用笔数',
    `latest_used_at` datetime DEFAULT NULL COMMENT '最后占用额度时间',
    `per_amount_limit` decimal(20, 8) NOT NULL COMMENT '单笔限额',
    `remark` varchar(128) NOT NULL COMMENT '备注',
    `status` int DEFAULT '0' COMMENT '渠道限额状态 0-初始化，1-激活，2-停用',
    `create_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `update_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT  CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    unique `uniq_channel_biz_type_currency` (
        `channel_id`,
        `biz_type`,
        `currency`
    )
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '支付渠道配额表';

CREATE TABLE `pay_channel_fee` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `biz_type` int default 0 not null comment '业务类型 0-充值，1-提款',
    `channel_id` varchar(64) NOT NULL COMMENT '支付渠道ID',
    `currency` varchar(8) NOT NULL COMMENT '币种',
    `fee_mode` varchar(8) NOT NULL COMMENT '手续费收取模式：0-每笔固定 fixed，1-每笔百分比 percent，2-阶梯费率 tierd',
    `fixed_fee` decimal(20, 8) NOT NULL COMMENT '每笔固定',
    `percent_fee_rate` decimal(20, 8) NOT NULL COMMENT '每笔百分比',
    `remark` varchar(128) NOT NULL COMMENT '备注',
    `status` int DEFAULT '0' COMMENT '状态 0-初始化，1-激活，2-停用',
    `create_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `update_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT  CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    unique `uniq_channel_biz_type_currency_fee_mode` (
        `channel_id`,
        `biz_type`,
        `currency`,
        `fee_mode`
    )
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '支付渠道费率表';

CREATE TABLE `pay_channel_fee_tier` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `biz_type` int default 0 not null comment '业务类型 1充值，2提款',
    `channel_id` varchar(64) NOT NULL COMMENT '支付渠道ID',
    `currency` varchar(8) NOT NULL COMMENT '币种',
    `tier_mode` varchar(8) NOT NULL COMMENT '分层模式 单笔，总额',
    `start_amount` decimal(20, 8) NOT NULL COMMENT '起始金额（闭区间）',
    `end_amount` decimal(20, 8) NOT NULL COMMENT '结束金额（开区间）',
    `fee_rate` decimal(20, 8) NOT NULL COMMENT '费率',
    `remark` varchar(256) NOT NULL COMMENT '备注',
    `create_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `update_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT  CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    unique `uniq_channel_biz_type_currency` (
        `channel_id`,
        `biz_type`,
        `currency`
    )
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '支付渠道费率分层表';

-- CREATE TABLE `pay_merchant_service_fee` (
--     `id` bigint NOT NULL AUTO_INCREMENT,
--     `merchant_id` bigint default 0 not null comment '商户ID',
--     `biz_type` int default 0 not null comment '业务类型 1充值，2提款',
--     `currency` varchar(8) NOT NULL COMMENT '币种',
--     `fee_mode` varchar(8) NOT NULL COMMENT '费率模式 fixed固定费率，Tierd分层费率',
--     `fee_rate` decimal(20, 8) NOT NULL COMMENT '服务费率',
--     `remark` varchar(128) NOT NULL COMMENT '备注',
--     `status` int DEFAULT '0' COMMENT '状态 0-初始化，1-激活，2-暂停，3-停用',
--     `create_by` varchar(64) DEFAULT NULL COMMENT '创建者',
--     `update_by` varchar(64) DEFAULT NULL COMMENT '更新者',
--     `created_at` datetime DEFAULT NULL COMMENT '创建时间',
--     `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
--     `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
--     PRIMARY KEY (`id`),
--     unique `uniq_merchant_channel_pay_mode_currency` (
--         `merchant_id`,
--         `channel_id`,
--         `pay_mode`,
--         `currency`
--     )
-- ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '商户服务费费率表';

-- CREATE TABLE `pay_merchant_service_fee_tier` (
--     `id` bigint NOT NULL AUTO_INCREMENT,
--     `merchant_id` bigint default 0 not null comment '商户ID',
--     `biz_type` int default 0 not null comment '业务类型 1充值，2提款',
--     `currency` varchar(8) NOT NULL COMMENT '币种',
--     `tier_mode` varchar(8) NOT NULL COMMENT '分层模式 单笔，总额',
--     `start_amount` decimal(20, 8) NOT NULL COMMENT '起始金额（闭区间）',
--     `end_amount` decimal(20, 8) NOT NULL COMMENT '结束金额（开区间）',
--     `fee_rate` decimal(20, 8) NOT NULL COMMENT '服务费率',
--     `remark` varchar(128) NOT NULL COMMENT '备注',
--     `create_by` varchar(64) DEFAULT NULL COMMENT '创建者',
--     `update_by` varchar(64) DEFAULT NULL COMMENT '更新者',
--     `created_at` datetime DEFAULT NULL COMMENT '创建时间',
--     `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
--     `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
--     PRIMARY KEY (`id`),
--     unique `uniq_merchant_channel_pay_mode_currency` (
--         `merchant_id`,
--         `channel_id`,
--         `pay_mode`,
--         `currency`
--     )
-- ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '商户服务费率分层表';


CREATE TABLE `pay_channel_order` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `channel_id` varchar(64) NOT NULL COMMENT '支付渠道ID',
    `merchant_id` varchar(64) not null comment '商户ID',
    `order_id` varchar(64) NOT NULL COMMENT '订单ID',
    `tx_order_id` varchar(64) DEFAULT NULL COMMENT '交易订单号',
    `tx_type` int DEFAULT '0' COMMENT '交易类型 1-充值，2-提款，3-退款',
    `tx_amount` DECIMAL(20, 8) NOT NULL COMMENT '交易金额',
    `tx_fee_amount` DECIMAL(20, 8) NOT NULL COMMENT '交易手续费',
    `tx_currency` varchar(8) NOT NULL COMMENT '交易币种',
    `tx_time` datetime DEFAULT NULL COMMENT '交易时间',
    `tx_detail` varchar(1024) NOT NULL COMMENT '交易详情',
    `status` int DEFAULT '0' COMMENT '状态 0-待处理，1-成功，2-失败',
    `third_party_settlement_amount` DECIMAL(20, 8) NOT NULL COMMENT '支付结算金额(第三方支付渠道)',
    `third_party_fee_amount` DECIMAL(20, 8) NOT NULL COMMENT '支付手续费(第三方支付渠道)',
    `third_party_order_id` varchar(64) DEFAULT NULL COMMENT '第三方支付渠道订单号',
    `pay_url` varchar(1024) DEFAULT NULL COMMENT '渠道支付链接',
    `pay_request` varchar(1024) DEFAULT NULL COMMENT '支付渠道请求参数',
    `pay_response` varchar(1024) DEFAULT NULL COMMENT '支付渠道返回结果',
    `callback_time` datetime DEFAULT NULL COMMENT '回调时间',
    `callback_url` varchar(1024) NOT NULL COMMENT '回调url',
    `callback_response` varchar(1024) DEFAULT NULL COMMENT '支付渠道回调结果',
    `remark` varchar(256) NOT NULL COMMENT '备注',
    `create_by` varchar(64) DEFAULT NULL COMMENT '创建者',
    `update_by` varchar(64) DEFAULT NULL COMMENT '更新者',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT  CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '最后更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    unique `uniq_merchant_channel_order_id` (
        `merchant_id`,
        `channel_id`,
        `order_id`
    )
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 comment '支付渠道订单表';
