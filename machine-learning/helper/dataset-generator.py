from faker import Faker
import csv
import random

# Configuração do Faker
fake = Faker()

# Caminho do arquivo CSV a ser gerado
output_file = 'products.csv'

# Número de itens a serem gerados
num_items = 100000

# Cabeçalhos do CSV
headers = [
    "Product ID", "SKU", "Name", "Category", "Price", "Original Price", "Discount", "Description",
    "Weight (kg)", "Dimensions (cm)", "Brand", "Color", "Warranty (months)", "Stock (units)", "Trending"
]

# Lista de categorias específicas
categories = [
    "Electronics",
    "Books",
    "Fashion",
    "Home & Kitchen",
    "Sports"
]

# Rastreando IDs únicos
product_ids = set()
skus = set()

# Geração do CSV
with open(output_file, mode='w', newline='', encoding='utf-8') as file:
    writer = csv.writer(file)
    writer.writerow(headers)

    for i in range(num_items):
        # Geração de IDs únicos
        while True:
            product_id = f"P{random.randint(10000, 999999)}"
            if product_id not in product_ids:
                product_ids.add(product_id)
                break

        while True:
            sku = f"SKU-{random.randint(1000, 99999)}"
            if sku not in skus:
                skus.add(sku)
                break

        # Geração dos demais atributos
        name = fake.word().capitalize() + " " + random.choice(categories).split()[0]
        category = random.choice(categories)
        price_original = round(random.uniform(50, 2000), 2)  # Preço original aleatório
        discount = random.randint(5, 50)  # Desconto entre 5% e 50%
        price = round(price_original * (1 - discount / 100), 2)
        description = fake.sentence(nb_words=15)
        weight = round(random.uniform(0.5, 10.0), 1)
        dimensions = f"{random.randint(10, 100)}x{random.randint(10, 100)}x{random.randint(10, 100)}"
        brand = fake.company()
        color = random.choice(["Black", "White", "Red", "Blue", "Green", "Yellow", "Gray"])
        warranty = random.randint(6, 24)
        stock = random.randint(0, 500)
        trending = random.choice([0, 1])

        # Escrevendo os dados no CSV
        writer.writerow([
            product_id, sku, name, category, f"R$ {price}", f"R$ {price_original}", f"{discount}%",
            description, weight, dimensions, brand, color, warranty, stock, trending
        ])

        # Flush a cada 100 itens
        if (i + 1) % 100 == 0:
            file.flush()

print(f"Arquivo CSV com {num_items} itens gerado: {output_file}")
