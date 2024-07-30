CREATE TABLE "products" (
  "id" SERIAL PRIMARY KEY,
  "product_id" varchar UNIQUE,
  "product_name" varchar,
  "category" varchar,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "customers" (
  "id" SERIAL PRIMARY KEY,
  "customer_id" varchar UNIQUE,
  "customer_name" varchar,
  "email" varchar,
  "address" varchar,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "regions" (
  "id" SERIAL PRIMARY KEY,
  "region_name" varchar,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);



CREATE TABLE "orders" (
  "id" SERIAL PRIMARY KEY,
  "order_id" integer UNIQUE,
  "customer_id" varchar,
  "region_id" integer,
  "date_of_sale" timestamp,
  "payment_method" varchar,
  "shipping_cost" float,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "order_details" (
  "id" SERIAL PRIMARY KEY,
  "order_id" integer UNIQUE,
  "product_id" varchar,
  "quantity_sold" integer,
  "unit_price" float,
  "discount" float,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

ALTER TABLE "orders" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("customer_id");

ALTER TABLE "orders" ADD FOREIGN KEY ("region_id") REFERENCES "regions" ("id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("order_id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id");