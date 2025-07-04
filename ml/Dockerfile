# FROM python:3.9.6-slim

# RUN apt-get update && apt-get install -y \
#     libgl1 \
#     libglib2.0-0 \
#     vim \
#     && rm -rf /var/lib/apt/lists/*
    
# WORKDIR /app
# COPY ./req.txt /app/
# RUN pip install --no-cache-dir -r req.txt

# COPY ./main.py /app
# EXPOSE 8000

# ENTRYPOINT ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]

# uvicorn main:app --host 0.0.0.0 --port 8001


FROM python:3.11-slim

RUN apt-get update && apt-get install -y \
    libgl1 \
    libglib2.0-0 \
    vim \
    && rm -rf /var/lib/apt/lists/*
    
WORKDIR /app
COPY ./req.txt /app/
RUN pip install --no-cache-dir -r req.txt

COPY ./main.py /app
EXPOSE 8000

ENTRYPOINT ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]

# uvicorn main:app --host 0.0.0.0 --port 8001
