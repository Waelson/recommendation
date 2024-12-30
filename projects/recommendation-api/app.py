from flask import Flask, request, jsonify
import pandas as pd
import numpy as np
import joblib

# Inicializar o Flask
app = Flask(__name__)

# Carregar os dados e o modelo pré-processados
data = pd.read_csv("model/processed_data.csv")
features = np.load("model/features.npy")
knn = joblib.load("model/knn_model.pkl")

# Função para gerar recomendações
def recommend(product_id, n=5):
    try:
        idx = data[data["Product ID"] == product_id].index[0]
        distances, indices = knn.kneighbors([features[idx]], n_neighbors=n + 1)
        recommended_ids = data.iloc[indices[0][1:]]["Product ID"].tolist()
        return recommended_ids
    except IndexError:
        return []

# Rota para a API
@app.route('/recommend', methods=['GET'])
def get_recommendations():
    print("\nObtendo recomendacao...")
    product_id = request.args.get('product_id')
    n = int(request.args.get('n', 5))  # Número de recomendações, padrão 5

    if not product_id:
        return jsonify({"error": "Product ID é obrigatório"}), 400

    recommendations = recommend(product_id, n)
    return jsonify({"product_id": product_id, "recommendations": recommendations})

# Rodar o servidor
if __name__ == '__main__':
    print("\nApp recommendation-api running at: http://localhost:5000")
    app.run(host='0.0.0.0', port=8000)
