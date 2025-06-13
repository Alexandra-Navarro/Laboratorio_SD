-- Tabla: escuela
CREATE TABLE escuela (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL UNIQUE,
    direccion VARCHAR(255) NOT NULL UNIQUE,
    comuna VARCHAR(100) NOT NULL
);

-- Tabla: sala
CREATE TABLE sala (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    escuela_id INTEGER NOT NULL,
    CONSTRAINT fk_sala_escuela FOREIGN KEY (escuela_id) REFERENCES escuela(id)
);

-- Tabla: variable_ambiental
CREATE TABLE variable_ambiental (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    unidad_medida VARCHAR(50) NOT NULL
);

-- Tabla: sensor
CREATE TABLE sensor (
    id SERIAL PRIMARY KEY,
    modelo VARCHAR(100) NOT NULL,
    estado VARCHAR(50) NOT NULL,
    fecha_instalacion DATE NOT NULL,
    sala_id INTEGER NOT NULL,
    variable_id INTEGER NOT NULL,
    CONSTRAINT fk_sensor_sala FOREIGN KEY (sala_id) REFERENCES sala(id),
    CONSTRAINT fk_sensor_variable FOREIGN KEY (variable_id) REFERENCES variable_ambiental(id)
);

-- Tabla: medicion
CREATE TABLE medicion (
    id SERIAL PRIMARY KEY,
    fecha TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    valor FLOAT NOT NULL,
    sensor_id INTEGER NOT NULL,
    CONSTRAINT fk_medicion_sensor FOREIGN KEY (sensor_id) REFERENCES sensor(id)
);

-- Tabla: alerta
CREATE TABLE alerta (
    id SERIAL PRIMARY KEY,
    tipo VARCHAR(50) CHECK (tipo IN ('informativo', 'preventivo', 'critico')),
    descripcion TEXT,
    valor_detectado FLOAT NOT NULL,
    umbral VARCHAR(50) CHECK (umbral IN ('bajo', 'alto', 'variable')),
    fecha TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    sala_id INTEGER NOT NULL,
    CONSTRAINT fk_alerta_sala FOREIGN KEY (sala_id) REFERENCES sala(id)
);

-- Tabla: usuario
CREATE TABLE usuario (
    rut_usuario VARCHAR(10) PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    rol VARCHAR(50) CHECK (rol IN ('admin', 'user')),
    escuela_id INTEGER NOT NULL,
    CONSTRAINT fk_usuario_escuela FOREIGN KEY (escuela_id) REFERENCES escuela(id)
);

-- Tabla: umbrales
CREATE TABLE umbrales (
    id SERIAL PRIMARY KEY,
    umbral_bajo FLOAT,
    umbral_alto FLOAT,
    umbral_bajo_prev FLOAT,
    umbral_alto_prev FLOAT,
    sala_id INTEGER NOT NULL,
    variable_id INTEGER NOT NULL,
    CONSTRAINT fk_umbrales_sala FOREIGN KEY (sala_id) REFERENCES sala(id),
    CONSTRAINT fk_umbrales_variable FOREIGN KEY (variable_id) REFERENCES variable_ambiental(id)
);
