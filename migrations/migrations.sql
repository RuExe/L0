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
    id                 serial PRIMARY KEY,
    track_number       varchar     not null,
    entry              varchar(30) not null,
    delivery           varchar     not null,
--     delivery_id        int REFERENCES deliveries (id),
    payment            varchar     not null,
    items              varchar     not null,
    locale             varchar(10) not null,
    internal_signature varchar     not null,
    customer_id        varchar     not null,
    delivery_service   varchar     not null,
    shardkey           varchar     not null,
    sm_id              smallint    not null,
    date_created       timestamp   not null,
    oof_shard          varchar     not null
)