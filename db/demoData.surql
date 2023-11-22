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
-- TABLE: categories_to_warehouses
-- ------------------------------

DEFINE TABLE categories_to_warehouses SCHEMAFULL;

DEFINE FIELD in ON categories_to_warehouses TYPE record<categories>;
DEFINE FIELD out ON categories_to_warehouses TYPE record<warehouse>;

DEFINE INDEX unique_relationships ON categories_to_warehouses FIELDS in, out UNIQUE;

-- ------------------------------
-- TABLE: entities
-- ------------------------------

DEFINE TABLE entities SCHEMALESS;

-- ------------------------------
-- TABLE: entities_to_warehouses
-- ------------------------------

DEFINE TABLE entities_to_warehouses SCHEMAFULL;

DEFINE FIELD in ON entities_to_warehouses TYPE record<entities>;
DEFINE FIELD out ON entities_to_warehouses TYPE record<warehouse>;

DEFINE INDEX unique_relationships ON entities_to_warehouses FIELDS in, out UNIQUE;

-- ------------------------------
-- TABLE: fields
-- ------------------------------

DEFINE TABLE fields SCHEMAFULL;

DEFINE FIELD help ON fields TYPE string ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD title ON fields TYPE string ASSERT $value != NULL AND $value != NONE;
DEFINE FIELD type ON fields TYPE string ASSERT $value != NULL AND $value != NO;

-- ------------------------------
-- TABLE: fields_to_warehouses
-- ------------------------------

DEFINE TABLE fields_to_warehouses SCHEMAFULL;

DEFINE FIELD in ON fields_to_warehouses TYPE record<fields>;
DEFINE FIELD out ON fields_to_warehouses TYPE record<warehouse>;

DEFINE INDEX unique_relationships ON fields_to_warehouses FIELDS in, out UNIQUE;

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
DEFINE FIELD logo ON warehouse TYPE string ASSERT string::is::url($value);
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

UPDATE categories:55ldvropjbij4ofyrr6r CONTENT { description: 'A', id: categories:55ldvropjbij4ofyrr6r, parents: [], properties: [fields:19cnkznj77b09zuqmugc], title: 'kategoria A' };
UPDATE categories:anu09x36m6ctj1al1bhm CONTENT { description: 'ABA', id: categories:anu09x36m6ctj1al1bhm, parents: [categories:55ldvropjbij4ofyrr6r, categories:rkysc4zkbckj1558zzcn], properties: [fields:4816rl13ymz2k5lny7r1], title: 'kategoria ABA' };
UPDATE categories:jjl4y7ws2keh68nalonj CONTENT { description: 'AA', id: categories:jjl4y7ws2keh68nalonj, parents: [categories:55ldvropjbij4ofyrr6r], properties: [fields:ld7inkmf9tg5vaf9dwrs, fields:36a9tmke5hnuy5h2pczu], title: 'kategoria AA' };
UPDATE categories:rkysc4zkbckj1558zzcn CONTENT { description: 'AB', id: categories:rkysc4zkbckj1558zzcn, parents: [categories:55ldvropjbij4ofyrr6r], properties: [fields:gsnulwrqa7s3uzfx8uwn], title: 'kategoria AB' };
UPDATE categories:w2dnetzjs44z1l7kfo66 CONTENT { description: 'AAA', id: categories:w2dnetzjs44z1l7kfo66, parents: [categories:55ldvropjbij4ofyrr6r, categories:jjl4y7ws2keh68nalonj], properties: [fields:zeiqo8oe5don0sk7sjmi], title: 'kategoria AAA' };

-- ------------------------------
-- TABLE DATA: categories_to_warehouses
-- ------------------------------

RELATE categories:rkysc4zkbckj1558zzcn -> categories_to_warehouses:6fxjl6cms5g8do3ysze3 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: categories_to_warehouses:6fxjl6cms5g8do3ysze3, in: categories:rkysc4zkbckj1558zzcn, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE categories:w2dnetzjs44z1l7kfo66 -> categories_to_warehouses:c7vjqc0ckrtbjadjhe06 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: categories_to_warehouses:c7vjqc0ckrtbjadjhe06, in: categories:w2dnetzjs44z1l7kfo66, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE categories:jjl4y7ws2keh68nalonj -> categories_to_warehouses:ef7fk5g7vxh2nv9o6qdp -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: categories_to_warehouses:ef7fk5g7vxh2nv9o6qdp, in: categories:jjl4y7ws2keh68nalonj, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE categories:anu09x36m6ctj1al1bhm -> categories_to_warehouses:g61t4kv01cf3hiqsdik3 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: categories_to_warehouses:g61t4kv01cf3hiqsdik3, in: categories:anu09x36m6ctj1al1bhm, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE categories:55ldvropjbij4ofyrr6r -> categories_to_warehouses:szkfc1r55herfearh8ro -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: categories_to_warehouses:szkfc1r55herfearh8ro, in: categories:55ldvropjbij4ofyrr6r, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };

-- ------------------------------
-- TABLE DATA: entities
-- ------------------------------


-- ------------------------------
-- TABLE DATA: entities_to_warehouses
-- ------------------------------


-- ------------------------------
-- TABLE DATA: fields
-- ------------------------------

UPDATE fields:19cnkznj77b09zuqmugc CONTENT { help: 'Progressive homogeneous knowledge user', id: fields:19cnkznj77b09zuqmugc, title: 'Tagopia', type: 'datetime' };
UPDATE fields:22ky17f1h9kajpu1cisu CONTENT { help: 'Progressive explicit help-desk', id: fields:22ky17f1h9kajpu1cisu, title: 'Babblestorm', type: 'datetime' };
UPDATE fields:36a9tmke5hnuy5h2pczu CONTENT { help: 'Enterprise-wide incremental time-frame', id: fields:36a9tmke5hnuy5h2pczu, title: 'Bubbletube', type: 'integer' };
UPDATE fields:395v22p52ej8kyq69v6i CONTENT { help: 'Organic exuding access', id: fields:395v22p52ej8kyq69v6i, title: 'Photobug', type: 'integer' };
UPDATE fields:3muvyvyn9i8lt2npnyy3 CONTENT { help: 'Intuitive logistical data-warehouse', id: fields:3muvyvyn9i8lt2npnyy3, title: 'Oodoo', type: 'datetime' };
UPDATE fields:4816rl13ymz2k5lny7r1 CONTENT { help: 'Optimized transitional instruction set', id: fields:4816rl13ymz2k5lny7r1, title: 'Yakijo', type: 'date' };
UPDATE fields:4sxqxo1jb6mor72bryi6 CONTENT { help: 'Multi-channelled fresh-thinking success', id: fields:4sxqxo1jb6mor72bryi6, title: 'Youtags', type: 'date' };
UPDATE fields:53n17pyt5y6n69lg2qbe CONTENT { help: 'Persevering well-modulated portal', id: fields:53n17pyt5y6n69lg2qbe, title: 'Yacero', type: 'datetime' };
UPDATE fields:5c0i8xkj7eviobxinqrb CONTENT { help: 'Cross-platform incremental migration', id: fields:5c0i8xkj7eviobxinqrb, title: 'Edgepulse', type: 'date' };
UPDATE fields:5fzyvj3pkpgis73rm1fj CONTENT { help: 'Compatible scalable orchestration', id: fields:5fzyvj3pkpgis73rm1fj, title: 'Thoughtworks', type: 'string' };
UPDATE fields:5kz0lr9n0y1dbnfdiso6 CONTENT { help: 'Open-source encompassing portal', id: fields:5kz0lr9n0y1dbnfdiso6, title: 'Yotz', type: 'integer' };
UPDATE fields:5ubys5rvbokejaxxtuxw CONTENT { help: 'Optional upward-trending definition', id: fields:5ubys5rvbokejaxxtuxw, title: 'Dabshots', type: 'date' };
UPDATE fields:6sefbzhyorydohn3hdbx CONTENT { help: 'Intuitive holistic success', id: fields:6sefbzhyorydohn3hdbx, title: 'Fadeo', type: 'integer' };
UPDATE fields:6vsx977htp0kwxwymx6w CONTENT { help: 'Cross-platform grid-enabled pricing structure', id: fields:6vsx977htp0kwxwymx6w, title: 'Thoughtblab', type: 'integer' };
UPDATE fields:6x0myzxngxmdpy3uj0mf CONTENT { help: 'Cloned fault-tolerant support', id: fields:6x0myzxngxmdpy3uj0mf, title: 'Ainyx', type: 'date' };
UPDATE fields:75okpxlk2a3ulc6r1ec3 CONTENT { help: 'Polarised static emulation', id: fields:75okpxlk2a3ulc6r1ec3, title: 'Oozz', type: 'integer' };
UPDATE fields:7n390e8ondwr6idxsh5g CONTENT { help: 'Implemented local moratorium', id: fields:7n390e8ondwr6idxsh5g, title: 'Quinu', type: 'datetime' };
UPDATE fields:8c88ft524cwbigzerhap CONTENT { help: 'Switchable 24 hour open architecture', id: fields:8c88ft524cwbigzerhap, title: 'Mynte', type: 'integer' };
UPDATE fields:cdi8tw44128217q9umjs CONTENT { help: 'Total global pricing structure', id: fields:cdi8tw44128217q9umjs, title: 'Tambee', type: 'integer' };
UPDATE fields:dpy250o7d46xnf11qgae CONTENT { help: 'Exclusive executive artificial intelligence', id: fields:dpy250o7d46xnf11qgae, title: 'Dabtype', type: 'datetime' };
UPDATE fields:g9nr41gupa0hp8dyu8i9 CONTENT { help: 'Total stable solution', id: fields:g9nr41gupa0hp8dyu8i9, title: 'Youspan', type: 'date' };
UPDATE fields:gbp8wsvx3dr25xn8mv34 CONTENT { help: 'Inverse bottom-line website', id: fields:gbp8wsvx3dr25xn8mv34, title: 'Gigazoom', type: 'date' };
UPDATE fields:gsnulwrqa7s3uzfx8uwn CONTENT { help: 'Quality-focused intermediate emulation', id: fields:gsnulwrqa7s3uzfx8uwn, title: 'Wikizz', type: 'string' };
UPDATE fields:h57rniv916vztals3x6l CONTENT { help: 'Persevering motivating solution', id: fields:h57rniv916vztals3x6l, title: 'Photobean', type: 'datetime' };
UPDATE fields:i21efu2pkfnsin7ouc3o CONTENT { help: 'Integrated mission-critical process improvement', id: fields:i21efu2pkfnsin7ouc3o, title: 'Izio', type: 'datetime' };
UPDATE fields:il23j9vp7aotxpzxsa1f CONTENT { help: 'Right-sized attitude-oriented matrices', id: fields:il23j9vp7aotxpzxsa1f, title: 'Brainbox', type: 'date' };
UPDATE fields:iqfco7l2o5q0of0utzi2 CONTENT { help: 'Proactive non-volatile firmware', id: fields:iqfco7l2o5q0of0utzi2, title: 'Reallinks', type: 'string' };
UPDATE fields:kpduu5ywyu2beeks8rsa CONTENT { help: 'Streamlined 4th generation protocol', id: fields:kpduu5ywyu2beeks8rsa, title: 'Jabbercube', type: 'date' };
UPDATE fields:lbdy11fj5m2z84st9vgw CONTENT { help: 'Fundamental zero administration matrices', id: fields:lbdy11fj5m2z84st9vgw, title: 'Fanoodle', type: 'datetime' };
UPDATE fields:ld7inkmf9tg5vaf9dwrs CONTENT { help: 'Front-line composite service-desk', id: fields:ld7inkmf9tg5vaf9dwrs, title: 'Omba', type: 'datetime' };
UPDATE fields:m1fuu52splfm6hm1kmas CONTENT { help: 'Profit-focused solution-oriented policy', id: fields:m1fuu52splfm6hm1kmas, title: 'Brainverse', type: 'string' };
UPDATE fields:mqtfqu1g3dzhcslkmbpb CONTENT { help: 'Switchable grid-enabled workforce', id: fields:mqtfqu1g3dzhcslkmbpb, title: 'Minyx', type: 'integer' };
UPDATE fields:n3oi5zo8kv6gst89pt3h CONTENT { help: 'Cross-group 3rd generation policy', id: fields:n3oi5zo8kv6gst89pt3h, title: 'Browsebug', type: 'date' };
UPDATE fields:n5n71ul8rpnl434n3roh CONTENT { help: 'Public-key solution-oriented alliance', id: fields:n5n71ul8rpnl434n3roh, title: 'Kamba', type: 'string' };
UPDATE fields:ot4j3ehrzrtwzpo57jrs CONTENT { help: 'Configurable attitude-oriented interface', id: fields:ot4j3ehrzrtwzpo57jrs, title: 'Quire', type: 'string' };
UPDATE fields:pbo4d1hce6g9hcrg4jol CONTENT { help: 'Quality-focused full-range framework', id: fields:pbo4d1hce6g9hcrg4jol, title: 'Twinder', type: 'integer' };
UPDATE fields:pgtaabswctarr33v3766 CONTENT { help: 'Exclusive 24/7 hub', id: fields:pgtaabswctarr33v3766, title: 'Livefish', type: 'string' };
UPDATE fields:qvq9hxwsrdrds56yc9y1 CONTENT { help: 'Cross-platform upward-trending productivity', id: fields:qvq9hxwsrdrds56yc9y1, title: 'Shufflester', type: 'string' };
UPDATE fields:r19fixltzudm8egf5snl CONTENT { help: 'Synchronised context-sensitive extranet', id: fields:r19fixltzudm8egf5snl, title: 'Wordpedia', type: 'datetime' };
UPDATE fields:rhhz7s5c7scy0r5c28cm CONTENT { help: 'Devolved homogeneous emulation', id: fields:rhhz7s5c7scy0r5c28cm, title: 'Einti', type: 'string' };
UPDATE fields:rhw5c5fr1x2h5okouo67 CONTENT { help: 'Balanced contextually-based parallelism', id: fields:rhw5c5fr1x2h5okouo67, title: 'Jaloo', type: 'integer' };
UPDATE fields:sim3nqs6d8oqppk5259h CONTENT { help: 'Object-based eco-centric knowledge base', id: fields:sim3nqs6d8oqppk5259h, title: 'Bubblebox', type: 'datetime' };
UPDATE fields:uigo5pc9bau8y5wcpgb2 CONTENT { help: 'Up-sized intangible archive', id: fields:uigo5pc9bau8y5wcpgb2, title: 'Aimbu', type: 'date' };
UPDATE fields:v14ki1lwg13kil4w8jle CONTENT { help: 'Cross-group human-resource capability', id: fields:v14ki1lwg13kil4w8jle, title: 'Tazz', type: 'string' };
UPDATE fields:v8g8vs78zby4lyejz6rs CONTENT { help: 'Open-source bi-directional internet solution', id: fields:v8g8vs78zby4lyejz6rs, title: 'Gabspot', type: 'datetime' };
UPDATE fields:vj4bq28drrearyfuxw0p CONTENT { help: 'Polarised tangible info-mediaries', id: fields:vj4bq28drrearyfuxw0p, title: 'Realblab', type: 'string' };
UPDATE fields:y5ot9nrxud2rtuch5x6x CONTENT { help: 'Cross-platform content-based model', id: fields:y5ot9nrxud2rtuch5x6x, title: 'Digitube', type: 'date' };
UPDATE fields:ym8br170rbobumzvyjpy CONTENT { help: 'Switchable client-driven projection', id: fields:ym8br170rbobumzvyjpy, title: 'Zoonder', type: 'date' };
UPDATE fields:zeiqo8oe5don0sk7sjmi CONTENT { help: 'Networked intermediate conglomeration', id: fields:zeiqo8oe5don0sk7sjmi, title: 'Tanoodle', type: 'string' };
UPDATE fields:zva43qs3f0lcgkkssoaw CONTENT { help: 'User-friendly bifurcated solution', id: fields:zva43qs3f0lcgkkssoaw, title: 'Cogilith', type: 'date' };

-- ------------------------------
-- TABLE DATA: fields_to_warehouses
-- ------------------------------

RELATE fields:iqfco7l2o5q0of0utzi2 -> fields_to_warehouses:29l23i1bozeii1yzt8wj -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:29l23i1bozeii1yzt8wj, in: fields:iqfco7l2o5q0of0utzi2, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:il23j9vp7aotxpzxsa1f -> fields_to_warehouses:34l4n13foxt8tqf2ovpq -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:34l4n13foxt8tqf2ovpq, in: fields:il23j9vp7aotxpzxsa1f, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:zva43qs3f0lcgkkssoaw -> fields_to_warehouses:4pka7co9uxaeory8gbli -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:4pka7co9uxaeory8gbli, in: fields:zva43qs3f0lcgkkssoaw, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:gbp8wsvx3dr25xn8mv34 -> fields_to_warehouses:5j4mz1y3rpfunrpv29ob -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:5j4mz1y3rpfunrpv29ob, in: fields:gbp8wsvx3dr25xn8mv34, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:6vsx977htp0kwxwymx6w -> fields_to_warehouses:5q11t201tt333ac0ox9n -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:5q11t201tt333ac0ox9n, in: fields:6vsx977htp0kwxwymx6w, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:pbo4d1hce6g9hcrg4jol -> fields_to_warehouses:6uzhhydcv59ixjraifx9 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:6uzhhydcv59ixjraifx9, in: fields:pbo4d1hce6g9hcrg4jol, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:zeiqo8oe5don0sk7sjmi -> fields_to_warehouses:7krl6nt4gb8d7btoptc2 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:7krl6nt4gb8d7btoptc2, in: fields:zeiqo8oe5don0sk7sjmi, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:rhhz7s5c7scy0r5c28cm -> fields_to_warehouses:86po1aw9hzeb4jc7n886 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:86po1aw9hzeb4jc7n886, in: fields:rhhz7s5c7scy0r5c28cm, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:3muvyvyn9i8lt2npnyy3 -> fields_to_warehouses:8gv7sv55jzzcevxw63fd -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:8gv7sv55jzzcevxw63fd, in: fields:3muvyvyn9i8lt2npnyy3, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:g9nr41gupa0hp8dyu8i9 -> fields_to_warehouses:8lmoa6h0x5ebgzg217dd -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:8lmoa6h0x5ebgzg217dd, in: fields:g9nr41gupa0hp8dyu8i9, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:qvq9hxwsrdrds56yc9y1 -> fields_to_warehouses:a9lfdm79hq369qbihabj -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:a9lfdm79hq369qbihabj, in: fields:qvq9hxwsrdrds56yc9y1, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:6sefbzhyorydohn3hdbx -> fields_to_warehouses:bhf1zwhin2slonczh4od -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:bhf1zwhin2slonczh4od, in: fields:6sefbzhyorydohn3hdbx, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:lbdy11fj5m2z84st9vgw -> fields_to_warehouses:bk88cm5ne952txeom8ks -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:bk88cm5ne952txeom8ks, in: fields:lbdy11fj5m2z84st9vgw, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:i21efu2pkfnsin7ouc3o -> fields_to_warehouses:cae47lqxt8ymzl0j5nmd -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:cae47lqxt8ymzl0j5nmd, in: fields:i21efu2pkfnsin7ouc3o, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:m1fuu52splfm6hm1kmas -> fields_to_warehouses:coig7t1gw8zym6rprme3 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:coig7t1gw8zym6rprme3, in: fields:m1fuu52splfm6hm1kmas, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:ot4j3ehrzrtwzpo57jrs -> fields_to_warehouses:edpk0htdxcpf9nw2b7cf -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:edpk0htdxcpf9nw2b7cf, in: fields:ot4j3ehrzrtwzpo57jrs, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:7n390e8ondwr6idxsh5g -> fields_to_warehouses:eiit7f9w4d4awpiolos9 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:eiit7f9w4d4awpiolos9, in: fields:7n390e8ondwr6idxsh5g, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:19cnkznj77b09zuqmugc -> fields_to_warehouses:ej44vr3aky3efdixdxxk -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:ej44vr3aky3efdixdxxk, in: fields:19cnkznj77b09zuqmugc, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:5c0i8xkj7eviobxinqrb -> fields_to_warehouses:evgfb82tdbmttle0ijbn -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:evgfb82tdbmttle0ijbn, in: fields:5c0i8xkj7eviobxinqrb, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:395v22p52ej8kyq69v6i -> fields_to_warehouses:f1gs2fuaqjobkhey8nbw -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:f1gs2fuaqjobkhey8nbw, in: fields:395v22p52ej8kyq69v6i, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:kpduu5ywyu2beeks8rsa -> fields_to_warehouses:g6edu5v7zh4efihez2y9 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:g6edu5v7zh4efihez2y9, in: fields:kpduu5ywyu2beeks8rsa, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:5fzyvj3pkpgis73rm1fj -> fields_to_warehouses:g6m1bea9kztn9dm3d2po -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:g6m1bea9kztn9dm3d2po, in: fields:5fzyvj3pkpgis73rm1fj, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:rhw5c5fr1x2h5okouo67 -> fields_to_warehouses:hefhjkq1w5srqmgwctvw -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:hefhjkq1w5srqmgwctvw, in: fields:rhw5c5fr1x2h5okouo67, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:5ubys5rvbokejaxxtuxw -> fields_to_warehouses:hm8xre1rja22uw75nf3n -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:hm8xre1rja22uw75nf3n, in: fields:5ubys5rvbokejaxxtuxw, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:dpy250o7d46xnf11qgae -> fields_to_warehouses:hmvsi4jmj0brnzl21dke -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:hmvsi4jmj0brnzl21dke, in: fields:dpy250o7d46xnf11qgae, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:ld7inkmf9tg5vaf9dwrs -> fields_to_warehouses:i80wsj42ke4eg34feynd -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:i80wsj42ke4eg34feynd, in: fields:ld7inkmf9tg5vaf9dwrs, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:v8g8vs78zby4lyejz6rs -> fields_to_warehouses:ioykpckwxrpzbf9ghijh -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:ioykpckwxrpzbf9ghijh, in: fields:v8g8vs78zby4lyejz6rs, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:gsnulwrqa7s3uzfx8uwn -> fields_to_warehouses:jlsjn5hjberynsmmzjxq -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:jlsjn5hjberynsmmzjxq, in: fields:gsnulwrqa7s3uzfx8uwn, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:36a9tmke5hnuy5h2pczu -> fields_to_warehouses:juase0mf2xukgmppyqu4 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:juase0mf2xukgmppyqu4, in: fields:36a9tmke5hnuy5h2pczu, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:sim3nqs6d8oqppk5259h -> fields_to_warehouses:kjd3o3fnwue631wgdkcp -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:kjd3o3fnwue631wgdkcp, in: fields:sim3nqs6d8oqppk5259h, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:ym8br170rbobumzvyjpy -> fields_to_warehouses:l9mwej699fx7a98vtuwu -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:l9mwej699fx7a98vtuwu, in: fields:ym8br170rbobumzvyjpy, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:75okpxlk2a3ulc6r1ec3 -> fields_to_warehouses:mxmgnu2bfzgnvux9sv3k -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:mxmgnu2bfzgnvux9sv3k, in: fields:75okpxlk2a3ulc6r1ec3, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:mqtfqu1g3dzhcslkmbpb -> fields_to_warehouses:o1b42r4ybmm9jhsa5oe2 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:o1b42r4ybmm9jhsa5oe2, in: fields:mqtfqu1g3dzhcslkmbpb, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:y5ot9nrxud2rtuch5x6x -> fields_to_warehouses:onjpp0h5pvz5y3orq38h -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:onjpp0h5pvz5y3orq38h, in: fields:y5ot9nrxud2rtuch5x6x, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:vj4bq28drrearyfuxw0p -> fields_to_warehouses:pffdserxu2sqqd9qaz95 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:pffdserxu2sqqd9qaz95, in: fields:vj4bq28drrearyfuxw0p, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:22ky17f1h9kajpu1cisu -> fields_to_warehouses:pivoa0qajfzc5twlgynf -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:pivoa0qajfzc5twlgynf, in: fields:22ky17f1h9kajpu1cisu, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:r19fixltzudm8egf5snl -> fields_to_warehouses:qtw4ahekp63ohe3ykecr -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:qtw4ahekp63ohe3ykecr, in: fields:r19fixltzudm8egf5snl, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:4816rl13ymz2k5lny7r1 -> fields_to_warehouses:ri9t02numu3bjc2ubv2d -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:ri9t02numu3bjc2ubv2d, in: fields:4816rl13ymz2k5lny7r1, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:uigo5pc9bau8y5wcpgb2 -> fields_to_warehouses:soxll9qbhxx8b1dtrq7s -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:soxll9qbhxx8b1dtrq7s, in: fields:uigo5pc9bau8y5wcpgb2, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:n3oi5zo8kv6gst89pt3h -> fields_to_warehouses:sv2qlz86tby4wd9cxxwl -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:sv2qlz86tby4wd9cxxwl, in: fields:n3oi5zo8kv6gst89pt3h, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:8c88ft524cwbigzerhap -> fields_to_warehouses:tnw37mria54rfxm8jf6o -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:tnw37mria54rfxm8jf6o, in: fields:8c88ft524cwbigzerhap, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:4sxqxo1jb6mor72bryi6 -> fields_to_warehouses:u5hl7ygqssgu50qvhnnc -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:u5hl7ygqssgu50qvhnnc, in: fields:4sxqxo1jb6mor72bryi6, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:pgtaabswctarr33v3766 -> fields_to_warehouses:ua3y82wmdk195mubzz66 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:ua3y82wmdk195mubzz66, in: fields:pgtaabswctarr33v3766, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:h57rniv916vztals3x6l -> fields_to_warehouses:ubhitmalt1cz5a0o5i7a -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:ubhitmalt1cz5a0o5i7a, in: fields:h57rniv916vztals3x6l, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:v14ki1lwg13kil4w8jle -> fields_to_warehouses:utcjdpt65dw490yqybt6 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:utcjdpt65dw490yqybt6, in: fields:v14ki1lwg13kil4w8jle, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:5kz0lr9n0y1dbnfdiso6 -> fields_to_warehouses:vrz05n3ndbgozf73yemp -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:vrz05n3ndbgozf73yemp, in: fields:5kz0lr9n0y1dbnfdiso6, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:6x0myzxngxmdpy3uj0mf -> fields_to_warehouses:wr2bk67hk8x9ufkcge35 -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:wr2bk67hk8x9ufkcge35, in: fields:6x0myzxngxmdpy3uj0mf, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:53n17pyt5y6n69lg2qbe -> fields_to_warehouses:xseua2qlvjsdpwjkkmzx -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:xseua2qlvjsdpwjkkmzx, in: fields:53n17pyt5y6n69lg2qbe, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:cdi8tw44128217q9umjs -> fields_to_warehouses:y3zw182rpcaajsjzesan -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:y3zw182rpcaajsjzesan, in: fields:cdi8tw44128217q9umjs, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };
RELATE fields:n5n71ul8rpnl434n3roh -> fields_to_warehouses:y6imc4v365z13i6dbfge -> warehouse:018b1e6ca3397d8780fb21315ccbfbc1 CONTENT { __: true, id: fields_to_warehouses:y6imc4v365z13i6dbfge, in: fields:n5n71ul8rpnl434n3roh, out: warehouse:018b1e6ca3397d8780fb21315ccbfbc1 };

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

