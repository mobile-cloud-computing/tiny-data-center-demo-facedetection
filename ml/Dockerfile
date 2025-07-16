FROM arm32v7/python:3.11.13

RUN apt-get update && apt-get install -y \
    libgl1 \
    && rm -rf /var/lib/apt/lists/*
    
WORKDIR /app
COPY ./req.txt /app/
RUN pip install --no-cache-dir -r req.txt

COPY ./main.py /app
EXPOSE 8000

ENTRYPOINT ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]