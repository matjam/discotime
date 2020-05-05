from discord.ext import commands
from os import getenv
from datetime import datetime
import logging

logger = logging.getLogger(__name__)


bot = commands.Bot(command_prefix="!", description="A bot that converts timezones.")


@bot.event
async def on_ready():
    print(f"Logged in as {bot.user.name} ({bot.user.id})")


@bot.command()
async def time(ctx):
    now = datetime.now()
    current_time = now.strftime("%H:%M:%S")
    await ctx.send(current_time)


def main():
    logging.basicConfig(
        format="%(asctime)s - %(name)s - %(levelname)s - %(message)s",
        level=logging.INFO,
    )

    logging.info("Started")
    bot.run(getenv("DISCORD_AUTH"))
    logging.info("Finished")


if __name__ == "__main__":
    main()
