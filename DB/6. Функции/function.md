### Задача 1
У Вас возникла необходимость определить товары, в описании которых присутствует слово «МГц»,  
однако при выводе данных Вам необходимо заменить данное слово на «MHz».

Выборка должна включать наименование товара и его описание с уже заменённым словом.

``` 
SELECT products.product_name,
       REPLACE (products.description, 'МГц', 'MHz') AS "replace"
FROM products
WHERE products.description LIKE '%МГц%'
```

### Задача 2
У Вас возникла необходимость вывести инициалы покупателей в отдельном столбце.

Выборка должна включать столбец с полным именем и столбец, содержащий только инициалы,  
при этом исключите из выборки покупателя с идентификатором 1028.

``` 
SELECT buyers.buyer_name,
       RIGHT(TRIM(buyers.buyer_name), 4) AS "initials"
FROM buyers
WHERE buyer_id <> 1028
```

### Задача 3
Вам поставили задачу определить, какие товары, связанные с процессором, приобретал покупатель «Андреев В.А.».

Напишите универсальный запрос, который будет выводить такие товары, приобретенные данным покупателем.

Выборка должна включать имя покупателя и наименование товара.

``` 
SELECT buyers.buyer_name, products.product_name
FROM buyers
INNER JOIN orders ON buyers.buyer_id = orders.buyer_id
INNER JOIN orders_products ON orders.order_id = orders_products.order_id
INNER JOIN products ON orders_products.product_id = products.product_id
WHERE buyer_name = 'Андреев В.А.' AND LOWER(product_name) LIKE '%процессор%'
```

### Задача 4
У Вас возникла необходимость определить среднюю стоимость заказа для каждого покупателя за период с Апреля по Июнь 2020,  
при этом в выборку должны попасть покупатели, сделавшие более одного заказа,  
а также итоговую среднюю стоимость необходимо округлить до 2-х знаков после запятой.

Выборка должна включать имя покупателя, среднюю стоимость заказа и количество заказов.  
Итоговый результат отсортируйте по уменьшению средней стоимости заказа.  

``` 
SELECT buyers.buyer_name,
       ROUND(AVG(orders.order_summa), 2) AS "average_coast",
       COUNT(*) AS "quantity_orders"
FROM buyers
INNER JOIN orders ON buyers.buyer_id = orders.buyer_id
WHERE order_date BETWEEN '2020.04.01' AND '2020.06.30'
GROUP BY buyers.buyer_name
HAVING COUNT(*) > 1
ORDER BY quantity_orders DESC
```

### Задача 5
У Вас возникла необходимость получить детальную информацию о товарах (включая описание и наименование типа товара),  
которые приобретал покупатель «Иванов И.И.», при этом если описание товара не заполнено, необходимо выводить фразу «Нет описания».  
Кроме этого предусмотрите, чтобы все текстовые поля выводились без пробелов в начале и в конце.

В выборке каждый товар должен быть в единственном экземпляре. Строки отсортируйте по идентификатору товара.  

``` 
SELECT DISTINCT TRIM(buyers.buyer_name),
       products.product_id,
       TRIM(products.product_name),
       COALESCE (TRIM(products.description), 'Нет описания') AS "description",
       products.price,
       TRIM(product_types.name)
FROM buyers
INNER JOIN orders ON orders.buyer_id = buyers.buyer_id
INNER JOIN orders_products ON orders_products.order_id = orders.order_id
INNER JOIN products ON products.product_id = orders_products.product_id
INNER JOIN product_types ON product_types.type_id = products.type_id
WHERE buyer_name = 'Иванов И.И.'
ORDER BY product_id
```
