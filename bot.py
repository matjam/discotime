from redis_client import Redis
from discord.ext import commands
from os import getenv
from datetime import datetime
import logging
import asyncio

logger = logging.getLogger(__name__)


bot = commands.Bot(command_prefix="!", description="A bot that converts timezones.")
redis = Redis()


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
    bot.run(getenv("DISCORD_AUTH"))
    logging.info("Finished")


if __name__ == "__main__":
    asyncio.run(main())
