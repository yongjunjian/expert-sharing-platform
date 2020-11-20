CREATE TABLE IF NOT EXISTS `user`(
  `userId` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `userName` VARCHAR(32) NOT NULL COMMENT '用户姓名',

  `userAccount` VARCHAR(32) NOT NULL COMMENT '用户帐号',
  `phoneNumber` VARCHAR(16) DEFAULT NULL COMMENT '注册用的电话号码, 可用于登录',
  `userPassword` VARCHAR(32) NOT NULL COMMENT '用户密码',
  `credit` INT DEFAULT 0 COMMENT '用户信用积分',
  `details` VARCHAR(1024) DEFAULT NULL COMMENT '用户详情,json格式',
  PRIMARY KEY (`userId`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `expert`(
  `exportId` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '专家ID',
  `userId` BIGINT NOT NULL COMMENT '该专家对应的用户ID',
  `expertLevel` TINYINT DEFAULT 1 COMMENT '专家等级',
  `domainId` INT DEFAULT 1 COMMENT '专家领域',
  `fortune` INT DEFAULT 0 COMMENT '赚取的`财富量，虚拟币',
)
 
CREATE TABLE IF NOT EXISTS `domain` (
  `domainId` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '领域ID',
  `domainName` VARCHAR(32) NOT NULL COMMENT '领域名',
  `description` TEXT NOT NULL COMMENT '领域描述',
  `parentId` INT NOT NULL COMMENT '父领域的id,根结点的父领域id为0',
  PRIMARY KEY (`domainId`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `task` (
  `taskId` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `domainId` INT NOT NULL COMMENT '所属领域',
  `problemDesc` VARCHAR(1024) NOT NULL COMMENT '问题描述',
  `standardForSolution` VARCHAR(1024) NOT NULL COMMENT '问题成功解决的标准',
  `bonus` INT NOT NULL COMMENT '承诺的问题解决后的奖金, 虚拟币',
  `fine` INT NOT NULL COMMENT '问题未能及时解决的罚金, 虚拟币',
  `deadline` TIMESTAMP NOT NULL COMMENT '问题解决的最后期限',
  `publisherId` INT NOT NULL COMMENT '发布者的用户id',
  `expertId` INT DEFAULT NULL COMMENT ' 专家id',
  `status` TINYINT DEFAULT 0 COMMENT '任务状态,0-未开始 1-解决中 2-已解决 3-已取消',
  PRIMARY KEY (`taskId`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `solution` (
  `solutionId` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '方案ID',
  `domainId` INT NOT NULL COMMENT '所属领域',
  `taskId` BIGINT NOT NULL COMMENT '关联的任务id',
  `problemDesc` VARCHAR(1024) NOT NULL COMMENT '问题描述',
  `abstract` VARCHAR(1024) NOT NULL COMMENT '方案摘要',
  `detail` TEXT NOT NULL COMMENT '方案详情',
  `communicateInfo` VARCHAR(1024) NOT NULL COMMENT '沟通记录，json格式',
  `publisherId` BIGINT NOT NULL COMMENT '发布者的用户id',
  `publisherPermitted` INT NOT NULL COMMENT '发布者是否授权出售该方案',
  `expertId` BIGINT DEFAULT NULL COMMENT '执行者的用户id',
  `solverPermitted` BIGINT NOT NULL COMMENT '解决者是否授权出售该方案',
  `price` INT DEFAULT NULL COMMENT '该条方案的价格',
  `shareRate` FLOAT DEFAULT  1  COMMENT '分成比例，发布者:专家',
  PRIMARY KEY (`solutionId`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `message`(
  `messageId` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '消息ID',
  `time` TIMESTAMP NOT NULL  COMMENT '时间戳',
  `taskId` BIGINT NOT NULL  COMMENT '任务id',
  `direction` TINYINT DEFAULT 0  COMMENT '方向,0:p->s, 1:s->p',
  `type` TINYINT DEFAULT 0 COMMENT '类型: 0-文本 1-图片 2-音频 3-视频',
  `data` LONGBLOB NOT NULL COMMENT '数据内容',
PRIMARY KEY (`messageId`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8
