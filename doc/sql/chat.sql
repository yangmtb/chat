
CREATE TABLE `chat_account` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(32) NOT NULL COMMENT '用户名',
  `password` varchar(64) NOT NULL COMMENT '密码',
  `nickname` varchar(32) NOT NULL COMMENT '昵称',
  `phone` varchar(15) NOT NULL COMMENT '手机号',
  `email` varchar(32) COMMENT '邮箱',
  `salt` varchar(32) NOT NULL COMMENT '盐',
  `ctime` timestamp DEFAULT NOW() COMMENT '创建时间',
  `mtime` timestamp COMMENT '修改时间',
  `dtime` timestamp COMMENT '删除时间',
  `level` tinyint(3) unsigned DEFAULT '1' COMMENT '等级',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='帐号管理';
