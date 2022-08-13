CREATE TABLE company (
                         companyId int PRIMARY KEY,
                         companyName varchar(15),
                         companyCountry varchar(15)
);

CREATE TABLE phone (
                       phoneId SERIAL PRIMARY KEY,
                       phoneModel varchar(15),
                       companyId int,
                       FOREIGN KEY (companyId) REFERENCES company(companyId),
                       price int
);

INSERT INTO company VALUES (1, 'huawei', 'china');
INSERT INTO company VALUES (2, 'meizu', 'china');
INSERT INTO company VALUES (3, 'xiaomi', 'china');
INSERT INTO company VALUES (4, 'apple', 'usa');
INSERT INTO company VALUES (5, 'moto', 'usa');
INSERT INTO company VALUES (6, 'asdf', 'korea');

INSERT INTO phone (phoneModel, companyId, price) VALUES ('huawei_1', 1, 111);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('huawei_2', 1, 222);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('huawei_3', 1, 333);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('huawei_4', 1, 444);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('huawei_5', 1, 555);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('huawei_6', 1, 666);

INSERT INTO phone (phoneModel, companyId, price) VALUES ('meizu_1', 2, 222);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('meizu_2', 2, 222);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('meizu_3', 2, 333);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('meizu_4', 2, 444);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('meizu_5', 2, 555);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('meizu_6', 2, 666);

INSERT INTO phone (phoneModel, companyId, price) VALUES ('xiaomi_1', 3, 111);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('xiaomi_2', 3, 122);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('xiaomi_3', 3, 333);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('xiaomi_4', 3, 444);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('xiaomi_5', 3, 555);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('xiaomi_6', 3, 666);

INSERT INTO phone (phoneModel, companyId, price) VALUES ('apple_1', 4, 111);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('apple_2', 4, 133);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('apple_3', 4, 333);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('apple_4', 4, 444);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('apple_5', 4, 555);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('apple_6', 4, 666);

INSERT INTO phone (phoneModel, companyId, price) VALUES ('moto_1', 5, 111);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('moto_2', 5, 133);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('moto_3', 5, 333);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('moto_4', 5, 444);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('moto_5', 5, 555);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('moto_6', 5, 666);

INSERT INTO phone (phoneModel, companyId, price) VALUES ('asdf_1', 6, 111);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('asdf_2', 6, 122);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('asdf_3', 6, 333);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('asdf_4', 6, 444);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('asdf_5', 6, 555);
INSERT INTO phone (phoneModel, companyId, price) VALUES ('asdf_6', 6, 666);

SELECT * FROM company;
SELECT COUNT(companyId) FROM phone;


SELECT companyName, COUNT(phoneId), SUM(price)
FROM company
JOIN phone ON company.companyId = phone.companyId
GROUP BY companyName;

SELECT COUNT(phoneId)
FROM phone
JOIN company ON company.companyId = phone.companyId
WHERE company.companyCountry = 'china';

SELECT companyName
FROM company
JOIN phone ON company.companyId = phone.companyId
GROUP BY companyName
ORDER BY AVG(price) DESC
LIMIT 1;

SELECT phoneModel
FROM phone
WHERE price IN (SELECT MAX(price)
                FROM phone
                GROUP BY companyId);



SELECT MAX(price)
FROM company
JOIN phone ON company.companyId = phone.companyId
GROUP BY companyName;

SELECT companyId, MAX(price)
FROM phone
GROUP BY companyID;


SELECT phoneModel, price, companyName
FROM company, phone
WHERE company.companyId = phone.companyId
  AND phone.price = (
    SELECT MAX(phone.price)
    FROM phone
    WHERE company.companyId = phone.companyId
);
