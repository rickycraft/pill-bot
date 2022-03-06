#!./env/bin/python

from stats import Stats
from telegram import Update
from telegram.ext import CallbackContext, CommandHandler, MessageHandler, Filters
from config import Config
from bot import Bot
from db import Database
from datetime import datetime, timedelta
from reminder import Reminder
from box import Box
from echo import Echo


class Server:

    def __init__(self):
        self.cfg = Config()
        self.db = Database(self.cfg)
        self.bot = Bot(self.cfg)
        self.reminders = Reminder(self.bot, self.db)
        # Adding all handlers
        self.add_handlers()

    def add_handlers(self):
        self.bot.add_handler(CommandHandler('start', self.db.save_info))
        self.bot.add_handler(MessageHandler(self.bot.auth_filter & Filters.regex(r'pill'), self.pill))
        # Reminders
        self.reminders.register_handlers()
        # Custom commands
        self.bot.add_handler(Echo(self.bot).handler)
        self.bot.add_handler(Box(self.db, self.bot).handler)
        self.bot.add_handler(Stats(self.bot, self.db).handler)
        # Catchall
        self.bot.add_handler(MessageHandler(Filters.all, self.default))

    def pill(self, update: Update, context: CallbackContext):
        now = datetime.today().astimezone(self.cfg.zone)
        # check if not early
        if ((self.cfg.start_time - timedelta(hours=1)) <= now <= self.cfg.end_time):
            if (self.db.taken):
                update.message.reply_text(self.cfg.messages['user_already_confirmed'])
            else:
                self.db.save_day()
                self.reminders.stop_reminder()
                self.bot.confirm()
        else:
            update.message.reply_text(self.cfg.messages['early'])

    def default(self, update: Update, context: CallbackContext):
        update.message.reply_text(self.cfg.messages["command_not_supported"])


if __name__ == '__main__':
    print("Creating server", datetime.today().isoformat())
    server = Server()
    server.bot.start()
