from telegram.ext.handler import Handler
from bot import Bot
from telegram.ext import CallbackContext, CommandHandler
from telegram import Update
from datetime import datetime


class Echo:

    command = "echo"

    def __init__(self, bot: Bot):
        self.bot = bot

    def action(self, update: Update, context: CallbackContext) -> None:
        update.message.reply_text(datetime.today().isoformat())

    @property
    def handler(self) -> Handler:
        return CommandHandler(self.command, self.action, self.bot.admin_filter)
