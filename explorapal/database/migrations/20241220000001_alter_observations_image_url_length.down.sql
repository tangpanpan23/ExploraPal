-- 回滚：恢复 observations 表 image_url 字段长度
ALTER TABLE `observations` MODIFY COLUMN `image_url` VARCHAR(500) NOT NULL COMMENT '图片URL';
