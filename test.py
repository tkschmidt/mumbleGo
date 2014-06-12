#!/usr/bin/env python

from struct import *
import socket, sys, time, datetime
fixed_time=239179
datum = fixed_time
buf = pack(">iQ", 0, datum)
buf
'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\xa6K'