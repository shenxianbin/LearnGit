create table IF NOT exists `players`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT COMMENT "UID",
    `lv` int(11) not null DEFAULT 0,
    `soul` int(11) not null DEFAULT 0 comment '金币',
    `gold` int(11) not null DEFAULT 0 comment '钻石',
    `trophy` int(11) not null DEFAULT 0,
    `totalCharge` int(11) not null DEFAULT 0,
    `addedUpTime` int(11) not null DEFAULT 0 comment '累计游戏天数，更新lastLoginTime为跨天时候更新',
    `lastLoginTime` datetime not null DEFAULT "0000-00-00" comment '保存最后登录时间，DATEDIFF() 计算时间差',
    `ip` varchar(120) not null DEFAULT "",
    `area` varchar(120) not null DEFAULT "",
    `createTime` datetime not null DEFAULT CURRENT_TIMESTAMP,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='用户表';

create table IF NOT exists `king_skills` (
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `playerId` int(11) UNSIGNED not null,
    `schemeId` int(11) not null,
    `level` int(11) not null,
     primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='魔王技能';

CREATE TABLE IF NOT exists `heros` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `playerId` int(11) unsigned NOT NULL,
    `schemeId` int(11) NOT NULL,
    `level` int(11) NOT NULL,
    `stage` int(11) NOT NULL,
    `rank` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `playerId_schemeId` (`playerId`,`schemeId`) USING BTREE
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='魔使表，所有魔使';

CREATE TABLE IF NOT exists `soldiers` (
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