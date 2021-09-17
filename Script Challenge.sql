

CREATE TABLE PERSONA (
    id        INTEGER NOT NULL AUTO_INCREMENT,
    dpi       INTEGER UNIQUE,
    nombre    VARCHAR(100),
    apellidos VARCHAR(100),
	PRIMARY KEY (id)
);

CREATE TABLE VACUNA (
    id     INTEGER NOT NULL AUTO_INCREMENT,
    vacuna VARCHAR(100),
    PRIMARY KEY (id)
);

CREATE TABLE DOSIS (
    dosis      INTEGER,
    fecha      DATE,
    vacuna_id  INTEGER NOT NULL,
    persona_id INTEGER NOT NULL,
    PRIMARY KEY (vacuna_id, persona_id),
    FOREIGN KEY (vacuna_id) REFERENCES VACUNA(id),
    FOREIGN KEY (persona_id) REFERENCES PERSONA(id)
);



