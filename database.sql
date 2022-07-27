DROP TABLE IF EXISTS userinfo;
--DEFAULT(CURRENT_DATE)
CREATE TABLE IF NOT EXISTS userinfo (
    userid BIGINT UNSIGNED NOT NULL,
    username TEXT DEFAULT "",
    country INTEGER DEFAULT 0,
    banned INTEGER DEFAULT 0,
    bannedreason TEXT DEFAULT "",
    privilages INTEGER DEFAULT 0,
    joinedat DATE DEFAULT NULL
);   

DROP TABLE IF EXISTS authenticationpairs;

CREATE TABLE IF NOT EXISTS authenticationpairs (
    userid BIGINT UNSIGNED NOT NULL,
    password TEXT DEFAULT ""
);

DROP TABLE IF EXISTS friends;

CREATE TABLE IF NOT EXISTS friends (
    userid BIGINT UNSIGNED NOT NULL,
    friendid BIGINT UNSIGNED NOT NULL
);

DROP TABLE IF EXISTS userpage;

CREATE TABLE IF NOT EXISTS userpage (
    userid BIGINT UNSIGNED NOT NULL,
    title TEXT DEFAULT "",
    website TEXT DEFAULT "",
    aboutme TEXT DEFAULT "",
    twitter TEXT DEFAULT ""
);

DROP TABLE IF EXISTS inventory;

CREATE TABLE IF NOT EXISTS inventory (
    userid BIGINT UNSIGNED NOT NULL,
    itemid BIGINT UNSIGNED NOT NULL,
    itemcount BIGINT UNSIGNED DEFAULT 1
);

DROP TABLE IF EXISTS worlds;

CREATE TABLE IF NOT EXISTS worlds (
    worldid BIGINT UNSIGNED NOT NULL,
    mapid BIGINT UNSIGNED NOT NULL
);

DROP TABLE IF EXISTS usersession;

CREATE TABLE IF NOT EXISTS usersession (
    userid BIGINT UNSIGNED NOT NULL,
    worldid BIGINT UNSIGNED NOT NULL,
    prevworldid BIGINT UNSIGNED NOT NULL,

    position_x bigint DEFAULT 0,
    position_y bigint DEFAULT 0,

    headitem BIGINT UNSIGNED DEFAULT 0,
    torsoitem BIGINT UNSIGNED DEFAULT 0,
    legsitem BIGINT UNSIGNED DEFAULT 0,
    shoesitem BIGINT UNSIGNED DEFAULT 0,

    headitemsecondary BIGINT UNSIGNED DEFAULT 0,
    torsoitemsecondary BIGINT UNSIGNED DEFAULT 0,
    legsitemsecondary BIGINT UNSIGNED DEFAULT 0,
    shoesitemsecondary BIGINT UNSIGNED DEFAULT 0,

    primaryweapon BIGINT UNSIGNED DEFAULT 0,
    secondaryweapon BIGINT UNSIGNED DEFAULT 0,
    pet BIGINT UNSIGNED DEFAULT 0
);

