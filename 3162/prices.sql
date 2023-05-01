-- CREATE TABLE stock_prices(
--     symbol text not null,
--     quote_date date not null,
--     price numeric(10,3) not null
-- );
DROP TABLE stock_prices;

CREATE TABLE stock_prices As
(SELECT symbol, CURRENT_DATE AS quote_date, 20.150 As price
FROM my_stocks);
