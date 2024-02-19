CREATE TABLE IF NOT EXISTS food_listing (
    id SERIAL PRIMARY KEY,
    establishment_name VARCHAR(255) NOT NULL,
    food_type VARCHAR(100) NOT NULL,
    location VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    quantity_available INT NOT NULL,
    description TEXT
);

INSERT INTO food_listing (establishment_name, food_type, location, price, quantity_available, description) 
VALUES 
    ('Café da Maria', 'Café da Manhã', 'Rua das Flores, 123', 15.90, 5, 'Opções de café da manhã variadas'),
    ('Pizzaria Bella Napoli', 'Pizza', 'Av. Principal, 456', 29.99, 3, 'Pizza de calabresa, mussarela e muito mais'),
    ('Restaurante Sabor Nordestino', 'Comida Nordestina', 'Rua do Sertão, 789', 24.50, 8, 'Pratos típicos da culinária nordestina'),
    ('Sushi Express', 'Sushi', 'Rua dos Peixes, 567', 39.90, 2, 'Sushis frescos e variados'),
    ('Pastelaria da Vovó', 'Pastel', 'Av. Central, 890', 5.99, 10, 'Pastéis recheados com diversos sabores'),
    ('Hamburgueria Top Burguer', 'Hambúrguer', 'Rua dos Hamburguers, 321', 19.99, 6, 'Hambúrgueres artesanais com opções vegetarianas'),
    ('Padaria Doce Sabor', 'Padaria', 'Av. das Delícias, 111', 8.50, 4, 'Pães, bolos, salgados e mais'),
    ('Casa do Açaí', 'Açaí', 'Rua das Palmeiras, 222', 12.99, 7, 'Açaí na tigela com acompanhamentos diversos'),
    ('Sorveteria Gelatto', 'Sorvete', 'Av. dos Sabores, 333', 6.50, 9, 'Sorvetes de diversos sabores'),
    ('Sabor do Churrasco', 'Churrasco', 'Rua do Churrasco, 444', 35.99, 1, 'Carnes selecionadas e acompanhamentos');