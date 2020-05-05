from redis_client import Redis
from discord.ext import commands
from os import getenv
from datetime import datetime
import logging
from asyncio import get_event_loop

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


async def main():
    logging.basicConfig(
        format="%(asctime)s - %(name)s - %(levelname)s - %(message)s",
        level=logging.INFO,
    )

    logging.info("Starting")

    # and having to know this is async is why this sucks
    await Redis.connect(getenv("REDIS_URL"))

    try:
        bot.start(getenv("DISCORD_AUTH"))
    except Exception as e:
        bot.close()
        raise e

    logging.info("Finished")


if __name__ == "__main__":
    loop = get_event_loop()
    loop.run_until_complete(main())
