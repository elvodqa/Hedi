DROP TABLE IF EXISTS userinfo;
--DEFAULT(CURRENT_DATE)
CREATE TABLE IF NOT EXISTS userinfo (
    userid BIGINT UNSIGNED NOT NULL,
    username TEXT DEFAULT "",
    password TEXT DEFAULT "",
    country INTEGER DEFAULT 0,
    banned INTEGER DEFAULT 0,
    bannedreason TEXT DEFAULT "",
    privilages INTEGER DEFAULT 0,
    joinedat DATE DEFAULT NULL
);   

DROP TABLE IF EXISTS authenticationpairs;

-- CREATE TABLE IF NOT EXISTS authenticationpairs (
--     userid BIGINT UNSIGNED NOT NULL,
--     password TEXT DEFAULT ""
-- );

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
    aboutme TEXT DEFAULT ""
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
    mapid BIGINT UNSIGNED NOT NULL,
    worldname TEXT DEFAULT "",
    owner TEXT DEFAULT "",
    playerlimit BIGINT DEFAULT 255,
    playercount BIGINT DEFAULT 0
);

DROP TABLE IF EXISTS usersession;

CREATE TABLE IF NOT EXISTS usersession (
    userid BIGINT UNSIGNED NOT NULL,
    online INT DEFAULT 0,
    worldid BIGINT UNSIGNED DEFAULT 0,
    prevworldid BIGINT UNSIGNED DEFAULT 0,

    positionx bigint DEFAULT 0,
    positiony bigint DEFAULT 0,

    itemhead BIGINT UNSIGNED DEFAULT 0,
    itemtorso BIGINT UNSIGNED DEFAULT 0,
    itemlegs BIGINT UNSIGNED DEFAULT 0,
    itemshoes BIGINT UNSIGNED DEFAULT 0,

    itemheadsecondary BIGINT UNSIGNED DEFAULT 0,
    itemtorsosecondary BIGINT UNSIGNED DEFAULT 0,
    itemlegssecondary BIGINT UNSIGNED DEFAULT 0,
    itemshoessecondary BIGINT UNSIGNED DEFAULT 0,

    weaponprimary BIGINT UNSIGNED DEFAULT 0,
    weaponsecondary BIGINT UNSIGNED DEFAULT 0,
    pet BIGINT UNSIGNED DEFAULT 0
);


