CREATE TABLE IF NOT EXISTS orders
(
    order_uid  varchar unique primary key not null ,
    track_number varchar,
    entry varchar,
    locale varchar,
    internal_signature varchar,
    customer_id varchar,
    delivery_service varchar,
    shardkey varchar,
    sm_id integer,
    date_created varchar,
    oof_shard varchar
);

CREATE TABLE IF NOT EXISTS deliveries
(
    order_uid varchar not null ,
    name varchar primary key not null ,
    phone varchar,
    zip varchar,
    city varchar,
    address varchar,
    region varchar,
    email varchar
);

CREATE TABLE IF NOT EXISTS items (
    chrt_id integer primary key not null ,
    track_number varchar not null ,
    price integer,
    rid varchar,
    name varchar,
    sale integer,
    size varchar,
    total_price integer,
    nm_id integer,
    brand varchar,
    status integer
);

CREATE TABLE IF NOT EXISTS payments (
    transaction varchar not null ,
    request_id varchar primary key not null ,
    currency varchar,
    provider varchar,
    amount integer,
    payment_dt integer,
    bank varchar,
    delivery_cost integer,
    goods_total integer,
    custom_fee integer
);
