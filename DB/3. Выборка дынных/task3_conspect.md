### Псевдонимы

```
SELECT g.product_id AS id,
       g.product_name AS product_name,
       g.price AS "Цена"
FROM goods AS g;
```

Обращение к столбцам таблицы через псевдоним
```
SELECT g.product_id, g.price
FROM goods AS g
WHERE g.price > 100;
```

Псевдонимы столбцов в списке выборки
```
SELECT product_id, price AS product_price
FROM goods
ORDER BY product_price;
```

### Предикаты
* Предикат – это любое выражение, результатом которого являются значения
  TRUE, FALSE или UNKNOWN, иными словами, истина, ложь или неизвестно.

```
SELECT product_id,
       product_name,
       price
FROM goods
WHERE price > 10;
```

### Операторы SQL
* ```=``` (равно) – определяет, равняются ли сравниваемые выражения.
* ```>``` (больше) – определяет, превышает ли одно выражение другое.
* ```<``` (меньше) – оператор используется для проверки того, что одно выражение меньше другого.
* ```<>``` (не равно) – определяет неравенство двух выражений.
* ```!=``` (не равно) – оператор, который также определяет неравенство двух выражений.
* ```>=``` (больше или равно) – оператор используется для проверки превышения либо равенства двух выражений.
* ```<=``` (меньше или равно) – оператор используется для проверки того, что одно выражение меньше или равно другому.
* ```LIKE``` или NOT LIKE – определяет, содержит или не содержит текстовая строка символы, заданные в шаблоне.
* ```BETWEEN``` или NOT BETWEEN – определяет, входит ли или не входит одно выражение в заданный диапазон значений.
* ```IN``` или NOT IN – определяет, совпадает или не совпадает указанное значение с одним из значений из заданного списка или вложенного запроса.
* ```EXISTS``` – используется во вложенном запросе для проверки существования строк, возвращенных вложенным запросом (вложенные запросы рассмотрим на следующих занятиях).
* ```OR(||) - ИЛИ``` - хотя бы одно условие  ```AND(&&) - И``` оба условия 

--Пример с использованием BETWEEN
```
SELECT product_id,
       product_name,
       price
FROM goods
WHERE price BETWEEN 100 AND 500;
```

-- Пример с использованием LIKE
```
SELECT product_id,
       product_name,
       price
FROM goods
WHERE product_name LIKE 'С%'; -первый символ, '%С%' - есть ли вообще символ
```

-- Пример с использованием IN
```
SELECT product_id,
       product_name,
       price
FROM goods
WHERE price IN (50, 100); - Только 50 и 100 price = 50 OR price = 100;
```

-- Пример с использованием OR
```
SELECT product_id,
       product_name,
       price
FROM goods
WHERE price = 50 OR product_name = 'Монитор'; - 50 или 'Монитор' // true / true
```

-- Пример с использованием IS NOT NULL
```
SELECT product_id,
       product_name,
       price
FROM goods
WHERE price IS NOT NULL; - заполненные столбцы / NULL - незаполненные столбцы 
```

### Сортировка

-- Пример с использованием ORDER BY - ***по возрастанию***
```
SELECT product_name, price
FROM goods
ORDER BY price;
```

-- Пример с использованием ORDER BY DESC - ***по убыванию***
```
SELECT product_name, price
FROM goods
ORDER BY price DESC;
```

-- Пример с использованием ORDER BY и двух сортировок
```
SELECT category, product_name, price
FROM goods
ORDER BY category, price; - Сначала отсортирует по category, потом по price
```

### Фильтры

-- Пример с использованием LIMIT
```
SELECT product_id, product_name, price
FROM goods
LIMIT 3; - Выведет только первые 3 строки
```

-- Пример с использованием LIMIT
```
SELECT product_id, product_name, price
FROM goods
ORDER BY Price DESC
LIMIT 3; - Выведет только первые 3 строки отсортированные по убыванию
```

-- Пример с использованием LIMIT и OFFSET
```
SELECT product_id, product_name, price
FROM goods
ORDER BY Price DESC
LIMIT 3 OFFSET 1; Выведет только первые 3 строки отсортированные по убыванию пропуская первую строку
```

***SELECT возвращает все записи из источника, но записи в источнике могут повторяться или, может быть,
у Вас так сконструирован запрос, что они могут повторяться, при этом в результирующем наборе данных
эти повторы Вам не нужны.***

-- Пример с использованием DISTINCT
```
SELECT DISTINCT product_name, price
FROM goods;
```

### Агрегирующие функции

***Агрегирующие функции – это функции, которые выполняют агрегацию данных,
т.е. статистическую операцию на наборе данных, и возвращают одиночное итоговое значение.***

* COUNT() – вычисляет количество значений в столбце (значения NULL не учитываются).
Если написать COUNT(*), то будут учитываться все записи, т.е. все строки;
* SUM() – суммирует значения в столбце;
* MAX() – определяет максимальное значение в столбце;
* MIN() – определяет минимальное значение в столбце;
* AVG() – определяет среднее значение в столбце.

-- Пример с использованием агрегирующих функций
```
SELECT COUNT(*)   AS "Количество строк",
       SUM(price) AS "Сумма по столбцу price",
       MAX(price) AS "Максимальное значение в столбце price",
       MIN(price) AS "Минимальное значение в столбце price",
       AVG(price) AS "Среднее значение в столбце price"
FROM goods;
```

### GROUP BY
***А теперь давайте представим, что нам нужно все то же самое, только с группировкой по категориям.
Другими словами, мы хотим знать количество товаров в определенной категории, максимальную,
минимальную и среднюю цену товаров в каждой из категорий (сумму мы уберем, а то как-то нелогично суммировать цену).
Мы для этого используем предложение GROUP BY***

-- Пример с использованием GROUP BY
```
SELECT category,
       COUNT(*)   AS "Количество строк",
       MAX(price) AS "Максимальное значение в столбце Price",
       MIN(price) AS "Минимальное значение в столбце Price",
       AVG(price) AS "Среднее значение в столбце Price"
FROM goods
GROUP BY category;
```

-- Второй пример с использованием GROUP BY
```
SELECT category,
COUNT(*)   AS "Количество строк",
MAX(price) AS "Максимальное значение в столбце price",
MIN(price) AS "Минимальное значение в столбце price",
AVG(price) AS "Среднее значение в столбце price"
FROM goods
WHERE product_id <> 1
GROUP BY category;
```

### HAVING

***HAVING – это условие, фильтрующее сгруппированные данные, образованные агрегатной функцией.***

-- Пример с использованием HAVING
```
SELECT category AS "Id категории",
COUNT(*) AS "Количество товаров"
FROM goods
GROUP BY category
HAVING COUNT(*) > 1; - сгруппирует category > 1
```

-- Второй пример с использованием HAVING
```
SELECT category AS "Id категории",
COUNT(*) AS "Количество товаров"
FROM goods
WHERE product_id >= 3
GROUP BY category
HAVING COUNT(*) > 1;
```

### CASE

* Простое выражение CASE – т.е. простое сравнение одного значения с набором других значений;
* Поисковое выражение CASE – в данном случае CASE содержит набор логических выражений,
которые вычисляются, чтобы вернуть результат.

Синтаксис CASE

```
CASE input_expression
   WHEN when_expression THEN result_expression [ ...n ]
 [ ELSE else_result_expression ]
END
```

Описание:

* input_expression – выражение, которое необходимо проверить;
* WHEN when_expression – выражение, с которым сравнивается input_expression;
* THEN result_expression – выражение, которое будет возвращено, если текущее условие выполняется;
* ELSE else_result_expression – дополнительный параметр, который предназначен для случаев,
когда ни одно из перечисленных в CASE условий не выполнилось. Это необязательный параметр.
Если ELSE не указано, а условия не выполнились, вернётся NULL.

Поисковое выражение CASE
```
CASE
   WHEN Boolean_expression THEN result_expression [ ...n ]
 [ ELSE else_result_expression ]
END
```

Описание:

* WHEN Boolean_expression – логическое выражение, которое служит для вычисления результата.
Это своего рода проверочное условие и таких условий может быть несколько;
* THEN result_expression – выражение, которое будет возвращено, если текущее условие выполняется;
* ELSE else_result_expression – дополнительный параметр, который предназначен для случаев,
когда ни одно из перечисленных в CASE условий не выполнилось. Это необязательный параметр.
Если ELSE не указано, а условия не выполнились, вернётся NULL.

-- Простое выражение CASE
```
SELECT product_id,
       CASE product_id WHEN 1 THEN 'Один'
                       WHEN 2 THEN 'Два'
                       WHEN 3 THEN 'Три'
                       WHEN 4 THEN 'Четыре'
                       WHEN 5 THEN 'Пять'
                       ELSE ''
       END AS IdText
FROM goods;
```

-- Пример поискового выражения CASE
```
SELECT product_id,
       CASE WHEN product_id = 1 THEN 'Один'
            WHEN product_id = 2 THEN 'Два'
            WHEN product_id = 3 THEN 'Три'
            WHEN product_id = 4 THEN 'Четыре'
            WHEN product_id = 5 THEN 'Пять'
            ELSE ''
       END AS IdText
FROM goods;
```

-- Пример более сложного поискового выражения CASE
```
SELECT product_id, price,
       CASE WHEN price  >  100 THEN 'Больше 100'
            WHEN price = 100 THEN 'Равно 100'
            WHEN price < 100 THEN 'Меньше 100'
            WHEN price = 50 AND product_id = 1 THEN 'Цена равна 50 и Id равен 1'
            ELSE 'Нет подходящего условия'
       END AS Result
FROM goods;
```

### SELECT INTO

***SELECT INTO – это инструкция языка SQL,
которая создает новую таблицу и вставляет в нее результирующие строки из SQL запроса.***

```
SELECT product_id, category, product_name, price
INTO test_table
FROM goods;

SELECT * FROM test_table;
```