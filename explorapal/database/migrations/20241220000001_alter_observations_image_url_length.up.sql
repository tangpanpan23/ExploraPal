-- 增加 observations 表 image_url 字段长度以支持 base64 图片数据
ALTER TABLE `observations` MODIFY COLUMN `image_url` LONGTEXT NOT NULL COMMENT '图片URL或Base64数据';
