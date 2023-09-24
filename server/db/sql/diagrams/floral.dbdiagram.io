Project floral {
  database_type: 'PostgreSQL'
  Note: 'Floral e-store aggregator'
}

Table floral.store {
  id integer [pk, increment]
  name varchar(100) [not null, unique]
  password varchar(500) [not null]
  email varchar(50) [not null, unique]
  phone_number varchar(16) [not null]
  created timestamp [not null]
}

// ALTER TABLE "floral"."store" ALTER COLUMN created SET DEFAULT NOW();

Table floral.user {
  id integer [pk, increment]
  first_name varchar(100) [not null]
  last_name varchar(100) [not null]
  email varchar(50) [not null, unique]
  password varchar(500) [not null]
  phone_number varchar(16) [not null, unique]
  created timestamp [not null]
}

// ALTER TABLE "floral"."user" ALTER COLUMN created SET DEFAULT NOW();

Table floral.product_category {
  id integer [pk, increment]
  name VARCHAR(60) [not null]
  description VARCHAR(150) [default: '']
  created timestamp [not null]
  modified timestamp [not null]
}

// ALTER TABLE "floral"."product_category" ALTER COLUMN created SET DEFAULT NOW();
// ALTER TABLE "floral"."product_category" ALTER COLUMN modified SET DEFAULT NOW();

Table floral.product {
  id integer [pk, increment]
  store_id integer [not null]
  name varchar(100) [not null]
  description varchar(1000) [default: '']
  image_url varchar(300) [not null]
  price integer [not null]
  created timestamp [not null]
  modified timestamp [not null]
  category integer [note: "id of the product's category"]
  
  indexes {
    store_id
  }
}

Ref: floral.product.store_id > floral.store.id [delete: no action, update: cascade]
Ref: floral.product.category > floral.product_category.id [delete: set null, update: cascade]

Table floral.product_common_traits {
  product_id integer [pk, not null]
  min_height smallint [note:'in cm']
  max_height smallint [note:'in cm']
  created timestamp [not null]
  modified timestamp [not null]
}

// ALTER TABLE "floral"."product_common_traits" ALTER COLUMN created SET DEFAULT NOW();
// ALTER TABLE "floral"."product_common_traits" ALTER COLUMN modified SET DEFAULT NOW();

Ref: floral.product.id - floral.product_common_traits.product_id [delete: no action, update: cascade]

Table floral.user_favorite {
  user_id integer [not null]
  product_id integer [not null]

  indexes {
    user_id
    product_id
    (user_id, product_id) [pk]
  }
}

Ref: floral.user_favorite.user_id > floral.user.id [delete: cascade, update: cascade]
Ref: floral.user_favorite.product_id > floral.product.id [delete: cascade, update: cascade]

Table floral.cart_position {
  user_id integer [not null]
  product_id integer [not null]
  quantity integer [not null]

  indexes {
    user_id
    product_id
    (user_id, product_id) [unique]
  }
}

Ref: floral.cart_position.user_id > floral.user.id [delete: cascade, update: cascade]
Ref: floral.cart_position.product_id > floral.product.id [delete: cascade, update: cascade]

Enum floral.order_status {
  created
  in_progress
  processed
  delivery
  canceled
  completed
}

Table floral.order {
  id integer [pk, increment]
  status floral.order_status
  user_id integer [not null]
  created timestamp [not null]
  status_modified timestamp [not null]

  indexes {
    user_id
  }
}

// ALTER TABLE "floral"."order" ALTER COLUMN created SET DEFAULT NOW();
// ALTER TABLE "floral"."order" ALTER COLUMN status_modified SET DEFAULT NOW();

Ref: floral.order.user_id > floral.user.id [delete: no action]

Table floral.order_position {
  order_id integer [not null]
  product_id integer [not null]
  quantity integer [not null]

  indexes {
    order_id
    product_id
    (product_id, order_id) [unique]
  }
}

Ref: floral.order_position.order_id > floral.order.id [delete: no action, update: cascade]
Ref: floral.order_position.product_id > floral.product.id [delete: no action, update: cascade]

Table floral.review {
  id bigint [pk, increment]
  user_id integer
  product_id integer [not null]
  rating float [not null]
  review_text VARCHAR(2500)
  created timestamp [not null]
  modified timestamp [not null]

  indexes {
    user_id
    product_id
    (user_id, product_id) [unique]
  }
}

// ALTER TABLE "floral"."review" ALTER COLUMN created SET DEFAULT NOW();
// ALTER TABLE "floral"."review" ALTER COLUMN modified SET DEFAULT NOW();

Ref: floral.review.user_id > floral.user.id [delete: set null, update: cascade]
Ref: floral.review.product_id > floral.product.id [delete: no action, update: cascade]
