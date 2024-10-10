### Задача 1

**Вам постоянно требуется знать всю информацию о заказах и товарах, которые приобретают покупатели в Вашем магазине.  
Для удобства Вы решили создать представление sales_products, которое будет возвращать детализированные данные о продажах.**  

Данные должны включать:  

* Идентификатор и имя покупателя  
* Идентификатор, номер, дата и сумма заказа  
* Идентификатор, наименование и стоимость товара  
* Идентификатор и наименование типа товара  
* Напишите инструкцию создания соответствующего представления.  

```
CREATE VIEW full_info
AS
SELECT buyers.buyer_id,
       buyers.buyer_name,
       orders.order_id,
       orders.order_number,
       orders.order_date,
       orders.order_summa,
       products.product_id,
       products.product_name,
       products.price,
       product_types.type_id,
       product_types.name

FROM buyers
INNER JOIN orders ON orders.buyer_id = buyers.buyer_id
INNER JOIN orders_products ON orders_products.order_id = orders.order_id
INNER JOIN products ON products.product_id = orders_products.product_id
INNER JOIN product_types ON product_types.type_id = products.type_id
```

### Задача 2

**Напишите запрос, который покажет, сколько физических товаров приобрел каждый из покупателей за все время,  
при этом в качестве источника используйте созданное ранее представление.**

**Результирующий набор должен включать имя покупателя и количество приобретенных физических товаров,  
при этом отсортируйте полученный результат по уменьшению количества покупок.**  

``` 
SELECT buyer_name,
       COUNT(*) AS "quantity_sales"
FROM full_info
WHERE type_id = 1
GROUP BY buyer_name
ORDER BY quantity_sales DESC
```

### Задача 3

**Вам постоянно требуется знать, сколько единиц каждого товара было продано за все время,  
поэтому Вы решили создать представление quantity_sales_products для удобства.**  

Представление должно возвращать:  

* Идентификатор товара  
* Наименование товара  
* Количество продаж  
* Напишите инструкцию создания соответствующего представления.  

```
CREATE VIEW quantity_sales_products
AS
SELECT products.product_id,
       products.product_name,
       COUNT(orders_products.product_id) AS "quantity_sales"
FROM products
LEFT JOIN orders_products ON orders_products.product_id = products.product_id
GROUP BY products.product_id, products.product_name
```

### Задача 4

**Напишите запрос, который покажет 2 самых продаваемых товара,  
при этом используйте в качестве источника созданное ранее представление.**  

Выборка должна включать наименование товара и количество продаж.  

``` 
SELECT product_name, quantity_sales
FROM quantity_sales_products
ORDER BY quantity_sales DESC
LIMIT 2
```

`DROP VIEW IF EXISTS quantity_sales_products, full_info`