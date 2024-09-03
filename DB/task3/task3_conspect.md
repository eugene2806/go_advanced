Псевдонимы ---
SELECT g.product_id AS id,
       g.product_name AS product_name,
       g.price AS "Цена"
FROM goods AS g;
---

Обращение к столбцам таблицы через псевдоним
SELECT g.product_id, g.price
FROM goods AS g
WHERE g.price > 100;
---

Псевдонимы столбцов в списке выборки
SELECT product_id, price AS product_price
FROM goods
ORDER BY product_price;
---


