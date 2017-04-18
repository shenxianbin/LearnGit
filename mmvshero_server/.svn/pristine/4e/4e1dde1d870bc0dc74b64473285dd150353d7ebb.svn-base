create table IF not exists `login`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `playerId` int(11) UNSIGNED not null,
    `ip` varchar(120) not null,
    `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='登录日志表';

create table IF NOT exists `stage`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `playerId` int(11) not null,
    `level` int(11) not null,
    `schemeId` int(11) not null,
    `status` enum('begin','end') not null,
    `isPassed` boolean not null,
    `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='PVE关卡日志表';

create table IF NOT exists `pay`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `playerId` int(11) not null,
    `level` int(11) not null,
    `type` enum('item','none') not null,
    `schemeId` int(11) not null,
    `gold` int(11) not null,
    `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='玩家付费日志';

create table IF NOT exists `charge`(
    `id` int(11) UNSIGNED not NULL AUTO_INCREMENT,
    `playerId` int(11) not null,
    `level` int(11) not null,
    `schemeId` int(11) not null,
    `gold` int(11) not null,
    `price` int(11) UNSIGNED not null,
    `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,

    primary key (`id`)
)engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='玩家充值日志';