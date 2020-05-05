#!/usr/bin/env python

from setuptools import setup

setup(
    name="discotime",
    version="1.0",
    description="A discord timezone conversion bot",
    author="Nathan Ollerenshaw",
    author_email="chrome@stupendous.net",
    url="https://github.com/matjam/discotime/",
    install_requires=["discord.py", "dateparser"],
    packages=["discotime"],
)
