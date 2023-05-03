-- Селектим пользователей по имени
EXPLAIN ANALYSE SELECT * from users where name = 'Anthony';

-- Создаем индекс на имя пользователя
CREATE INDEX users_name_idx ON users(name);

-- Удаляем индекс на имя
DROP INDEX users_name_idx;


-- Селектим пользователя по id
EXPLAIN ANALYSE SELECT * from users where id = 1545753;

-- Создаем индекс на id пользователя
CREATE INDEX users_id_idx ON users(id);

-- Удаляем индекс на id
DROP INDEX users_id_idx;


-- Исследуем добавление нового пользователя
EXPLAIN ANALYSE INSERT INTO users (name, birthday) VALUES ('Tema', now())


SELECT COUNT(*) FROM users;