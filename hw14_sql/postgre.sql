create database hw14_sql;
create schema if not exists hw14_sql;

create table hw14_sql.users(
    id serial primary key,
    name text not null ,
    email text unique ,
    password text not null
);

create table hw14_sql.orders(
    id serial primary key ,
    user_id serial references hw14_sql.users(id) on delete cascade ,
    order_date timestamp default current_timestamp,
    total_amount decimal(10, 2) not null
);

create table hw14_sql.products(
    id serial primary key ,
    name text not null ,
    price decimal(10, 2) not null
);

create table hw14_sql.orderProducts(
    order_id serial not null,
    product_id serial not null,
    primary key (order_id, product_id),
    foreign key (order_id) references hw14_sql.orders(id) on delete cascade,
    foreign key (product_id) references hw14_sql.products(id) on delete cascade
);

insert into hw14_sql.users(name, email, password) VALUES ('Daniil', 'daniil@yandex.ru', 'qwerty');
insert into hw14_sql.users(name, email, password) VALUES ('Irina', 'irina@gmail.com', 'qwerty1234');
insert into hw14_sql.users(name, email, password) VALUES ('Alexander', 'alex@gmail.com', '12345');
insert into hw14_sql.users(name, email, password) VALUES ('Tatyana', 'tatyana@mail.ru', '098765');

insert into hw14_sql.products(name, price) VALUES ('TV Xiomi 42`', 34990);
insert into hw14_sql.products(name, price) VALUES ('TV Samsung QLED 65`', 119990);
insert into hw14_sql.products(name, price) VALUES ('TV LG C1 OLED 85`', 299990);

update hw14_sql.users set name='Oleg' where name='Tatyana';
update hw14_sql.users set email='oleg@mail.ru' where name='Oleg';

delete from hw14_sql.users where name='Oleg';

update hw14_sql.products set price=199990 where name='TV LG C1 OLED 85`';

insert into hw14_sql.orders(user_id, order_date, total_amount) VALUES (1, current_timestamp, 119990);
insert into hw14_sql.orders(user_id, order_date, total_amount) VALUES (1, current_timestamp, 34990);

delete from hw14_sql.orders where total_amount=34990;

select * from hw14_sql.users;
select * from hw14_sql.users where email like '%@gmail.com';

select * from hw14_sql.orders;
select * from hw14_sql.orders where user_id=1;

select u.id AS user_id,
SUM(o.total_amount) AS total_order_amount,
AVG(o.total_amount) AS average_product_price
from hw14_sql.users u
left join hw14_sql.orders o on u.id = o.user_id
where u.id=1
group by u.id;

create index idx_users_email on hw14_sql.users(email);
create index idx_orders_user_id on hw14_sql.orders(user_id);
create index idx_product_name on hw14_sql.products(name);