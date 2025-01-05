import joblib
import pandas as pd
import numpy as np
from sklearn.preprocessing import MinMaxScaler
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity
from sklearn.neighbors import NearestNeighbors

# Carregar o dataset
# Substitua 'product.csv' pelo caminho correto do seu arquivo
data = pd.read_csv("./dataset/products.csv")
data["Price"] = data["Price"].str.replace("R\\$", "", regex=True).astype(float)
data["Original Price"] = data["Original Price"].str.replace("R\\$", "", regex=True).astype(float)

# Exibir as primeiras linhas do dataset
print("Visualizando os dados:")
print(data.head())

# Verificar valores ausentes e preenchê-los
print("\nChecando valores ausentes:")
print(data.isnull().sum())
data.fillna("", inplace=True)

# Normalizar colunas numéricas
print("\nNormalizando colunas numéricas...")
scaler = MinMaxScaler()
data[["Price", "Original Price", "Weight (kg)", "Stock (units)"]] = scaler.fit_transform(
    data[["Price", "Original Price", "Weight (kg)", "Stock (units)"]]
)

# Transformar descrições em vetores numéricos usando TF-IDF
print("\nProcessando a coluna Description com TF-IDF...")
tfidf = TfidfVectorizer(max_features=100)
description_vectors = tfidf.fit_transform(data["Description"]).toarray()

# Aplicar One-Hot Encoding em colunas categóricas (Category e Color)
print("\nAplicando One-Hot Encoding em colunas categóricas...")
data = pd.get_dummies(data, columns=["Category", "Color"])

# Combinar todas as features
print("\nCombinando features...")
features = np.hstack([
    data[["Price", "Original Price", "Weight (kg)", "Stock (units)"]].values,
    description_vectors
])

# Treinar modelo baseado em similaridade (cosine similarity)
print("\nCalculando matriz de similaridade...")
similarity_matrix = cosine_similarity(features)

# Treinar modelo k-NN para recomendação
print("\nTreinando modelo k-NN...")
knn = NearestNeighbors(metric='cosine', algorithm='brute')
knn.fit(features)

# Salvar dados e modelo
print("\nSalvando dados do modelo...")
data.to_csv("../projects/recommendation-api/model/processed_data.csv", index=False)
np.save("../projects/recommendation-api/model/features.npy", features)
joblib.dump(knn, "../projects/recommendation-api/model/knn_model.pkl")
print("Modelo e dados salvos com sucesso!")

# Função para recomendar itens
print("\nDefinindo função de recomendação...")
def recommend(product_id, n=5):
    try:
        idx = data[data["Product ID"] == product_id].index[0]
        distances, indices = knn.kneighbors([features[idx]], n_neighbors=n+1)
        recommended_ids = data.iloc[indices[0][1:]]["Product ID"].tolist()
        return recommended_ids
    except IndexError:
        return "Product ID não encontrado no dataset."

# Exemplo de recomendação
example_product_id = data["Product ID"].iloc[0]
print(f"\nRecomendações para o produto {example_product_id}:")
print(recommend(example_product_id))

# Avaliação simples usando Trending como indicador de relevância
print("\nAvaliando precisão do modelo...")
true_relevant = data["Trending"].values
recommended_indices = knn.kneighbors([features[0]], n_neighbors=6)[1][0][1:]
precision = sum(true_relevant[i] for i in recommended_indices) / len(recommended_indices)
print(f"Precision@5: {precision:.2f}")
