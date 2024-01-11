BEGIN TRANSACTION;

DEFINE TABLE user SCHEMAFULL;
DEFINE TABLE warehouse SCHEMAFULL;
DEFINE TABLE manages SCHEMAFULL;
DEFINE TABLE categories SCHEMAFULL;
DEFINE TABLE fields SCHEMAFULL;
DEFINE TABLE entities SCHEMALESS;
DEFINE TABLE categories_to_warehouses SCHEMAFULL;
DEFINE TABLE fields_to_warehouses SCHEMAFULL;
DEFINE TABLE entities_to_categories SCHEMAFULL;

DEFINE FIELD email ON TABLE user TYPE string
  ASSERT $value != NULL AND $value != NONE AND string::is::email($value);
DEFINE FIELD firstName ON TABLE user TYPE string
  ASSERT $value != NULL AND $value != NONE AND string::len($value) > 0;
DEFINE FIELD lastName ON TABLE user TYPE string
  ASSERT $value != NULL AND $value != NONE AND string::len($value) > 0;
DEFINE FIELD passwd ON TABLE user TYPE string
  ASSERT $value != NULL AND $value != NONE AND string::len($value) > 0;
DEFINE FIELD avatar ON TABLE user TYPE option<string>
  ASSERT $value == NONE || string::is::url($value);
DEFINE FIELD default_warehouse ON TABLE user TYPE option<record(warehouse)>;
DEFINE FIELD owns ON TABLE user TYPE option<array>
  DEFAULT []
  ASSERT array::distinct($value) == $value;
DEFINE FIELD owns.* ON TABLE user TYPE record(warehouse);

DEFINE INDEX userEmailIndex ON TABLE user COLUMNS email UNIQUE;

DEFINE FIELD name ON TABLE warehouse TYPE string
  ASSERT $value != NULL AND $value != NONE AND string::len($value) > 0;
DEFINE FIELD desc ON TABLE warehouse TYPE string
  ASSERT $value != NULL AND $value != NONE AND string::len($value) > 0;
DEFINE FIELD logo ON TABLE warehouse TYPE option<string>
  ASSERT string::is::url($value);
DEFINE FIELD owner ON TABLE warehouse TYPE record(user)
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD isPhysical ON TABLE warehouse TYPE bool
  VALUE $value OR false;
DEFINE FIELD capacity ON TABLE warehouse TYPE int
  VALUE $value OR 0;

DEFINE INDEX name_owner_pairs ON TABLE warehouse COLUMNS name,owner UNIQUE;

DEFINE FIELD title ON TABLE categories TYPE string
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD description ON TABLE categories TYPE string
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD properties ON TABLE categories type array
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD properties.* ON TABLE categories type record(fields);
DEFINE FIELD parents ON TABLE categories type array
DEFAULT [];
DEFINE FIELD parents.* ON TABLE categories type record(categories);

DEFINE FIELD in ON TABLE manages TYPE record(user);
DEFINE FIELD out ON TABLE manages TYPE record(warehouse);
DEFINE FIELD role ON TABLE manages TYPE number
  ASSERT $value != NULL AND $value != NONE;

DEFINE INDEX unique_relationships ON TABLE manages COLUMNS in, out UNIQUE;

DEFINE FIELD in ON TABLE categories_to_warehouses type record(categories);
DEFINE FIELD out ON TABLE categories_to_warehouses type record(warehouse);
DEFINE INDEX unique_relationships ON TABLE categories_to_warehouses COLUMNS in, out UNIQUE;

DEFINE FIELD in ON TABLE fields_to_warehouses type record(fields);
DEFINE FIELD out ON TABLE fields_to_warehouses type record(warehouse);
DEFINE INDEX unique_relationships ON TABLE fields_to_warehouses COLUMNS in, out UNIQUE;

DEFINE FIELD in ON TABLE entities_to_categories type record(entities);
DEFINE FIELD out ON TABLE entities_to_categories type record(categories);
DEFINE INDEX unique_relationships ON TABLE entities_to_categories COLUMNS in, out UNIQUE;


DEFINE FIELD title ON TABLE fields type string
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD help ON TABLE fields type string
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD type ON TABLE fields type string
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD format ON TABLE fields TYPE option<string>;

DELETE user, warehouse, manages, categories, fields, entities;

COMMIT TRANSACTION;
