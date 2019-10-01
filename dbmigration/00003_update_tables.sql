-- +goose Up
UPDATE users SET username='test' WHERE username LIKE 'root'       ;
UPDATE users SET username='gil', name='gilbert', surname='ndenzi', extra_column='ryumugabe' WHERE username LIKE 'vojtechvitek';