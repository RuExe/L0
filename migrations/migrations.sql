CREATE TABLE deliveries
(
    id      serial PRIMARY KEY,
    name    varchar not null,
    phone   varchar not null,
    zip     varchar not null,
    city    varchar not null,
    address varchar not null,
    region  varchar not null,
    email   varchar not null
);

CREATE TABLE orders
(
    order_uid          varchar PRIMARY KEY,
    track_number       varchar     not null UNIQUE,
    entry              varchar(30) not null,
    delivery_id        int         not null REFERENCES deliveries (id),
    locale             varchar(10) not null,
    internal_signature varchar     not null,
    customer_id        varchar     not null,
    delivery_service   varchar     not null,
    shardkey           varchar     not null,
    sm_id              smallint    not null,
    date_created       timestamp   not null,
    oof_shard          varchar     not null
);

CREATE TABLE payments
(
    transaction   varchar PRIMARY KEY REFERENCES orders (order_uid),
    request_id    varchar not null,
    currency      varchar not null,
    provider      varchar not null,
    amount        int     not null,
    payment_dt    int     not null,
    bank          varchar not null,
    delivery_cost int     not null,
    goods_total   int     not null,
    custom_fee    int     not null
);

CREATE TABLE items
(
    rid          varchar PRIMARY KEY,
    chrt_id      int     not null,
    track_number varchar not null REFERENCES orders (track_number),
    price        int     not null,
    name         varchar not null,
    sale         int     not null,
    size         varchar not null,
    total_price  int     not null,
    nm_id        int     not null,
    brand        varchar not null,
    status       int     not null
);