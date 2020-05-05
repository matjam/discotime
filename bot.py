from discord.ext import commands
from os import getenv
from datetime import datetime


bot = commands.Bot(command_prefix="!", description="A bot that converts timezones.")


@bot.event
async def on_ready():
    print(f"Logged in as {bot.user.name} ({bot.user.id})")


@bot.command()
async def time(ctx):
    now = datetime.now()
    current_time = now.strftime("%H:%M:%S")
    await ctx.send(current_time)


bot.run(getenv("DISCORD_AUTH"))
