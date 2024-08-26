CREATE DOMAIN PRODUCT_NAME AS VARCHAR(20);
CREATE SEQUENCE order_name_seq AS integer;

CREATE TABLE product (
    product_name PRODUCT_NAME NOT NULL PRIMARY KEY
);

CREATE TABLE product_workflow_1 (
    product_name PRODUCT_NAME NOT NULL PRIMARY KEY,
    worfkflow_property_1 VARCHAR(100),
    FOREIGN KEY (product_name) REFERENCES product(product_name)
);

CREATE TABLE orders (
    order_id integer NOT NULL DEFAULT nextval('order_name_seq') PRIMARY KEY,
    order_name VARCHAR(20) NOT NULL,
    order_description VARCHAR(100),
    product PRODUCT_NAME NOT NULL,
    FOREIGN KEY (product) REFERENCES product(product_name)
);