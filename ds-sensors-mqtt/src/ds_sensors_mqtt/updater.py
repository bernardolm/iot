import logging
import os
import sys
import time

from flask import Flask
from github_webhook import Webhook


class Updater():

    def __init__(self):
        pass

    def do(self, event=None):
        app = Flask(__name__)

        @app.route("/")        # Standard Flask endpoint
        def ping():
            logging.debug('got ping')
            event.set()
            return 'OK'

        webhook = Webhook(app)

        @webhook.hook()
        def on_push(data):
            logging.debug('got github push')
            event.set()
            return 'OK'

        port = int(os.environ.get('UPDATER_PORT', '9345'))
        logging.debug(f'running http server to updater in port {port}')

        app.run(host="0.0.0.0", port=port)
