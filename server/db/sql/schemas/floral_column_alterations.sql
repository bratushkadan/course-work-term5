ALTER TABLE
  "floral"."store"
ALTER COLUMN
  created
SET
  DEFAULT NOW();

ALTER TABLE
  "floral"."user"
ALTER COLUMN
  created
SET
  DEFAULT NOW();

ALTER TABLE
  "floral"."product_category"
ALTER COLUMN
  created
SET
  DEFAULT NOW();

ALTER TABLE
  "floral"."product_category"
ALTER COLUMN
  modified
SET
  DEFAULT NOW();

ALTER TABLE
  "floral"."product_common_traits"
ALTER COLUMN
  created
SET
  DEFAULT NOW();

ALTER TABLE
  "floral"."product_common_traits"
ALTER COLUMN
  modified
SET
  DEFAULT NOW();

ALTER TABLE
  "floral"."order"
ALTER COLUMN
  created
SET
  DEFAULT NOW();

ALTER TABLE
  "floral"."order"
ALTER COLUMN
  status_modified
SET
  DEFAULT NOW();

ALTER TABLE
  "floral"."review"
ALTER COLUMN
  created
SET
  DEFAULT NOW();

ALTER TABLE
  "floral"."review"
ALTER COLUMN
  modified
SET
  DEFAULT NOW();