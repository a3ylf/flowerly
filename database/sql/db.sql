
CREATE TABLE plants (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    scientific_name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    category VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    stock_quantity INTEGER NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    origin_location VARCHAR(255) NOT NULL
);

INSERT INTO plants (name, scientific_name, description, category, price, stock_quantity, image_url, origin_location)
VALUES 
('Rosa Vermelha', 'Rosa rubiginosa', 'Rosas vermelhas de alta qualidade, perfeitas para buques.', 'flores', 5.99, 150, 'rosa_vermelha.png', 'Areia'),
('Orquidea Branca', 'Phalaenopsis amabilis', 'Orquidea Phalaenopsis branca, ideal para presentes.', 'plantas', 25.00, 60, 'orquidea_branca.png', 'Bananeiras'),
('Bromelia', 'Bromelia laciniosa', 'Planta tropical resistente, comum em regioes aridas.', 'plantas', 18.50, 85, 'bromelia.png', 'Campina Grande'),
('Hibisco', 'Hibiscus rosa-sinensis', 'Flores vibrantes usadas em paisagismo e chas.', 'flores', 7.99, 200, 'hibisco.png', 'Joao Pessoa'),
('Cacto Mandacaru', 'Cereus jamacaru', 'Cacto tipico do semiarido nordestino.', 'suculentas', 12.00, 100, 'mandacaru.png', 'Patos'),
('Palmeira Imperial', 'Roystonea oleracea', 'Palmeira de grande porte, muito usada em jardins.', 'arvores', 120.00, 15, 'palmeira_imperial.png', 'Areia'),
('Jiboia', 'Epipremnum aureum', 'Planta trepadeira muito popular em interiores.', 'plantas', 10.99, 75, 'jiboia.png', 'Cajazeiras'),
('Samambaia', 'Nephrolepis exaltata', 'Planta ornamental muito utilizada em varandas.', 'plantas', 14.50, 95, 'samambaia.png', 'Mari'),
('Ipe Amarelo', 'Handroanthus albus', 'Arvore ornamental com flores amarelas vibrantes.', 'arvores', 85.00, 40, 'ipe_amarelo.png', 'Areia'),
('Cafe', 'Coffea arabica', 'Planta cultivada para a producao de cafe.', 'arvores', 30.00, 120, 'cafe.png', 'Brejo da Paraiba'),
('Flor de Maio', 'Schlumbergera truncata', 'Planta suculenta com flores coloridas.', 'suculentas', 15.00, 60, 'flor_de_maio.png', 'Mari'),
('Maranta', 'Maranta leuconeura', 'Planta de folhas ornamentais, tambem conhecida como planta rezadeira.', 'plantas', 20.00, 50, 'maranta.png', 'Guarabira'),
('Costela de Adao', 'Monstera deliciosa', 'Planta de folhas grandes e recortadas, muito usada em decoracao.', 'plantas', 35.00, 40, 'costela_de_adao.png', 'Campina Grande'),
('Violeta', 'Saintpaulia ionantha', 'Planta de pequenas flores, popular em interiores.', 'flores', 8.50, 180, 'violeta.png', 'Alagoa Grande'),
('Aroeira', 'Schinus terebinthifolius', 'Arvore com frutos usados como tempero.', 'arvores', 65.00, 30, 'aroeira.png', 'Mari'),
('Anturio', 'Anthurium andraeanum', 'Planta com inflorescencia vermelha, muito usada em decoracao.', 'flores', 22.00, 70, 'anturio.png', 'Alagoa Nova'),
('Espada de Sao Jorge', 'Dracaena trifasciata', 'Planta resistente e facil de cuidar, muito usada em protecao espiritual.', 'plantas', 18.00, 80, 'espada_de_sao_jorge.png', 'Santa Rita'),
('Manaca da Serra', 'Tibouchina mutabilis', 'Arvore ornamental com flores roxas e brancas.', 'arvores', 90.00, 20, 'manaca_da_serra.png', 'Esperanca'),
('Suculenta Echeveria', 'Echeveria elegans', 'Planta ornamental de facil cuidado.', 'suculentas', 12.00, 150, 'echeveria.png', 'Bananeiras'),
('Jacaranda', 'Jacaranda mimosifolia', 'Arvore ornamental com flores roxas.', 'arvores', 110.00, 25, 'jacaranda.png', 'Solanea');


-- Concede permissões ao novo usuário

CREATE TABLE client (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    cpf VARCHAR(11) UNIQUE not null,
    rua TEXT not null,
    num smallint not null
);

INSERT INTO client (name, email, password, cpf,rua,num)
VALUES ('client', 'client@example.com', '$2a$10$w9HUkNydSBqJnUOngDrLN..O5yZVHM/D9wXlEHhlV7fpM6SQVZXNS', '12345678901','bruno almeida', 122);

CREATE TABLE vendor (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,   
    cpf varchar(11) unique not null
);
INSERT INTO vendor (name, email, password, cpf)
VALUES ('Admin', 'admin@example.com', '$2a$10$w9HUkNydSBqJnUOngDrLN..O5yZVHM/D9wXlEHhlV7fpM6SQVZXNS', '12345678901');

create table cart (
    id serial primary key,
    client_id int references client(id) on delete cascade
);

create table cart_item (
    cart_id int references cart(id) on delete cascade,     
    product_id int references plants(id) on delete cascade,
    quantity int not null check (quantity > 0),            
    primary key (cart_id, product_id)
);
INSERT INTO cart (client_id)
VALUES (1);
INSERT INTO cart (client_id)
VALUES (1);
INSERT INTO cart_item (cart_id,product_id,quantity)
VALUES (1,3,4);
INSERT INTO cart_item (cart_id,product_id,quantity)
VALUES (1,2,5);

CREATE TABLE purchase (
    id SERIAL PRIMARY KEY,
    cart_id INT NOT NULL REFERENCES cart(id) ON DELETE CASCADE,
    purchase_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    total_amount DECIMAL(10, 2) NOT NULL,
    payment_status VARCHAR(50) NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    vendor_id INT NULL,
  FOREIGN KEY(vendor_id) REFERENCES vendor(id)
);


ALTER TABLE purchase ADD CONSTRAINT fk_cart
FOREIGN KEY (cart_id) REFERENCES cart(id);

ALTER TABLE cart_item ADD CONSTRAINT fk_product
FOREIGN KEY (product_id) REFERENCES plants(id);




CREATE OR REPLACE FUNCTION update_stock(
    p_id INTEGER,
    p_stock_quantity INTEGER
) RETURNS VOID AS $$
BEGIN
    UPDATE plants
    SET stock_quantity = p_stock_quantity
    WHERE id = p_id;
END;
$$ LANGUAGE plpgsql;

CREATE VIEW client_purchases_view AS
SELECT 
    client.id AS client_id,
    client.name AS client_name,
    purchase.id AS purchase_id,
    purchase.purchase_date,
    purchase.total_amount,
    purchase.payment_status,
    purchase.payment_method,
    cart_item.product_id,
    plants.name AS product_name,
    cart_item.quantity
FROM 
    client
JOIN 
    cart ON client.id = cart.client_id
JOIN 
    purchase ON cart.id = purchase.cart_id
JOIN 
    cart_item ON cart.id = cart_item.cart_id
JOIN 
    plants ON cart_item.product_id = plants.id;

CREATE ROLE boss WITH LOGIN SUPERUSER PASSWORD '123456';
GRANT ALL PRIVILEGES ON DATABASE flowerly TO boss;

