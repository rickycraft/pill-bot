#!/bin/bash

git clone https://github.com/rickycraft/pill-bot.git
cd pill-bot/install

echo "install python venv"
python3 -m venv ../env
source ../env/bin/activate
pip install -r requirements.txt
deactivate

# Reading python token
unset TOKEN
read TOKEN -p "Telegram bot token: "
if [ -z "$TOKEN" ]; then
  echo "Token is empty"
  exit 1
fi
echo $TOKEN >../bot_token

echo '{"_default": {"1": {"id": "box","date": "2021-01-01"}}}' >../db.json

echo "install service"
cp pill-bot.service /etc/systemd/system
systemctl daemon-reload
systemctl enable pill-bot
echo "starting service"
systemctl start pill-bot
