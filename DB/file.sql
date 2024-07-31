SELECT * FROM clients;
----

SELECT orders.id, orders.created_at, clients.name, clients.email FROM orders
JOIN clients ON clients.id = orders.client_id;
----

SELECT orders.id, created_at, clients.name, clients.email,  products.name, products.price, orderdetails.quantity
FROM orders
INNER JOIN clients ON clients.id = orders.client_id
INNER JOIN orderdetails ON orders.id = orderdetails.orders_id
INNER JOIN products ON products.id = orderdetails.products_id;
---

UPDATE products SET
price = 60
WHERE id = 1;
---

DELETE FROM orderdetails
WHERE orders_id IN (5, 6);

DELETE FROM orders
WHERE id IN (5, 6);

DELETE FROM clients
WHERE id = 4;
