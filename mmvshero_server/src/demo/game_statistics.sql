create table IF not exists `login`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `quantity` int(11) UNSIGNED not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='登录统计结果表';
CREATE TRIGGER login BEFORE INSERT ON `login` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF not exists `player`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `quantity` int(11) UNSIGNED not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='注册用户统计结果表';
CREATE TRIGGER player BEFORE INSERT ON `player` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `level_distribution`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `level` int(11) not null,
    `value` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='玩家等级分布表';
CREATE TRIGGER level_distribution BEFORE INSERT ON `level_distribution` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table if not exists `hero`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计(活跃用户3，周活跃用户7，总用户0)',
    `level` int(11) not null comment 'player level',
    `schemeId` int(11) not null ,
    `averageRank` varchar(120) not null comment '平均品阶',
    `averageLevel` varchar(120) not null comment '平均等级',
    `averageSkillLv` varchar(120) not null comment '平均技能等级',
    `createTime` date NOT NULL,
    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='玩家魔使资源分布';
CREATE TRIGGER hero BEFORE INSERT ON `hero` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table if not exists `soldier`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计(活跃用户3，周活跃用户7，总用户0)',
    `level` int(11) not null comment 'player level',
    `schemeId` int(11) not null ,
    `averageLevel` varchar(120) not null comment '平均等级',
    `averageSkillLv` varchar(120) not null comment '平均技能等级',
    `createTime` date NOT NULL,
    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='玩家魔物资源分布';
CREATE TRIGGER soldier BEFORE INSERT ON `soldier` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table if not exists `building`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计(活跃用户3，周活跃用户7，总用户0)',
    `level` int(11) not null comment 'player level',
    `schemeId` int(11) not null ,
    `averageLevel` varchar(120) not null comment '平均等级',
    `createTime` date NOT NULL,
    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='玩家建筑资源分布';
CREATE TRIGGER building BEFORE INSERT ON `building` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table if not exists `kingskill`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计(活跃用户3，周活跃用户7，总用户0)',
    `level` int(11) not null comment 'player level',
    `schemeId` int(11) not null ,
    `averageLevel` varchar(120) not null comment '平均等级',
    `createTime` date NOT NULL,
    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='玩家技能资源分布';
CREATE TRIGGER kingskill_stat BEFORE INSERT ON `kingskill` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `soul`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `level` int(11) not null comment 'player level',
    `soul` varchar(120) not null comment '总量',
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='玩家金币资源分布';
CREATE TRIGGER soul_stat BEFORE INSERT ON `soul` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `gold`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `level` int(11) not null comment 'player level',
    `gold` varchar(120) not null comment '总量',
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='玩家钻石资源分布';
CREATE TRIGGER gold_stat BEFORE INSERT ON `gold` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `stage_begin_times`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `level` int(11) not null,
    `value` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='PVE关卡执行次数表';
CREATE TRIGGER stage_begin_times_stat BEFORE INSERT ON `stage_begin_times` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `stage_progress`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `schemeId` int(11) not null,
    `value` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='PVE关卡进度情况表';
CREATE TRIGGER stage_progress_stat BEFORE INSERT ON `stage_progress` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `stage_difficult`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `schemeId` int(11) not null,
    `level` int(11) not null,
    `value` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='PVE关卡难度情况表';
CREATE TRIGGER stage_difficult_stat BEFORE INSERT ON `stage_difficult` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

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

create table IF NOT exists `resource_output_blood`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `loginType` int(11) not null comment '按几天内登录过统计',
    `level` int(11) not null,

    `value` varchar(120) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='资源产出分布表';
CREATE TRIGGER resource_output_blood_stat BEFORE INSERT ON `resource_output_blood` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());


###流失部分
create table IF NOT exists `lost_player_times`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `addedUpTime` int(11) not null comment '累计天数',
    `quantity` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='流失玩家累计游戏天数';
CREATE TRIGGER lost_player_times BEFORE INSERT ON `lost_player_times` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `lost_player_pay`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `value` varchar(120) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='流失玩家付费率';
CREATE TRIGGER lost_player_pay BEFORE INSERT ON `lost_player_pay` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `charge_total_by_day`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `quantity` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='充值总额';
CREATE TRIGGER charge_total_by_day BEFORE INSERT ON `charge_total_by_day` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `charge_total_by_level`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `level` int(11) not null,
    `quantity` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='各等级充值总额';
CREATE TRIGGER charge_total_by_level BEFORE INSERT ON `charge_total_by_level` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `first_charge`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `level` varchar(120) not null,
    `quantity` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='第一次充值时玩家的等级';
CREATE TRIGGER first_charge BEFORE INSERT ON `first_charge` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `first_pay`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `type` enum('item','none') not null,
    `schemeId` int(11) not null,
    `quantity` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='第一次付费时玩家的等级';
CREATE TRIGGER first_pay BEFORE INSERT ON `first_pay` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `pay_items`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `type` enum('item','none') not null,
    `schemeId` int(11) not null,
    `quantity` int(11) not null,
    `gold` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='消费道具';
CREATE TRIGGER pay_items BEFORE INSERT ON `pay_items` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `charge_sum`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `section` int(11) not null,
    `quantity` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='单账户付费最高额度';
CREATE TRIGGER charge_sum BEFORE INSERT ON `charge_sum` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());

create table IF NOT exists `charge_item`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `schemeId` int(11) not null comment '充值项id',
    `quantity` int(11) not null,
    `createTime` date NOT NULL,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='充值项比例';
CREATE TRIGGER charge_item BEFORE INSERT ON `charge_item` FOR EACH ROW SET NEW.createTime = IFNULL(NEW.createTime, NOW());