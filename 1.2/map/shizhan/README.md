```sql
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `uid` int(10) NOT NULL AUTO_INCREMENT,
    `name` varchar(30) DEFAULT '',
    `phone` varchar(20) DEFAULT '',
    `email` varchar(30) DEFAULT '',
    `password` varchar(100) DEFAULT '',
    PRIMARY KEY (`uid`)
) ENGINE = InnoDB AUTO_INCREMENT = 3 DEFAULT CHARSET = utf8 COMMENT = '用户表';

BEGIN;
INSERT INTO `user` VALUES(1, 'lisi1', '18888888888', 'lisi@gmail.com', '123.com');
INSERT INTO `user` VALUES(2, 'lisi2', '18888888888', 'lisi@gmail.com', '123.com');
INSERT INTO `user` VALUES(3, 'lisi3', '18888888888', 'lisi@gmail.com', '123.com');
INSERT INTO `user` VALUES(4, 'lisi4', '18888888888', 'lisi@gmail.com', '123.com');
INSERT INTO `user` VALUES(5, 'lisi5', '18888888888', 'lisi@gmail.com', '123.com');
INSERT INTO `user` VALUES(6, 'lisi6', '18888888888', 'lisi@gmail.com', '123.com');
INSERT INTO `user` VALUES(7, 'lisi7', '18888888888', 'lisi@gmail.com', '123.com');
INSERT INTO `user` VALUES(8, 'lisi8', '18888888888', 'lisi@gmail.com', '123.com');
INSERT INTO `user` VALUES(9, 'lisi9', '18888888888', 'lisi@gmail.com', '123.com');
INSERT INTO `user` VALUES(10, 'lisi10', '18888888888', 'lisi@gmail.com', '123.com');
COMMIT;
```