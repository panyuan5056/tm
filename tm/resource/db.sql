CREATE TABLE `fx_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `fx_auth` (`id`, `username`, `password`) VALUES (null, 'root', '111111');

 
CREATE TABLE `fx_strategy` (
	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	`class1` varchar(200) DEFAULT '' COMMENT '类型1',
	`class2` varchar(200) DEFAULT '' COMMENT '类型2',
	`class3` varchar(200) DEFAULT '' COMMENT '类型3',
	`name` varchar(100) DEFAULT '' COMMENT '名称',
	`dept` varchar(50) DEFAULT '' COMMENT '深度',
	`category` varchar(50) DEFAULT '' COMMENT '对应处理类型',
	`plan` varchar(50) DEFAULT '' COMMENT '发现方法',
	`desc` varchar(500) DEFAULT '' COMMENT '描述',
	`forward` varchar(100) DEFAULT '' COMMENT '对应名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
