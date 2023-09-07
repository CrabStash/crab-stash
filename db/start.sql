BEGIN TRANSACTION;

DEFINE TABLE user SCHEMAFULL;
DEFINE TABLE warehouse SCHEMAFULL;
DEFINE TABLE manages SCHEMAFULL;
DEFINE TABLE fields SCHEMALESS;
DEFINE TABLE entities SCHEMAFULL;
DEFINE TABLE categories SCHEMAFULL;

DEFINE FIELD email ON TABLE user TYPE string
  ASSERT $value != NULL AND $value != NONE AND is::email($value);
DEFINE FIELD firstName ON TABLE user TYPE string
  ASSERT $value != NULL AND $value != NONE AND string::len($value) > 0 AND is::alpha($value);
DEFINE FIELD lastName ON TABLE user TYPE string
  ASSERT $value != NULL AND $value != NONE AND string::len($value) > 0 AND is::alpha($value);
DEFINE FIELD passwd ON TABLE user TYPE string
  ASSERT $value != NULL AND $value != NONE AND string::len($value) > 0;
DEFINE FIELD avatar ON TABLE user TYPE string
  VALUE $value OR NULL
  ASSERT (is::url($value) OR $value == NULL);
DEFINE FIELD owns ON TABLE user TYPE array
  VALUE $value OR []
  ASSERT array::distinct($value) == $value OR $value == NULL;
DEFINE FIELD owns.* ON TABLE user TYPE record(warehouse);

DEFINE INDEX userEmailIndex ON TABLE user COLUMNS email UNIQUE;

DEFINE FIELD name ON TABLE warehouse TYPE string
  ASSERT $value != NULL AND $value != NONE AND string::len($value) > 0;
DEFINE FIELD desc ON TABLE warehouse TYPE string
  ASSERT $value != NULL AND $value != NONE AND string::len($value) > 0;
DEFINE FIELD logo ON TABLE warehouse TYPE string
  VALUE $value OR NULL
  ASSERT (is::url($value) OR $value == NULL);
DEFINE FIELD owner ON TABLE warehouse TYPE record(user)
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD isPhysical ON TABLE warehouse TYPE bool
  VALUE $value OR false;
DEFINE FIELD capacity ON TABLE warehouse TYPE int
  VALUE $value OR 0;

DEFINE INDEX name_owner_pairs ON TABLE warehouse COLUMNS name,owner UNIQUE;

DEFINE FIELD title ON TABLE categories TYPE object
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD title.pl ON TABLE categories type string
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD title.en ON TABLE categories type string
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD description ON TABLE categories TYPE object
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD description.pl ON TABLE categories type string
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD description.en ON TABLE categories type string
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD properties ON TABLE categories type array
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD properties.* ON TABLE categories type record(fields);
DEFINE FIELD parents ON TABLE categories type array;
DEFINE FIELD parents.* ON TABLE categories type record(categories);

DEFINE FIELD in ON TABLE manages TYPE record(user);
DEFINE FIELD out ON TABLE manages TYPE record(warehouse);
DEFINE FIELD role ON TABLE manages TYPE number
  ASSERT $value != NULL AND $value != NONE;

DEFINE INDEX unique_relationships ON TABLE manages COLUMNS in, out UNIQUE;

DEFINE FIELD data ON TABLE entities type object
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD meta ON TABLE entities type object
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD meta.status ON TABLE entities type string;
DEFINE FIELD meta.release ON TABLE entities type datetime;
DEFINE FIELD meta.discount ON TABLE entities type number;
DEFINE FIELD meta.discount_deadline ON TABLE entities type datetime;


COMMIT TRANSACTION;