CREATE TABLE users (
                         id int PRIMARY KEY,
                         name varchar(15),
                         surname varchar(15)
);

CREATE TABLE items (
                       sku int PRIMARY KEY,
                       item varchar(15),
                       user_id int,
                       FOREIGN KEY (user_id) REFERENCES users(id) ,
                       data date
);

INSERT INTO users VALUES (1, 'dasha', 'petrova');
INSERT INTO users VALUES (2, 'masha', 'volosatova');
INSERT INTO users VALUES (3, 'petya', 'kusakin');

INSERT INTO items VALUES (234532, 'meizu_1', 1, '2021-01-01');
INSERT INTO items VALUES (233452, 'meizu_1', 2, '2021-01-01');
INSERT INTO items VALUES (463343, 'meizu_1', 3, '2021-01-01');
INSERT INTO items VALUES (353533, 'meizu_1', 3, '2021-01-01');
INSERT INTO items VALUES (578567, 'meizu_1', 3, '2021-01-01');
INSERT INTO items VALUES (242343, 'meizu_1', 2, '2021-01-01');

SELECT user_id, COUNT(user_id)
FROM items
WHERE data BETWEEN '2021-02-01' AND '2021-02-28'
GROUP BY user_id

SELECT item, name, surname, data FROM users
    JOIN items ON users.id = items.user_id

SELECT name FORM Passeger
    WHERE LENGTH(name) = (SELECT MAX(LENGTH(name)) FROM Passeger)

SELECT town_to FROM Trip
WHERE passenger IN (SELECT trip
                    FROM Pass_in_trip
                             JOIN Passenger ON Pass_in_trip.passenger = Passenger.id
                    WHERE name = 'Bruce Willis')

SELECT time_in FROM Trip
WHERE town_to = 'London'
  AND id IN (SELECT trip
             FROM Pass_in_trip
                      JOIN Passenger ON Pass_in_trip.passenger = Passenger.id
             WHERE name = 'Steve Martin')