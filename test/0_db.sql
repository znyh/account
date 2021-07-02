create database kratos_demo;
use kratos_demo;

CREATE TABLE `articles` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `title` varchar(64) NOT NULL COMMENT '名称',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `ix_mtime` (`mtime`)
) COMMENT='文章表';

CREATE TABLE `account` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `Nick`  varchar(64) NOT NULL COMMENT 'Nick',
  `Password` varchar(64) NOT NULL COMMENT '密码',
  `Telephone` varchar(1024) NOT NULL COMMENT '手机号',
  `OS` int(1) NOT NULL COMMENT '系统',
  `Sex` int(1) NOT NULL COMMENT 'Sex',
  `UserID` int(64) NOT NULL COMMENT 'Sex',
  `HealUrl` varchar(1024) NOT NULL COMMENT 'HealUrl',
  `City` varchar(1024) NOT NULL COMMENT 'City',
  `Birth` varchar(1024) NOT NULL COMMENT 'Birth',
  `Age` int(10) NOT NULL COMMENT 'Age',
  `Online` varchar(1024) NOT NULL COMMENT 'Online',
  `LastLoginTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `CreateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `IsRobot` varchar(1024) NOT NULL COMMENT 'IsRobot',
  PRIMARY KEY (`id`),
  KEY `ix_mtime` (`LastLoginTime`)
) COMMENT='用户表';


-- func (d *dao) GuestRegister1(ctx context.Context, id int32, title string) (err error) {
-- 	query := "INSERT articles SET id=?,title=?"
-- 	stmt, err := d.db.Prepare(query)
-- 	if err != nil {
-- 		log.Error("insert prepare, (id:%+v, title:%+v) query:%s error: ", id, title, query, err.Error())
-- 		return
-- 	}
-- 	defer stmt.Close()
--
-- 	_, err = stmt.Exec(ctx, id, title)
-- 	if err != nil {
-- 		log.Error("insert exec error,(id:%+v, title:%+v) query:%s err: ", id, title, query, err.Error())
-- 		return
-- 	}
-- 	return
-- }