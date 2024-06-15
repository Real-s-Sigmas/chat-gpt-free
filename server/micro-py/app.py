import os
import uuid
import psycopg2
from psycopg2 import extras, Error
from flask import Flask, jsonify, request, session, make_response, send_from_directory
from flask_cors import CORS
import smtplib
from email.mime.text import MIMEText
import random
from datetime import datetime
from dotenv import load_dotenv
import base64
import logging
import asyncio

load_dotenv()

logging.basicConfig(
    level=logging.DEBUG,
    format="%(asctime)s %(levelname)s %(message)s",
    datefmt="%Y—%m—%d %H:%M:%S",
)

PASSWORD_PG = os.getenv('PASSWORD_PG')
PORT_PG = os.getenv('PORT_PG')
USER_PG = os.getenv('USER_PG')
HOST_PG = os.getenv('HOST_PG')
DB_NAME_PG = os.getenv('DB_NAME_PG')


def escape_quotes(text):
    return text.replace("'", "''")

def unescape_quotes(text):
    return text.replace("''", "'")


app = Flask(__name__)

app.secret_key = "/zxc/"
app.permanent_session_lifetime = 60 * 60 * 24 * 28
app.config["SESSION_COOKIE_SAMESITE"] = "None"
app.config["SESSION_COOKIE_SECURE"] =  'None'

# enable CORS
CORS(app, resources={r"*": {"origins": "*", 'supports_credentials': True}})

#Главная страница
@app.route('/', methods=['GET'])
def home():

    response_object = {'status': 'success'} #БаZа
    response_object['message'] = session.get('id')
    logging.warning('1')
    # logging.info(session.get('id')) #debug
    logging.warning(response_object)
    session.pop('id', None)
    return jsonify(response_object)

def add_tables():
    try:
        pg = psycopg2.connect(f"""
                host={HOST_PG}
                dbname={DB_NAME_PG}
                user={USER_PG}
                password={PASSWORD_PG}
                port={PORT_PG}
            """)
        cursor = pg.cursor(cursor_factory=psycopg2.extras.DictCursor)

        
    except (Exception, Error) as error:
        logging.error(f'DB: ', error)
        return_data = f"Error" 

    finally:
        if pg:
            cursor.close
            pg.close
            logging.info("Соединение с PostgreSQL закрыто")


if __name__ == '__main__':
      add_tables()
      app.run()


