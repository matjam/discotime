import aioredis


class Redis:
    client = None

    @classmethod
    async def connect(cls, url):
        cls.client = await aioredis.create_redis_pool(url)

    @classmethod
    def get(cls):
        return cls.client
