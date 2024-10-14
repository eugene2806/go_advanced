## Задача 1

До текущего момента мы пользовались схемой данных, в таблицах которой отсутствовали ограничения первичного ключа.  

Однако так быть не должно, поэтому Вам необходимо добавить в каждую таблицу ограничение первичного ключа.  

Напишите инструкцию добавления первичного ключа для каждой таблицы.  

``` 
ALTER TABLE buyers ADD CONSTRAINT pk_buyers PRIMARY KEY (buyer_id);
ALTER TABLE orders ADD CONSTRAINT pk_orders PRIMARY KEY (order_id);
ALTER TABLE orders_products ADD CONSTRAINT pk_orders_products PRIMARY KEY (order_id, product_id);
ALTER TABLE products ADD CONSTRAINT pk_products PRIMARY KEY (product_id);
ALTER TABLE product_types ADD CONSTRAINT pk_product_types PRIMARY KEY (type_id);
```

## Задача 2
В нашей схеме данных между сущностями существует определенная связь (вспоминаем занятие «Типы связей»).

Физически связь между таблицами, как Вы помните, реализуется с помощью внешних ключей.

Однако в наших таблицах физически связь никак не реализована, т.е. у нас нет никаких внешних ключей.  
Мы писали запросы, опираясь только на визуальное, логическое представление данных.

Реальная эксплуатация такой схемы данных будет вызывать огромное количество аномалий в процессе добавления,  
изменения и удаления данных в таблицах.

Чтобы исправить ситуацию, мы должны физически отразить связь между таблицами в базе данных.  
Это делается с помощью ограничений внешнего ключа.

Напишите инструкции добавления ограничений внешнего ключа в существующие таблицы, чтобы физически реализовать связь,  
которая будет соответствовать уже знакомой ER-диаграмме.  

``` 
ALTER TABLE orders ADD CONSTRAINT fk_buyers FOREIGN KEY (buyer_id) REFERENCES buyers (buyer_id);
ALTER TABLE products ADD CONSTRAINT fk_product_types FOREIGN KEY (type_id) REFERENCES product_types (type_id);
ALTER TABLE orders_products ADD CONSTRAINT fk_products FOREIGN KEY (product_id) REFERENCES products (product_id);
ALTER TABLE orders_products ADD CONSTRAINT fk_orders FOREIGN KEY (order_id) REFERENCES orders (order_id)
``` 

## Задача 3

Руководство магазина снова приняло решение ввести ограничение, в соответствии с которым номер заказа не должен  
повторяться в рамках рабочего дня.

Вам поставили задачу реализовать данное ограничение на уровне базы данных.

Напишите инструкцию добавления соответствующего ограничения в таблицу с заказами (orders).

``` 
ALTER TABLE orders ADD CONSTRAINT uq_order_number_date UNIQUE (order_number, order_date)
```

## Задача 4

Согласно требованиям нашего магазина сумма заказа не должна быть равной 0 или быть отрицательной.  

Поэтому Вам поставили задачу реализовать такое ограничение на уровне базы данных.  

Напишите инструкцию добавления соответствующего ограничения в таблицу с заказами (orders).  

```
ALTER TABLE orders ADD CONSTRAINT ck_summa CHECK (order_summa > 0)
```

## Задача 5

В процессе эксплуатации системы у Вас возникла необходимость временно удалить все ограничения из базы данных.

Напишите инструкции для удаления всех ограничений, которые Вы создали ранее при решении предыдущих задач.

``` 
ALTER TABLE orders DROP CONSTRAINT fk_buyers;
ALTER TABLE orders_products DROP CONSTRAINT fk_orders;
ALTER TABLE products DROP CONSTRAINT fk_product_types;

ALTER TABLE orders DROP CONSTRAINT pk_orders;
ALTER TABLE orders DROP CONSTRAINT uq_order_number_date;
ALTER TABLE orders DROP CONSTRAINT ck_summa;

ALTER TABLE orders_products DROP CONSTRAINT pk_orders_products;
ALTER TABLE orders_products DROP CONSTRAINT fk_products;

ALTER TABLE products DROP CONSTRAINT pk_products;

ALTER TABLE product_types DROP CONSTRAINT pk_product_types;

ALTER TABLE buyers DROP CONSTRAINT pk_buyers;
```



