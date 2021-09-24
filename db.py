from tinydb import TinyDB, Query, where
from datetime import time, timedelta, date, datetime
from statistics import mean, stdev
import math
from config import Config


class Database:

    def __init__(self, cfg: Config):
        self.db = TinyDB('db.json')
        self.users = self.db.table('users')
        self.entries = self.db.table('entries', cache_size=0)
        self.query = Query()
        self.cfg = cfg

    def save_day(self):
        """Save when is pill is takens"""
        today = datetime.today().replace(microsecond=0)
        entry = {
            'date': today.date().isoformat(),
            'time': today.time().isoformat(),
            'dt': today.isoformat()
        }
        return self.entries.insert(entry)

    def save_info(self, update, contex):
        """Logs info of anywan that uses the bots"""
        update.message.reply_text("THIS IS NOT FOR YOU")
        _user = update.effective_user
        _new_user = {
            'id': _user.id,
            'name': _user.full_name,
            'username': _user.name,
            'datetime': datetime.today().isoformat()
        }
        print("new user", _new_user)
        NewUser = Query()
        self.users.upsert(_new_user, NewUser.id == _user.id)

    @property
    def taken(self):
        return self.entries.get(Query().dt.test(self.in_between))

    def in_between(self, dt: str):
        return (self.cfg.start_time <= datetime.fromisoformat(dt).astimezone(self.cfg.zone) <= self.cfg.end_time)

    @ property
    def box_day(self):
        """Get the date of the first day of the pill """
        day0 = date.fromisoformat(
            self.db.get(Query()['id'] == 'box')['date']
        )

        delta_day = (date.today() - day0).days
        if (delta_day < 28):
            return day0
        else:
            day0 += timedelta(days=28)
            self.db.update(
                {'date': day0.isoformat()},
                Query()['id'] == 'box'
            )
            return day0

    def month_entries(self, month: int):
        return self.entries.search(
            lambda e: date.fromisoformat(e['date']).month == month
        )
