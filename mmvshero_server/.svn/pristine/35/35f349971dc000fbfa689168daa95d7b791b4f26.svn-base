create table IF NOT exists `players`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT COMMENT "UID",
    `lv` int(11) not null DEFAULT 0,
    `stone` int(11) not null DEFAULT 0,
    `gold` int(11) not null DEFAULT 0,
    `freeGold` int(11) not null DEFAULT 0,
    `trophy` int(11) not null DEFAULT 0,
    `totalCharge` int(11) not null DEFAULT 0,
    `addedUpTime` int(11) not null DEFAULT 0 comment '累计游戏天数，更新lastLoginTime为跨天时候更新',
    `lastLoginTime` datetime not null DEFAULT "0000-00-00" comment '保存最后登录时间，DATEDIFF() 计算时间差',
    `ip` varchar(120) not null DEFAULT "",
    `area` varchar(120) not null DEFAULT "",
    `createTime` datetime not null DEFAULT CURRENT_TIMESTAMP,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='用户表';

create table IF not exists `login_logs`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `playerId` int(11) UNSIGNED not null,
    `ip` varchar(120) not null,
    `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='登录表';

create table IF not exists `login_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `quantity` int(11) UNSIGNED not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='登录统计结果表';
CREATE TRIGGER login_stat BEFORE INSERT ON `login_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF not exists `player_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `quantity` int(11) UNSIGNED not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='注册用户统计结果表';
CREATE TRIGGER player_stat BEFORE INSERT ON `player_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());


create table IF NOT exists `king_skills` (
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `playerId` int(11) UNSIGNED not null,
    `schemeId` int(11) not null,
    `level` int(11) not null,

     primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='魔王技能';

DROP TABLE IF EXISTS `heros`;
CREATE TABLE `heros` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `playerId` int(11) unsigned NOT NULL,
    `uid` int(11) NOT NULL,
    `schemeId` int(11) NOT NULL,
    `level` int(11) NOT NULL,
    `stage` int(11) NOT NULL,
    `rank` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `playerId_schemeId` (`playerId`,`schemeId`) USING BTREE
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='魔使表，所有魔使';

DROP TABLE IF EXISTS `soldiers`;
CREATE TABLE `soldiers` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `playerId` int(11) unsigned NOT NULL,
  `schemeId` int(11) NOT NULL,
  `num` int(11) NOT NULL,
  `level` int(11) NOT NULL,
  `stage` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `playerId_schemeId` (`playerId`,`schemeId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='魔物表，所有魔物';

create table IF NOT exists `buildings`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `playerId` int(11) UNSIGNED not null,
    `uid` int(11) not null,
    `schemeId` int(11) not null,
    `lv` int(11) not null,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='建筑表';

#游戏数据部分

create table IF NOT exists `level_distribution_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `level` int(11) not null,
    `value` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='玩家等级分布表';
CREATE TRIGGER level_distribution_stat BEFORE INSERT ON `level_distribution_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());


create table IF NOT exists `resource_distribution_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `level` int(11) not null comment 'player level',
    `type` enum('soldier','hero','building','skill','stone') not null,
    `schemeId` int(11) not null ,
    `value` varchar(120) not null comment '满级率=魔物当前等级和/魔物当前最高等级和',
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='资源分布';
CREATE TRIGGER resource_distribution_stat BEFORE INSERT ON `resource_distribution_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `gold_distribution_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `level` int(11) not null comment 'player level',
    `gold` varchar(120) not null comment '',
    `free_gold` varchar(120) not null comment '',
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='金币分布';
CREATE TRIGGER gold_distribution_stat BEFORE INSERT ON `gold_distribution_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `stage_logs`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `playerId` int(11) not null,
    `level` int(11) not null,
    `schemeId` int(11) not null,
    `status` enum('begin','end') not null,
    `isPassed` boolean not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='PVE关卡日志表';


create table IF NOT exists `stage_begin_times_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `level` int(11) not null,
    `value` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='PVE关卡执行次数表';
CREATE TRIGGER stage_begin_times_stat BEFORE INSERT ON `stage_begin_times_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());


create table IF NOT exists `stage_progress_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `schemeId` int(11) not null,
    `value` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='PVE关卡进度情况表';
CREATE TRIGGER stage_progress_stat BEFORE INSERT ON `stage_progress_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());


create table IF NOT exists `stage_difficult_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `schemeId` int(11) not null,
    `level` int(11) not null,
    `value` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='PVE关卡难度情况表';
CREATE TRIGGER stage_difficult_stat BEFORE INSERT ON `stage_difficult_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());


create table IF NOT exists `pvp_times`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `type` enum('掠夺','资源争夺','竞技场') not null comment '',
    `level` int(11) not null,
    `pvpLevel` int(11) not null,
    `value` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='PVP次数表';

create table IF NOT exists `resource_collect_log`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `playerId` int(11) not null,
    `type` enum('blood','soul') not null,
    `value` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='资源建筑收集表';

create table IF NOT exists `resource_output_blood_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `level` int(11) not null,

    `value` varchar(120) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='资源产出分布表';
CREATE TRIGGER resource_output_blood_stat BEFORE INSERT ON `resource_output_blood_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `resource_output_soul_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `level` int(11) not null,

    `value` varchar(120) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='资源产出分布表';
CREATE TRIGGER resource_output_soul_stat BEFORE INSERT ON `resource_output_soul_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());


create table IF NOT exists `resource_building_blood_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `level` int(11) not null,

    `value` varchar(120) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='资源产出分布表';
CREATE TRIGGER resource_building_blood_stat BEFORE INSERT ON `resource_building_blood_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `resource_building_soul_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `level` int(11) not null,

    `value` varchar(120) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='资源产出分布表';
CREATE TRIGGER resource_building_soul_stat BEFORE INSERT ON `resource_building_soul_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `stone_exchange_log`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `playerId` int(11) not null,
    `schemeId` int(11) not null comment 'item scheme id',
    `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='符石货币兑换道具的日志表';

create table IF NOT exists `stone_exchange_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `schemeId` int(11) not null comment 'item scheme id',
    `quantity` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='符石货币兑换道具的分布表';
CREATE TRIGGER stone_exchange_stat BEFORE INSERT ON `stone_exchange_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());


###流失部分

create table IF NOT exists `lost_player_times_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `addedUpTime` int(11) not null comment '累计天数',
    `quantity` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='流失玩家累计游戏天数';
CREATE TRIGGER lost_player_times_stat BEFORE INSERT ON `lost_player_times_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());


create table IF NOT exists `lost_player_pay_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `value` varchar(120) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='流失玩家付费率';
CREATE TRIGGER lost_player_pay_stat BEFORE INSERT ON `lost_player_pay_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

#付费部分

create table IF NOT exists `pay_logs`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `playerId` int(11) not null,
    `level` int(11) not null,
    `type` enum('item','none') not null,
    `schemeId` int(11) not null,
    `gold` int(11) not null,
    `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='玩家付费日志';

create table IF NOT exists `charge_logs`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `playerId` int(11) not null,
    `level` int(11) not null,
    `schemeId` int(11) not null,
    `gold` int(11) not null,
    `price` int(11) UNSIGNED not null,
    `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='玩家充值日志';


create table IF NOT exists `charge_total_by_day_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `quantity` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='充值总额';
CREATE TRIGGER charge_total_by_day_stat BEFORE INSERT ON `charge_total_by_day_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());



create table IF NOT exists `charge_total_by_level_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `level` int(11) not null,
    `quantity` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='各等级充值总额';
CREATE TRIGGER charge_total_by_level_stat BEFORE INSERT ON `charge_total_by_level_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `first_charge_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `level` varchar(120) not null,
    `quantity` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='第一次充值时玩家的等级';
CREATE TRIGGER first_charge_stat BEFORE INSERT ON `first_charge_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `first_pay_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `type` enum('item','none') not null,
    `schemeId` int(11) not null,
    `quantity` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='第一次付费时玩家的等级';
CREATE TRIGGER first_pay_stat BEFORE INSERT ON `first_pay_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());


create table IF NOT exists `pay_items_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `type` enum('item','none') not null,
    `schemeId` int(11) not null,
    `quantity` int(11) not null,
    `gold` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='消费道具';
CREATE TRIGGER pay_items_stat BEFORE INSERT ON `pay_items_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());


create table IF NOT exists `charge_sum_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `section` int(11) not null,
    `quantity` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='单账户付费最高额度';
CREATE TRIGGER charge_sum_stat BEFORE INSERT ON `charge_sum_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());


create table IF NOT exists `charge_item_stat`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `schemeId` int(11) not null comment '充值项id',
    `quantity` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='充值项比例';
CREATE TRIGGER charge_item_stat BEFORE INSERT ON `charge_item_stat` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());