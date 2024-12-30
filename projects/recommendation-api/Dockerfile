FROM python:3.9-slim

# Define o diretório de trabalho dentro da imagem
WORKDIR /app

# Copiar todos os arquivos e diretórios para a imagem Docker
COPY . /app

# Instalar as dependências necessárias
RUN pip install --no-cache-dir flask pandas numpy scikit-learn

# Definir o comando para iniciar a aplicação
CMD ["python", "app.py"]
