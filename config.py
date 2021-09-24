from yaml import safe_load
from random import randint
from datetime import time, datetime, date, timedelta
from zoneinfo import ZoneInfo

CONFIG_FILE_NAME = "cfg.yaml"


class Config:

    def __init__(self):
        with open(CONFIG_FILE_NAME, 'r') as config_file:
            self.doc = safe_load(config_file)
            self.bot = self.doc['bot']
            self.messages = self.doc['messages']

            self.admin = self.bot['admin']
            self.user = self.bot['user']
            self.zone = ZoneInfo(self.bot["timezone"])

    @property
    def reminder_msg(self):
        reminders = self.messages['reminders']
        index = randint(0, len(reminders) - 1)
        return reminders[index]

    @property
    def confirm_msg(self):
        confirms = self.messages['confirms']
        index = randint(0, len(confirms) - 1)
        return confirms[index]

    @property
    def start_time(self):
        d = date.today()
        t = time(self.bot["start_time"])
        return datetime.combine(d, t, tzinfo=self.zone)

    @property
    def end_time(self):
        d = date.today()+timedelta(days=1)
        t = time(self.bot["end_time"])
        return datetime.combine(d, t, tzinfo=self.zone)

    @property
    def reminder_time(self):
        return (self.start_time - self.zone.utcoffset(self.start_time)).time()

    @property
    def weeks(self):
        return self.bot['weeks']
