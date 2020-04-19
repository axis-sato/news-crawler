
-- +migrate Up
CREATE TABLE IF NOT EXISTS `article_tags` (
  `article_id` INT UNSIGNED NOT NULL,
  `tag_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`article_id`, `tag_id`),
  INDEX `fk_article_tags_tags_idx` (`tag_id` ASC) VISIBLE,
  CONSTRAINT `fk_article_tags_articles`
    FOREIGN KEY (`article_id`)
    REFERENCES `articles` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_article_tags_tags`
    FOREIGN KEY (`tag_id`)
    REFERENCES `tags` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

-- +migrate Down
DROP TABLE `article_tags`;
