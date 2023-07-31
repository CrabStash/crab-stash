BEGIN TRANSACTION;

DEFINE TABLE user SCHEMAFULL;
DEFINE TABLE warehouse SCHEMAFULL;
DEFINE TABLE manages SCHEMAFULL;
DEFINE TABLE fields SCHEMALESS;
DEFINE TABLE entities SCHEMALESS;
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
DEFINE FIELD roles ON TABLE manages TYPE array
  ASSERT $value != NULL AND $value != NONE AND array::distinct($value) == $value;
DEFINE FIELD roles.* ON TABLE manages TYPE string;

DEFINE INDEX unique_relationships ON TABLE manages COLUMNS in, out UNIQUE;

COMMIT TRANSACTION;

CREATE user CONTENT {
  email: "omikolajczak@edu.cdv.pl",
  firstName: "Oskar",
  lastName: "Mikołajczak",
  passwd: "uwu",
  avatar: null,
  owns: []
};

CREATE user CONTENT {
  email: "omikolajczak@edu.cdv.eu",
  lastName: "Mikołajczak",
  passwd: "uwu",
  avatar: null,
  owns: []
};

CREATE user CONTENT {
  email: "omikolajczak@edu.cdv.com",
  firstName: "",
  lastName: "Mikołajczak",
  passwd: "uwu",
  avatar: null,
  owns: []
};

CREATE warehouse CONTENT {
  name: "Masny magazyn",
  desc: "Najmaśniejszy magazyn po tej stronie wisły",
  logo: null,
  owner: "user:jnnh8c8ruxd18epww4p3",
  fields: [],
  categories: [],
  entities: []
}

UPDATE user:jnnh8c8ruxd18epww4p3 MERGE {
  owns: ["warehouse:a0hc4m47w6a4oxoi0pow", "warehouse:s4ts7y4fzghu93yzkoni", "masno"]
};

RELATE user:jnnh8c8ruxd18epww4p3 -> manages -> warehouse:a0hc4m47w6a4oxoi0pow
  SET roles = ["owner"];


BEGIN TRANSACTION;
  
COMMIT TRANSACTION;
