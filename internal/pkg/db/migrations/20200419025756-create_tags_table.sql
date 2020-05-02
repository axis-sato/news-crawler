
-- +migrate Up
CREATE TABLE IF NOT EXISTS `tags` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `is_followed` TINYINT(2) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `tags`;
