### Задача 1

***У Вас возникла необходимость вывести список товаров с указанием наименования типа товара.
Напишите запрос, который выведет наименование товара, его цену, а также наименование типа товара.***

```
SELECT products.product_name, products.price, product_types.name
FROM products
INNER JOIN product_types  ON  product_types.type_id = products.type_id
```

### Задача 2

***Руководство поставило Вам задачу составить список заказов за Июнь 2020, включая детализацию по товарам.
Напишите запрос, который выведет номер и дату заказа, наименование товаров, которые включаются в тот или иной заказ,
а также цену товаров.***

```
SELECT orders.order_number,
       orders.order_date,
       products.product_name,
       products.price
FROM orders
INNER JOIN orders_products ON orders.order_id = orders_products.order_id
INNER JOIN products ON products.product_id = orders_products.product_id
WHERE order_date BETWEEN '2020.06.01' AND '2020.06.30'
```

### Задача 3

***Вам поставили задачу сформировать список из 5 самых крупных по сумме заказов с указанием имени покупателей,
сделавших данные заказы.
Напишите запрос, который покажет имя покупателя, номер, дату и сумму самых крупных по сумме заказов.***

```
SELECT buyers.buyer_name,
       orders.order_number,
       orders.order_date,
       orders.order_summa
FROM buyers
INNER JOIN orders ON  orders.buyer_id = buyers.buyer_id
ORDER BY order_summa DESC
LIMIT 5
```

### Задача 4

***Вам поставили задачу определить товары, которые не пользуются спросом, в частности товары, которые еще никто не покупал.
Напишите запрос, который выведет наименование таких товаров и их цену.***

```
SELECT products.product_name, products.price
FROM products
LEFT JOIN orders_products ON  products.product_id = orders_products.product_id
WHERE orders_products.order_id IS NULL
```

### Задача 5

***У Вас возникла необходимость определить все заказы, в которых участвует товар с наименованием «Процессор V5».
Напишите запрос, который выведет наименование товара, номер и дату заказов с данным товаром.
Данные отсортируйте по дате заказа.***

```
SELECT products.product_name,
       orders.order_number,
       orders.order_date
FROM products
INNER JOIN orders_products ON products.product_id = orders_products.product_id
INNER JOIN orders ON orders_products.order_id = orders.order_id
WHERE product_name LIKE '%Процессор V5%'
```

### Задача 6

***Вас попросили узнать, на какую сумму каждый из покупателей сделал заказы в Вашем магазине,
т.е. сколько всего денег потратил каждый из покупателей.
Напишите запрос, который выведет имя покупателя и общую сумму денег, которую он потратил,
при этом отсортируйте список по имени покупателя.***

```
SELECT buyers.buyer_name, SUM(orders.order_summa) AS "summa"
FROM buyers
LEFT JOIN orders ON buyers.buyer_id = orders.buyer_id
GROUP BY buyers.buyer_name
ORDER BY buyers.buyer_name
```

### Задача 7

***Вам поставили задачу определить, какие физические товары покупал «Иванов И.И.» в Мае 2020.
Напишите запрос, который выведет имя покупателя, дату покупки и наименование товаров, которые он покупал.***

```
SELECT buyers.buyer_name,
       orders.order_date,
       products.product_name
FROM buyers
INNER JOIN orders ON buyers.buyer_id = orders.buyer_id
INNER JOIN orders_products ON orders.order_id = orders_products.order_id
INNER JOIN products ON orders_products.product_id = products.product_id
WHERE buyer_name = 'Иванов И.И.' AND order_date BETWEEN '2020.05.01' AND '2020.05.30'
```

### Задача 8

***У Вас возникла необходимость подсчитать, сколько всего товаров было продано с разбивкой по типу товара,
т.е. сколько было продано физических товаров и сколько было продано цифровых товаров.
Напишите запрос, который выведет наименование типа товара и количество продаж.***

```
SELECT product_types.name,
       COUNT(*) AS "quantity_sales"
FROM product_types
INNER JOIN products ON product_types.type_id = products.type_id
INNER JOIN orders_products ON products.product_id = orders_products.product_id
GROUP BY product_types.name
```

### Задача 9

***Вам необходимо определить покупателей, которые купили больше 5 физических товаров.
Напишите запрос, который выведет имя покупателя и количество купленных им физических товаров,
при этом в выборку должны попасть только покупатели, купившие более 5 физических товаров,
покупатели купившие 5 или менее физических товаров нас не интересуют.***

```
SELECT buyers.buyer_name,
       COUNT(*) AS "purchase_quntity"
FROM buyers
INNER JOIN orders ON buyers.buyer_id = orders.buyer_id
INNER JOIN orders_products ON orders.order_id = orders_products.order_id
INNER JOIN products ON orders_products.product_id = products.product_id
INNER JOIN product_types ON products.type_id = product_types.type_id
WHERE product_types.type_id = 1
GROUP BY buyers.buyer_name
HAVING COUNT(orders.order_id) > 5
```