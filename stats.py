
import math
from telegram.ext.callbackcontext import CallbackContext
from telegram.ext.commandhandler import CommandHandler
from telegram.ext.handler import Handler
from db import Database
from bot import Bot
from telegram import Update
from statistics import StatisticsError, mean, stdev
from datetime import date, time


class Stats:

    command = "stats"

    def __init__(self, bot: Bot, db: Database) -> None:
        self.bot = bot
        self.db = db

    # TODO adjust for timezone
    def month_avg(self, month: int):
        entries = self.db.month_entries(month)
        if len(entries) == 0:
            raise IndexError("There are no data points")
        # Map the time
        avg_times = [time.fromisoformat(e['time']) for e in entries]
        # Account only for hour and minute  and discard values too far from target
        avg_times = [(t.hour * 60 + t.minute) for t in avg_times if t.hour >= 19]
        avg_time = mean(avg_times)
        dev = int(stdev(avg_times, avg_time))

        hour = math.floor(avg_time / 60)
        minute = math.floor(avg_time % 60)
        return (time(hour, minute, 0), dev)

    def action(self, update: Update, context: CallbackContext):
        if (len(context.args) == 1):
            arg0 = int(context.args[0])
        else:
            arg0 = date.today().month
        try:
            m_avg = self.month_avg(arg0)
            self.bot.notify_admin(
                f'Avg time {m_avg[0].isoformat()}\nStdev { m_avg[1] }min'
            )
        except (StatisticsError, IndexError) as stat_error:
            self.bot.notify_admin(str(stat_error))

    @property
    def handler(self) -> Handler:
        return CommandHandler(self.command, self.action, self.bot.admin_filter)
