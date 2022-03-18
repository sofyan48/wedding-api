-- +goose Up
-- +goose StatementBegin
CREATE TABLE `guest` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `wa_number` varchar(100) NOT NULL,
  `address` varchar(100) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  `is_send` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `guest_UN` (`wa_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `guest`;
-- +goose StatementEnd
