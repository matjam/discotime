import aioredis
from os import getenv


class Redis:
    redis = None

    def __init__(self):
        self.redis = await aioredis.create_redis_pool(getenv("REDIS_URL"))

    def get_client(self):
        return self.redis
