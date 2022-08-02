FROM python:3.9
WORKDIR /telebot

COPY ./requirements.txt .

RUN pip3 install -r requirements.txt

CMD ["python","-u", "main.py"]