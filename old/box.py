from telegram.ext.callbackcontext import CallbackContext
from telegram.ext.filters import Filters
from telegram.ext.handler import Handler
from telegram.ext.messagehandler import MessageHandler
from telegram.update import Update
from bot import Bot
from datetime import date, timedelta
from db import Database


class Box:

    command = r'box'

    def __init__(self, db: Database, bot: Bot):
        self.taken = db.taken
        self.weeks = bot.cfg.weeks
        self.bot = bot

        self.day0 = db.box_day
        today = date.today()
        self.curr_day = (today - self.day0).days

    def to_text(self):
        n_days = 7 * self.weeks
        line_len = 8
        ret = "➖"*line_len + "\n"
        arr = self.to_array()

        for i in range(n_days):
            if (arr[i]):
                ret += '|⭕'
            else:
                ret += '|⚪'

            if ((i + 1) % 7 == 0):
                ret += "|\n"
                ret += "➖"*line_len + "\n"
        return ret

    def to_array(self):
        wkday = self.day0.weekday()
        n_days = 7 * self.weeks
        n_days_passed = n_days if self.box_index > n_days else self.box_index
        arr = [False] * n_days
        i = wkday
        for _ in range(n_days_passed):
            if i > (n_days - 1):
                i = 0
            arr[i] = True
            i += 1
        return arr

    @property
    def box_index(self):
        adj = 1 if self.taken else 0
        return self.curr_day + adj

    def __str__(self):
        msg = f"Day { self.curr_day + 1}/28\n"
        return msg

    def action(self, update: Update, context: CallbackContext):
        update.message.reply_text(str(self))
        update.message.reply_text(self.to_text())

    @property
    def handler(self) -> Handler:
        return MessageHandler(
            self.bot.auth_filter & Filters.regex(self.command),
            self.action)
