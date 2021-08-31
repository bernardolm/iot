import logging
import os
import sys
import time

from flask import Flask
from github_webhook import Webhook


class Updater():

    def __init__(self, event=None):
        self._event = event

    def do(self):
        app = Flask(__name__)
        webhook = Webhook(app)

        @webhook.hook()
        def on_push(data):
            logging.debug('got github push')
            self._event.set()

        app.run(host="0.0.0.0", port=int(os.environ.get('UPDATER_PORT', 8234)))
