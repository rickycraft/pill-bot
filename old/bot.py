from telegram import ReplyKeyboardMarkup
from telegram.ext import Updater, Handler, Filters
from telegram.ext.callbackcontext import CallbackContext
from config import Config


class Bot:

    @staticmethod
    def get_key(filename: str):
        with open(filename, 'r') as t:
            return t.readline().strip()

    def __init__(self, cfg: Config):
        self.cfg = cfg
        self.updater = Updater(Bot.get_key(cfg.bot["token_file"]), use_context=True)
        self.bot = self.updater.bot
        self.dp = self.updater.dispatcher
        self.job_queue = self.updater.job_queue
        self.job_queue.set_dispatcher(self.dp)
        self.job_queue.start()

        self.user_keyboard = ReplyKeyboardMarkup(self.cfg.bot['user_keyboard'], resize_keyboard=True)
        self.admin_keyboard = ReplyKeyboardMarkup(self.cfg.bot['admin_keyboard'], resize_keyboard=False)

    def start(self):
        self.notify_admin("Starting the bot")
        self.updater.start_polling()
        self.updater.idle()

    def add_handler(self, handler: Handler):
        self.dp.add_handler(handler)

    @property
    def user_filter(self):
        return Filters.user(self.cfg.user)

    @property
    def admin_filter(self):
        return Filters.user(self.cfg.admin)

    @property
    def auth_filter(self):
        return (self.user_filter | self.admin_filter)

    def notify_admin(self, msg: str):
        self.bot.send_message(self.cfg.admin, msg, reply_markup=self.admin_keyboard)

    def notify_user(self, msg: str):
        self.bot.send_message(self.cfg.user, msg, reply_markup=self.user_keyboard)

    def broadcast(self, msg: str):
        self.notify_admin(msg)
        self.notify_user(msg)

    def remind(self, context: CallbackContext):
        self.notify_admin(self.cfg.messages['admin_reminder'])
        self.notify_user(self.cfg.reminder_msg)

    def confirm(self):
        self.notify_admin(self.cfg.messages['admin_confirm'])
        self.notify_user(self.cfg.confirm_msg)
