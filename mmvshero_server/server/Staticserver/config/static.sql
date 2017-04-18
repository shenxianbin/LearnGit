DROP TABLE IF EXISTS `players`;
CREATE TABLE `players` (
  `id` int(11) unsigned NOT NULL COMMENT 'UID',
  `lv` int(11) NOT NULL DEFAULT '0',
  `stone` int(11) NOT NULL DEFAULT '0',
  `gold` int(11) NOT NULL DEFAULT '0',
  `freeGold` int(11) NOT NULL DEFAULT '0',
  `trophy` int(11) NOT NULL DEFAULT '0',
  `totalCharge` int(11) NOT NULL DEFAULT '0',
  `addedUpTime` int(11) NOT NULL DEFAULT '0' COMMENT '累计游戏天数，更新lastLoginTime为跨天时候更新',
  `lastLoginTime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '保存最后登录时间，DATEDIFF() 计算时间差',
  `ip` varchar(120) NOT NULL DEFAULT '',
  `area` varchar(120) NOT NULL DEFAULT '',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

DROP TABLE IF EXISTS `king_skills`;
CREATE TABLE `king_skills` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `playerId` int(11) unsigned NOT NULL,
  `schemeId` int(11) NOT NULL,
  `level` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `playerId_schemeId` (`playerId`,`schemeId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='魔王技能';

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
  UNIQUE KEY `playerId_uid` (`playerId`,`uid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='魔使表，所有魔使';

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

DROP TABLE IF EXISTS `buildings`;
CREATE TABLE `buildings` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `playerId` int(11) unsigned NOT NULL,
  `uid` int(11) NOT NULL,
  `schemeId` int(11) NOT NULL,
  `lv` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `playerId_uid` (`playerId`,`uid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='建筑表';

DROP TABLE IF EXISTS `charge_logs`;
CREATE TABLE `charge_logs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `playerId` int(11) NOT NULL,
  `level` int(11) NOT NULL,
  `schemeId` int(11) NOT NULL,
  `gold` int(11) NOT NULL,
  `price` int(11) unsigned NOT NULL,
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家充值日志';

DROP TABLE IF EXISTS `login_logs`;
CREATE TABLE `login_logs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `playerId` int(11) unsigned NOT NULL,
  `ip` varchar(120) NOT NULL,
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='登录表';

DROP TABLE IF EXISTS `pay_logs`;
CREATE TABLE `pay_logs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `playerId` int(11) NOT NULL,
  `level` int(11) NOT NULL,
  `type` enum('item','bloodExchagePvp','bloodBuyMappoint','bloodBuyDecoration',
	'soulBuyMappoint','soulBuyDecoration','evolutionSpeedup','evolutionOnekey',
	'upBuildingSpeedup','upBuildingOnekey','upKingskillSpeedup','upKingskillOnekey') NOT NULL,
  `schemeId` int(11) NOT NULL,
  `gold` int(11) NOT NULL,
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家付费日志';

DROP TABLE IF EXISTS `resource_collect_log`;
CREATE TABLE `resource_collect_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `playerId` int(11) NOT NULL,
  `type` enum('blood','soul') NOT NULL,
  `value` int(11) NOT NULL,
  `createTime` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='资源建筑收集表';

DROP TABLE IF EXISTS `stage_logs`;
CREATE TABLE `stage_logs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `playerId` int(11) NOT NULL,
  `level` int(11) NOT NULL,
  `schemeId` int(11) NOT NULL,
  `status` enum('begin','end') NOT NULL,
  `isPassed` tinyint(1) NOT NULL,
  `createTime` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='PVE关卡日志表';

DROP TABLE IF EXISTS `stone_exchange_log`;
CREATE TABLE `stone_exchange_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `playerId` int(11) NOT NULL,
  `schemeId` int(11) NOT NULL COMMENT 'item scheme id',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='符石货币兑换道具的日志表';