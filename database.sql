
DROP TABLE IF EXISTS userinfo;

CREATE TABLE IF NOT EXISTS userinfo (
    userid biginit insigned NOT NULL,
    username TEXT DEFAULT "",
    country INTEGER DEFAULT 0,
    banned INTEGER DEFAULT 0,
    bannedreason TEXT DEFAULT "",
    privilages INTEGER DEFAULT 0,
    joinedat DATE DEFAULT (CURRENT_DATE),
);

DROP TABLE IF EXISTS authenticationpairs;

CREATE TABLE IF NOT EXISTS authenticationpairs (
    userid biginit unsigned NOT NULL,
    password TEXT DEFAULT "",
);

DROP TABLE IF EXISTS friends;

CREATE TABLE IF NOT EXISTS friends (
    userid biginit unsigned NOT NULL,
    friendid biginit unsigned NOT NULL,
);

DROP TABLE IF EXISTS userpage;

CREATE TABLE IF NOT EXISTS userpage (
    userid bigint unsigned NOT NULL,
    title TEXT DEFAULT "",
    website TEXT DEFAULT "",
    aboutme TEXT DEFAULT "",
    twitter TEXT DEFAULT "",
);

DROP TABLE IF EXISTS inventory;

CREATE TABLE IF NOT EXISTS inventory (
    userid bigint unsigned NOT NULL,
    itemid bigint unsigned NOT NULL,
    itemcount bigint unsigned DEFAULT 1,
);

DROP TABLE IF EXISTS worlds;

CREATE TABLE IF NOT EXISTS worlds (
    worldid bigint unsigned NOT NULL,
    mapid biginit unsigned NOT NULL,
);

DROP TABLE IF NOT EXISTS usersession;

CREATE TABLE IF NOT EXISTS usersession (
    userid bigint unsigned NOT NULL,
    worldid bigint unsigned NOT NULL,
    prevworldid bigint unsigned NOT NULL,

    position_x bigint DEFAULT 0,
    position_y bigint DEFAULT 0,

    headitem bigint unsigned DEFAULT 0,
    torsoitem bigint unsigned DEFAULT 0,
    legsitem bigint unsigned DEFAULT 0,
    shoesitem bigint unsigned DEFAULT 0,

    headitemsecondary bigint unsigned DEFAULT 0,
    torsoitemsecondary bigint unsigned DEFAULT 0,
    legsitemsecondary bigint unsigned DEFAULT 0,
    shoesitemsecondary bigint unsigned DEFAULT 0,

    primaryweapon bigint unsigned DEFAULT 0,
    secondaryweapon bigint unsigned DEFAULT 0,
    pet bigint unsigned DEFAULT 0,
);

