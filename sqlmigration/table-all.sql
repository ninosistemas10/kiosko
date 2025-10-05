
-- USUARIOS Y ROLES
CREATE TABLE usuarios (
    id              UUID            PRIMARY KEY,
    nombre          VARCHAR(100)    NOT NULL,
    email           VARCHAR(100)    NOT NULL,
    password        VARCHAR(255)    NOT NULL,
    rol             VARCHAR(200)    NOT NULL, 
    telefono        VARCHAR(20),
    direccion       TEXT,
    imagen_url      VARCHAR(255),
    activo          BOOLEAN         DEFAULT TRUE,
    created_at		INTEGER			NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
	updated_at		INTEGER
);


-- PROVEEDORES
CREATE TABLE proveedores (
    id                 UUID             PRIMARY KEY,
    nombre              VARCHAR(100)    NOT NULL,
    contacto            VARCHAR(100),
    telefono            VARCHAR(20),
    direccion           TEXT,
    email               VARCHAR(100),
    ruc                 VARCHAR(20),
    imagen_logo         VARCHAR(255),
    activo              BOOLEAN         DEFAULT TRUE,
    created_at		    INTEGER			NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
	updated_at		    INTEGER
);

--INVENTARIO

CREATE TABLE unidad_medida (
    id              UUID            PRIMARY KEY,
    nombre          VARCHAR(20)     NOT NULL,
    abreviatura     VARCHAR(5)      NOT NULL,
    activo          BOOLEAN         DEFAULT TRUE,
    created_at      INTEGER         NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
    updated_at      INTEGER
);

CREATE TABLE categorias(
    id              UUID            PRIMARY KEY,
    nombre		    VARCHAR(100)    NOT NULL UNIQUE,
    description     TEXT            NOT NULL,
	images			VARCHAR(250)	NOT NULL,
    activo          BOOLEAN         NOT NULL DEFAULT TRUE,
    created_at      INTEGER         NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
    updated_at      INTEGER
);


CREATE TABLE productos (
    id                  UUID           PRIMARY KEY,
    nombre              VARCHAR(100)    NOT NULL,
    descripcion         TEXT,
    precio_venta        DECIMAL(10,2)   NOT NULL,
    costo_promedio      DECIMAL(10,2),
    stock_minimo        INT,
    idUnidadMedida      UUID NOT NULL unidad_medida(id) ON DELETE CASCADE,
    idCategoria         UUID NOT NULL categorias(id) ON DELETE CASCADE,
    imagen              VARCHAR(255),
    activo              BOOLEAN     DEFAULT TRUE,
    destacado           BOOLEAN     DEFAULT FALSE,
    lleva_inventario    BOOLEAN DEFAULT TRUE,
    created_at          INTEGER         NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
    updated_at          INTEGER
);

CREATE TABLE subproductos (
    id                  UUID            PRIMARY KEY,
    idProducto          UUID            NOT NULL REFERENCES productos(id) ON DELETE CASCADE,
    nombre              VARCHAR(50)     NOT NULL,
    precio_adicional    DECIMAL(10,2)   DEFAULT 0,
    imagen              VARCHAR(255),
    activo              BOOLEAN DEFAULT TRUE,
    created_at          INTEGER         NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
    updated_at          INTEGER
);


-- COMPRAS

CREATE TABLE compras (
    id                   UUID           PRIMARY KEY,
    idProveedor          UUID           NOT NULL REFERENCES proveedores(id) ON DELETE CASCADE,
    fecha_compra         DATE           NOT NULL,
    fecha_recibido       DATE,
    numero_factura       VARCHAR(50),
    subtotal             DECIMAL(12,2)  NOT NULL,
    iva                  DECIMAL(12,2)  DEFAULT 0,
    total                DECIMAL(12,2)  NOT NULL,
    estado               VARCHAR(20)    CHECK (estado IN ('pendiente', 'recibido', 'parcial', 'cancelado')) DEFAULT 'pendiente',
    metodo_pago          VARCHAR(20),
    imagen_factura_url   VARCHAR(255),
    notas                TEXT,
    created_at           INTEGER         NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
    updated_at           INTEGER   
);


CREATE TABLE detalle_compras (
    id                  UUID            PRIMARY KEY,
    idCompra            UUID            NOT NULL REFERENCES compras(id_compra) ON DELETE CASCADE,
    idProducto          UUID            NOT NULL REFERENCES productos(id_producto),
    cantidad            DECIMAL(10,3)   NOT NULL,
    precio_unitario     DECIMAL(10,2)   NOT NULL,
    idUnidad            UUID            NOT NULL REFERENCES unidades_medida(id_unidad),
    lote                VARCHAR(50),
    fecha_vencimiento   DATE,
    created_at          INTEGER         NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
    updated_at          INTEGER   
);

--PEDIDOS

CREATE TABLE pedidos (
    id                  UUID            PRIMARY KEY,
    id_usuario          UUID REFERENCES usuarios(id),
    cliente_nombre      VARCHAR(100),
    fecha_pedido        TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    fecha_entrega       TIMESTAMP,
    tipo_entrega        VARCHAR(20) CHECK (tipo_entrega IN ('local', 'domicilio', 'recoger')),
    subtotal            DECIMAL(10,2) NOT NULL,
    descuento           DECIMAL(10,2) DEFAULT 0,
    iva                 DECIMAL(10,2) DEFAULT 0,
    total               DECIMAL(10,2) NOT NULL,
    estado              VARCHAR(20) CHECK (estado IN ('pendiente', 'confirmado', 'en_preparacion', 'listo', 'entregado', 'cancelado')) DEFAULT 'pendiente',
    id_mesa             INT,
    direccion_entrega   TEXT,
    notas               TEXT,
    activo              BOOLEAN DEFAULT TRUE
    created_at          INTEGER         NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
    updated_at          INTEGER   
);


CREATE TABLE detalle_pedidos (
    id_detalle          UUID            PRIMARY KEY,
    id_pedido           UUID            NOT NULL REFERENCES pedidos(id_pedido) ON DELETE CASCADE,
    id_producto         UUID            REFERENCES productos(id_producto),
    id_subproducto      UUID            REFERENCES subproductos(id_subproducto),
    cantidad            NUMERIC         NOT NULL,
    precio_unitario     DECIMAL(10,2)   NOT NULL,
    instrucciones       TEXT,
    imagen              VARCHAR(255),
    estado              VARCHAR(20) DEFAULT 'pendiente',

    CONSTRAINT chk_productos_pedido CHECK (
        (id_producto IS NOT NULL AND id_subproducto IS NULL AND id_promocion IS NULL)
        OR (id_producto IS NULL AND id_subproducto IS NOT NULL AND id_promocion IS NULL)
    )
);

--PAGOS

CREATE TABLE pagos (
    id_pago SERIAL PRIMARY KEY,
    id_pedido INT REFERENCES pedidos(id_pedido),
    id_usuario INT REFERENCES usuarios(id_usuario),
    fecha_pago TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    monto DECIMAL(10,2) NOT NULL,
    metodo_pago VARCHAR(20) CHECK (metodo_pago IN ('efectivo', 'tarjeta_debito', 'tarjeta_credito', 'transferencia', 'app_pago')),
    referencia VARCHAR(50),
    comprobante_url VARCHAR(255),
    estado VARCHAR(20) DEFAULT 'completado'
);

--MOVIMIENTOS DE INVENTARIO

CREATE TABLE movimientos_inventario (
    id_movimiento SERIAL PRIMARY KEY,
    id_producto INT NOT NULL REFERENCES productos(id_producto),
    tipo VARCHAR(20) CHECK (tipo IN ('compra', 'venta', 'ajuste', 'consumo', 'devolucion')),
    cantidad DECIMAL(10,3) NOT NULL,
    id_unidad INT REFERENCES unidades_medida(id_unidad),
    fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    id_usuario INT REFERENCES usuarios(id_usuario),
    id_referencia INT,
    tabla_referencia VARCHAR(30),
    costo_unitario DECIMAL(10,2),
    notas TEXT
);

--CAJA Y TURNOS

CREATE TABLE caja (
    id_caja SERIAL PRIMARY KEY,
    fecha_apertura TIMESTAMP NOT NULL,
    fecha_cierre TIMESTAMP,
    id_usuario_apertura INT REFERENCES usuarios(id_usuario),
    id_usuario_cierre INT REFERENCES usuarios(id_usuario),
    monto_inicial DECIMAL(10,2) NOT NULL,
    monto_final DECIMAL(10,2),
    estado VARCHAR(20) CHECK (estado IN ('abierta', 'cerrada', 'conciliada')) DEFAULT 'abierta',
    notas TEXT
);


CREATE TABLE turnos (
    id_turno SERIAL PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    hora_inicio TIME NOT NULL,
    hora_fin TIME NOT NULL,
);



CREATE TABLE resenas (
    id_resena SERIAL PRIMARY KEY,
    id_usuario INT NOT NULL REFERENCES usuarios(id_usuario),
    id_producto INT REFERENCES productos(id_producto),
    id_pedido INT REFERENCES pedidos(id_pedido),
    calificacion INT CHECK (calificacion BETWEEN 1 AND 5),
    comentario TEXT,
    imagenes TEXT[],
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    respuesta_admin TEXT
);

CREATE TABLE configuraciones (
    id_config SERIAL PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    valor TEXT,
    descripcion TEXT
);
