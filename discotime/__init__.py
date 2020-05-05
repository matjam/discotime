from discord.ext import commands
from os import getenv

bot = commands.Bot(command_prefix="!", description="A bot that converts timezones.")


@bot.event
async def on_ready():
    print(f"Logged in as {bot.user.name} ({bot.user.id})")


@bot.command()
async def time(ctx):
    print(f"tz command {ctx}")


bot.run(getenv("DISCORD_AUTH"))

bot.add_command(time)
