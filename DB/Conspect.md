## База

**Создание таблицы**
```
CREATE TABLE users
(
	id BIGINT NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	first_name VARCHAR(64) NOT NULL,
	last_name VARCHAR(64) NOT NULL,
	email VARCHAR(64) NOT NULL
);
```
**Удаление таблицы**
```
 DROP TABLE users (IF EXISTS)
```
**Вставка в таблицу**
```
INSERT INTO users (id, first_name, last_name, email)
VALUES (2, 'Mark', 'Zuker', 'mark@gmail.com');
INSERT INTO users (id, first_name, last_name, email)
VALUES (3, 'Elon', 'Musk', 'musk@gmail.com');
```

**Изменение таблицы**
```
ALTER TABLE goods ALTER COLUMN price SET NOT NULL
ALTER TABLE goods ADD price NUMERIC(18,2) NULL
```

**SELECT**
```
SELECT product_id, product_name
FROM goods;
```

**Обновление данных в таблице**
```
UPDATE users SET
email = 'whatever@gmail.com'
WHERE id = 1
```

**Удаление из таблицы**
```
DELETE FROM users
WHERE id=2 OR id=3
```

**Создание таблицы с внешним ключом**
```
CREATE TABLE spendings
(
	id BIGINT NOT NULL PRIMARY KEY,
	price INT NOT NULL,
	created_at TIMESTAMP DEFAULT now(),
	user_id BIGINT NOT NULL,
	
	CONSTRAINT user_id_fk FOREIGN KEY(user_id) REFERENCES users(id)
);
```

**Объединение таблиц**
```
SELECT* FROM spendings
JOIN users ON users.id = spendings.user_id
--------
SELECT spendings.*, users.first_name FROM spendings
JOIN users ON users.id = spendings.user_id
--------
SELECT orders.id, created_at, clients.name, clients.email,  products.name, products.price, orderdetails.quantity
FROM orders
INNER JOIN clients ON clients.id = orders.client_id
INNER JOIN orderdetails ON orders.id = orderdetails.orders_id
INNER JOIN products ON products.id = orderdetails.products_id
```

**Объединение таблиц со всеми юзерами (Включая тех что без пары)**
```
SELECT* FROM spendings
JOIN users ON users.id = spendings.user_id

SELECT spendings.*, users.first_name FROM spendings
RIGHT OUTER JOIN users ON users.id = spendings.user_id (RIGHT, LEFT, FULL)
```

```
INSERT INTO clients (id, name, email) VALUES (1, 'Alex', 'alex@gmail.com');
INSERT INTO clients (id, name, email) VALUES (2, 'Bob', 'bob@gmail.com');
INSERT INTO clients (id, name, email) VALUES (3, 'Mark', 'mark@gmail.com');
INSERT INTO clients (id, name, email) VALUES (4, 'Aragorn', 'aragorn@gmail.com');


```





