
-- +migrate Up
CREATE TABLE IF NOT EXISTS `articles` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(255) NOT NULL,
  `url` VARCHAR(255) NOT NULL,
  `image` TEXT NULL,
  `crawled_at` TIMESTAMP NOT NULL,
  `site_id` INT UNSIGNED NOT NULL,
  `original_id` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_articles_site_idx` (`site_id` ASC) VISIBLE,
  INDEX `crawled_at_idx` (`crawled_at` ASC) VISIBLE,
  UNIQUE INDEX `site_id_original_id_unique` (`site_id` ASC, `original_id` ASC) VISIBLE,
  CONSTRAINT `fk_articles_sites`
    FOREIGN KEY (`site_id`)
    REFERENCES `sites` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `articles`;
