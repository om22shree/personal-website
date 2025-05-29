FROM python:alpine3.19
WORKDIR /app

COPY requirements.txt .
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

COPY . .

EXPOSE 8080

CMD ["streamlit", "run", "main.py", "--server.port=8080"]