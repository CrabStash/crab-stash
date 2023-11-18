-- ------------------------------
-- OPTION
-- ------------------------------

OPTION IMPORT;

-- ------------------------------
-- TABLE: categories
-- ------------------------------

DEFINE TABLE categories SCHEMAFULL;

DEFINE FIELD description ON categories TYPE string ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD parents ON categories TYPE array;
DEFINE FIELD parents[*] ON categories TYPE record<categories>;
DEFINE FIELD properties ON categories TYPE array ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD properties[*] ON categories TYPE record<fields>;
DEFINE FIELD title ON categories TYPE string ASSERT $value != NULL AND $value != NONE;

-- ------------------------------
-- TABLE: entities
-- ------------------------------

DEFINE TABLE entities SCHEMAFULL;

DEFINE FIELD data ON entities TYPE object ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD meta ON entities TYPE object ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD meta.discount ON entities TYPE number;
DEFINE FIELD meta.discount_deadline ON entities TYPE datetime;
DEFINE FIELD meta.release ON entities TYPE datetime;
DEFINE FIELD meta.status ON entities TYPE string;

-- ------------------------------
-- TABLE: fields
-- ------------------------------

DEFINE TABLE fields SCHEMAFULL;

DEFINE FIELD title ON TABLE fields type string
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD help ON TABLE fields type string
  ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD type ON TABLE fields type string
  ASSERT $value != NULL AND $value != NONE;

-- ------------------------------
-- TABLE: manages
-- ------------------------------

DEFINE TABLE manages SCHEMAFULL;

DEFINE FIELD in ON manages TYPE record<user>;
DEFINE FIELD out ON manages TYPE record<warehouse>;
DEFINE FIELD role ON manages TYPE number ASSERT $value != NULL AND $value != NONE;

DEFINE INDEX unique_relationships ON manages FIELDS in, out UNIQUE;

-- ------------------------------
-- TABLE: user
-- ------------------------------

DEFINE TABLE user SCHEMAFULL;

DEFINE FIELD avatar ON user TYPE string DEFAULT 'https://google.com' ASSERT string::is::url($value);
DEFINE FIELD email ON user TYPE string ASSERT $value != NULL AND $value != NONE AND string::is::email($value);
DEFINE FIELD firstName ON user TYPE string ASSERT $value != NULL AND $value != NONE AND string::len($value) > 0;
DEFINE FIELD lastName ON user TYPE string ASSERT $value != NULL AND $value != NONE AND string::len($value) > 0;
DEFINE FIELD owns ON user TYPE option<array> DEFAULT [] ASSERT array::distinct($value) == $value;
DEFINE FIELD owns[*] ON user TYPE record<warehouse>;
DEFINE FIELD passwd ON user TYPE string ASSERT $value != NULL AND $value != NONE AND string::len($value) > 0;

DEFINE INDEX userEmailIndex ON user FIELDS email UNIQUE;

-- ------------------------------
-- TABLE: warehouse
-- ------------------------------

DEFINE TABLE warehouse SCHEMAFULL;

DEFINE FIELD capacity ON warehouse TYPE int VALUE $value OR 0;
DEFINE FIELD desc ON warehouse TYPE string ASSERT $value != NULL AND $value != NONE AND string::len($value) > 0;
DEFINE FIELD isPhysical ON warehouse TYPE bool VALUE $value OR false;
DEFINE FIELD logo ON warehouse TYPE option<string> ASSERT string::is::url($value);
DEFINE FIELD name ON warehouse TYPE string ASSERT $value != NULL AND $value != NONE AND string::len($value) > 0;
DEFINE FIELD owner ON warehouse TYPE record<user> ASSERT $value != NULL AND $value != NONE;

DEFINE INDEX name_owner_pairs ON warehouse FIELDS name, owner UNIQUE;

-- ------------------------------
-- TRANSACTION
-- ------------------------------

BEGIN TRANSACTION;

-- ------------------------------
-- TABLE DATA: categories
-- ------------------------------


-- ------------------------------
-- TABLE DATA: entities
-- ------------------------------


-- ------------------------------
-- TABLE DATA: fields
-- ------------------------------


-- ------------------------------
-- TABLE DATA: manages
-- ------------------------------

RELATE user:018b1e6973bd728b854f1079638eca85 -> manages:1ws8pexlupfunmzonhs5 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:1ws8pexlupfunmzonhs5, in: user:018b1e6973bd728b854f1079638eca85, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e6a80e17ad292be6835534c2cda -> manages:4ctpxhsuotvptd3prkjp -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:4ctpxhsuotvptd3prkjp, in: user:018b1e6a80e17ad292be6835534c2cda, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e6acd50736a82010fc4a1f09c38 -> manages:5e94pmu9257sv4nqmvlr -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:5e94pmu9257sv4nqmvlr, in: user:018b1e6acd50736a82010fc4a1f09c38, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e699bcb71febe551f627c594545 -> manages:7wwmgvsb9nn6qnqy6lop -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:7wwmgvsb9nn6qnqy6lop, in: user:018b1e699bcb71febe551f627c594545, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e6936c079798e154826f1803d01 -> manages:8rycaz7i1nlh74k1hsf7 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:8rycaz7i1nlh74k1hsf7, in: user:018b1e6936c079798e154826f1803d01, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 4 };
RELATE user:018b1e6a05067ea282e215caca3d3f75 -> manages:90wumjxdfwf7i8fvdrhk -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:90wumjxdfwf7i8fvdrhk, in: user:018b1e6a05067ea282e215caca3d3f75, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e6a579d7f11be9db5253c0a4b4f -> manages:9bb67j3zdd2tm7p4rwvt -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:9bb67j3zdd2tm7p4rwvt, in: user:018b1e6a579d7f11be9db5253c0a4b4f, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e6adde77d2abcc700ebf5772087 -> manages:b9l5xegemed0zb3eg0gu -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:b9l5xegemed0zb3eg0gu, in: user:018b1e6adde77d2abcc700ebf5772087, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e6aa8467b19b0537a4999b02861 -> manages:bbbhfi04maqivuurdixk -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:bbbhfi04maqivuurdixk, in: user:018b1e6aa8467b19b0537a4999b02861, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e6987247353abb9c8e8750fedde -> manages:bwgvemfj0aps0wu8qo9o -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:bwgvemfj0aps0wu8qo9o, in: user:018b1e6987247353abb9c8e8750fedde, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e69eff87266a1f2089e1db4e355 -> manages:ckh24zg2w3a454czkr7x -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:ckh24zg2w3a454czkr7x, in: user:018b1e69eff87266a1f2089e1db4e355, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e6a94dc794b87bf589455395d1b -> manages:ckjfer0kknax5fztkfu6 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:ckjfer0kknax5fztkfu6, in: user:018b1e6a94dc794b87bf589455395d1b, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e6a6dbc7d6288144ff2e80a4ef4 -> manages:fw0zvzshw7fa1qrmpt69 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:fw0zvzshw7fa1qrmpt69, in: user:018b1e6a6dbc7d6288144ff2e80a4ef4, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e695dc67d9c9dc8eed714174b64 -> manages:ib8lnz6dyrnc13dy9fsh -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:ib8lnz6dyrnc13dy9fsh, in: user:018b1e695dc67d9c9dc8eed714174b64, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e6abba374eab462de7da1868da0 -> manages:ioyjaj4k9wrd386erq9t -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:ioyjaj4k9wrd386erq9t, in: user:018b1e6abba374eab462de7da1868da0, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e6afdb17ba68922411e07888442 -> manages:j13bfbz5pkwog5f9yfaf -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:j13bfbz5pkwog5f9yfaf, in: user:018b1e6afdb17ba68922411e07888442, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e69da3b7a5a95dd10ddae74dd4d -> manages:j2qaoowzww1z3pb3xvz0 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:j2qaoowzww1z3pb3xvz0, in: user:018b1e69da3b7a5a95dd10ddae74dd4d, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e69c7627a179889cc1081c25f0a -> manages:k4w1d36wirnql49b7go7 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:k4w1d36wirnql49b7go7, in: user:018b1e69c7627a179889cc1081c25f0a, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };
RELATE user:018b1e6a3d5f7876a3e08e8a54c062b8 -> manages:rbvefv7bkbeoh22t8ixh -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: manages:rbvefv7bkbeoh22t8ixh, in: user:018b1e6a3d5f7876a3e08e8a54c062b8, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, role: 0 };

-- ------------------------------
-- TABLE DATA: user
-- ------------------------------

UPDATE user:018b1e6936c079798e154826f1803d01 CONTENT { avatar: 'https://google.com', email: 'rbastiman0@imgur.com', firstName: 'Rozelle', id: user:018b1e6936c079798e154826f1803d01, lastName: 'Bastiman', owns: [warehouse:018b1e6ca3397d8780fb21315ccbfbc1], passwd: '$2a$10$1lnRULuB/GI0d5.v7pwMKehDkkWTPz7tNT8G0eN2NYBjBDWEHvG3e' };
UPDATE user:018b1e695dc67d9c9dc8eed714174b64 CONTENT { avatar: 'https://google.com', email: 'bbernakiewicz1@123reg.co.uk', firstName: 'Boniface', id: user:018b1e695dc67d9c9dc8eed714174b64, lastName: 'Bernakiewicz', owns: [], passwd: '$2a$10$2HTJUW/6QKn57uyUJmo9X.ejTtnrxRvDh.oqmzkrmrVLRz1O2PZpC' };
UPDATE user:018b1e6973bd728b854f1079638eca85 CONTENT { avatar: 'https://google.com', email: 'bmcgarva2@sogou.com', firstName: 'Brittan', id: user:018b1e6973bd728b854f1079638eca85, lastName: 'McGarva', owns: [], passwd: '$2a$10$N5dLDoO4ggnI4s3EmscXzucITA9gJ2dHrubnZWd6AZwo.or2WWAMu' };
UPDATE user:018b1e6987247353abb9c8e8750fedde CONTENT { avatar: 'https://google.com', email: 'iping3@i2i.jp', firstName: 'Ilysa', id: user:018b1e6987247353abb9c8e8750fedde, lastName: 'Ping', owns: [], passwd: '$2a$10$VAnk0MAPw5NLOngTGOsRHOwKjg/zJX4r1dlOkBKy7GbQhmrg9VHxO' };
UPDATE user:018b1e699bcb71febe551f627c594545 CONTENT { avatar: 'https://google.com', email: 'hinglefield4@macromedia.com', firstName: 'Halimeda', id: user:018b1e699bcb71febe551f627c594545, lastName: 'Inglefield', owns: [], passwd: '$2a$10$oM722n2iSk/Jjwlb5VYj2Oaw1FwCSfiw4sh48wgWuclA4rkH8MGly' };
UPDATE user:018b1e69c7627a179889cc1081c25f0a CONTENT { avatar: 'https://google.com', email: 'egeratasch5@accuweather.com', firstName: 'Earvin', id: user:018b1e69c7627a179889cc1081c25f0a, lastName: 'Geratasch', owns: [], passwd: '$2a$10$E0tOO2UwvatARldtwqXoI.iWRexLawph26Sw.k2ODDB13fva7mKX2' };
UPDATE user:018b1e69da3b7a5a95dd10ddae74dd4d CONTENT { avatar: 'https://google.com', email: 'kibeson6@sciencedirect.com', firstName: 'Kim', id: user:018b1e69da3b7a5a95dd10ddae74dd4d, lastName: 'Ibeson', owns: [], passwd: '$2a$10$z8qUREppK/gn7G3KCM88murG.YWUdfPAXRxVfP7PtqwDi7ySVbxeq' };
UPDATE user:018b1e69eff87266a1f2089e1db4e355 CONTENT { avatar: 'https://google.com', email: 'lmcilmorie7@ezinearticles.com', firstName: 'Lillian', id: user:018b1e69eff87266a1f2089e1db4e355, lastName: 'McIlmorie', owns: [], passwd: '$2a$10$XTsBU6WZDquEV21y5d.Im.jUmr3ed21Tb6YroxfIgl/vu/IjkEbim' };
UPDATE user:018b1e6a05067ea282e215caca3d3f75 CONTENT { avatar: 'https://google.com', email: 'jsnaddin8@unesco.org', firstName: 'Jean', id: user:018b1e6a05067ea282e215caca3d3f75, lastName: 'Snaddin', owns: [], passwd: '$2a$10$R221jA4x7FMOmnCOgSfMMO6W2G22x5IcbwWwnRsVvDx8hJTGYJ5u.' };
UPDATE user:018b1e6a3d5f7876a3e08e8a54c062b8 CONTENT { avatar: 'https://google.com', email: 'mdumbartona@youku.com', firstName: 'Mattie', id: user:018b1e6a3d5f7876a3e08e8a54c062b8, lastName: 'Dumbarton', owns: [], passwd: '$2a$10$Yxvk7Lp1KrvvYr1hmTqNE.j0/0/U0h1ZXjULurzfRlPiEXO6DQlyC' };
UPDATE user:018b1e6a579d7f11be9db5253c0a4b4f CONTENT { avatar: 'https://google.com', email: 'sdevilb@examiner.com', firstName: 'Siegfried', id: user:018b1e6a579d7f11be9db5253c0a4b4f, lastName: 'Devil', owns: [], passwd: '$2a$10$Bgjh0bI80nMzWA/XR/spXOWvJEzvQtJZhd0/A7vfevF2pH35F3u2O' };
UPDATE user:018b1e6a6dbc7d6288144ff2e80a4ef4 CONTENT { avatar: 'https://google.com', email: 'kondrakc@google.co.jp', firstName: 'Kippy', id: user:018b1e6a6dbc7d6288144ff2e80a4ef4, lastName: 'Ondrak', owns: [], passwd: '$2a$10$ZekTtJ/g6vgjqryW0DCkceOuBfZwlrm5y0gUjskPB26eNDyWSXsHm' };
UPDATE user:018b1e6a80e17ad292be6835534c2cda CONTENT { avatar: 'https://google.com', email: 'huttingd@twitter.com', firstName: 'Haley', id: user:018b1e6a80e17ad292be6835534c2cda, lastName: 'Utting', owns: [], passwd: '$2a$10$8lWy0hJrBR5UnKdXc6UBB.QPuDMIKE1vJ7//xFMnmRZ40kw53KPbq' };
UPDATE user:018b1e6a94dc794b87bf589455395d1b CONTENT { avatar: 'https://google.com', email: 'aeversleye@hugedomains.com', firstName: 'Annamarie', id: user:018b1e6a94dc794b87bf589455395d1b, lastName: 'Eversley', owns: [], passwd: '$2a$10$77WBmuE1Njf2JB/N4JGlnePbEqT8VRA4gf6.E/OULS.PL.CDmK0rS' };
UPDATE user:018b1e6aa8467b19b0537a4999b02861 CONTENT { avatar: 'https://google.com', email: 'bswafieldf@cbsnews.com', firstName: 'Berkie', id: user:018b1e6aa8467b19b0537a4999b02861, lastName: 'Swafield', owns: [], passwd: '$2a$10$NFNvOIrzN6jCv/6FZ7CnzugoMnvx9GlpFmptJClRbk2s3x5G32M8e' };
UPDATE user:018b1e6abba374eab462de7da1868da0 CONTENT { avatar: 'https://google.com', email: 'gmordeyg@is.gd', firstName: 'Gibby', id: user:018b1e6abba374eab462de7da1868da0, lastName: 'Mordey', owns: [], passwd: '$2a$10$Svm9GRlPy/MeIpxz2BgOveqRGdwFxYfIATEZlqRke74mU1sra.w0S' };
UPDATE user:018b1e6acd50736a82010fc4a1f09c38 CONTENT { avatar: 'https://google.com', email: 'rforgieh@narod.ru', firstName: 'Raphael', id: user:018b1e6acd50736a82010fc4a1f09c38, lastName: 'Forgie', owns: [], passwd: '$2a$10$VOh1TwY4ztiFH.F6vmufn.kShR7Z8iHB4htQss8btpyJWkuORIqbm' };
UPDATE user:018b1e6adde77d2abcc700ebf5772087 CONTENT { avatar: 'https://google.com', email: 'lelgari@goo.gl', firstName: 'Land', id: user:018b1e6adde77d2abcc700ebf5772087, lastName: 'Elgar', owns: [], passwd: '$2a$10$6u.jP6B2jHOk.IUZmGNZ5O6DJQNw2tLDSplx9HE9857X9Ju1tLKq2' };
UPDATE user:018b1e6afdb17ba68922411e07888442 CONTENT { avatar: 'https://google.com', email: 'ssemeradj@gnu.org', firstName: 'Shayne', id: user:018b1e6afdb17ba68922411e07888442, lastName: 'Semerad', owns: [], passwd: '$2a$10$p63ygg0ZI/XvRgE1b84T5eA2urgNlJzVDgSUB9Y5dtI4abM89zkJi' };

-- ------------------------------
-- TABLE DATA: warehouse
-- ------------------------------

UPDATE warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { capacity: -1, desc: 'najmaśniejszy magazyn po tej stronie wisły', id: warehouse:018b1e6ca3397d8780fb21315ccbfbc1, isPhysical: false, logo: 'https://google.com', name: 'Masny magazyn', owner: user:018b1e6936c079798e154826f1803d01 };

-- ------------------------------
-- TRANSACTION
-- ------------------------------

COMMIT TRANSACTION;

