from datetime import date, time, timedelta
from telegram import Update
from telegram.ext import CallbackContext, Job,  MessageHandler, Filters
from db import Database
from bot import Bot


class Reminder:

    def __init__(self, bot: Bot, db: Database):
        self.bot = bot
        self.db = db
        self.messages = bot.cfg.messages
        self.reminder_delta = bot.cfg.bot['reminder_delta']

        self.bot.job_queue.run_daily(lambda _: self.start_reminder(), bot.cfg.reminder_time)
        self.reminder_job = Job(lambda _: print('stop reminder'))

    def start_reminder(self):
        # Return if already taken
        if (self.db.taken):
            return

        self.reminder_job = self.bot.job_queue.run_repeating(
            self.bot.remind, timedelta(minutes=self.reminder_delta), first=timedelta(seconds=1), name="repeating_reminder"
        )

    def stop_reminder(self):
        if (self.reminder_job.job != None):
            self.reminder_job.schedule_removal()

    # ONE TIME REMINDERS

    def get_reminder(self, update: Update, context: CallbackContext):
        jobs = self.bot.job_queue.get_jobs_by_name("onetime_reminder")
        index = 0
        for j in jobs:
            if (not j.removed):
                index += 1
                next_t = j.next_t + timedelta(hours=2)
                next_t = next_t.time().strftime("%H:%M")
                update.message.reply_text(self.messages["next_reminder"]+next_t)
        if (index == 0):
            update.message.reply_text(self.messages["no_reminder"])

    def remove_reminder(self, update: Update, contex: CallbackContext):
        jobs = self.bot.job_queue.get_jobs_by_name("onetime_reminder")
        index = 0
        for j in jobs:
            if (not j.removed):
                j.schedule_removal()
                index += 1
        update.message.reply_text(self.messages['removed_reminder'] % index)

    def add_reminder(self, update: Update, context: CallbackContext):
        hour = int(context.match.group(1))
        try:
            minute = int(context.match.group(2))
        except TypeError:
            minute = 0
        try:
            when = time(hour, minute)  # TODO timezone
            j = self.bot.job_queue.run_once(self.bot.remind, when, name="onetime_reminder")
            if (j.next_t.date() == date.today()):
                self.bot.broadcast(self.messages['reminder_success']+time(hour, minute).strftime("%H:%M"))
            else:
                self.bot.broadcast(self.messages['reminder_tomorrow']+time(hour, minute).strftime("%H:%M"))
        except ValueError:
            update.message.reply_text(self.messages['reminder_fail'])

    def register_handlers(self) -> None:
        self.bot.add_handler(MessageHandler(
            self.bot.auth_filter & Filters.regex(r'reminders'),
            self.get_reminder))
        self.bot.add_handler(MessageHandler(
            self.bot.auth_filter & Filters.regex(r'remove'),
            self.remove_reminder))
        self.bot.add_handler(MessageHandler(
            self.bot.auth_filter & Filters.regex(r'remind (\d?\d):?(\d\d)?'),
            self.add_reminder))
