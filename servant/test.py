#!/usr/bin/env python
#encoding=utf-8
'''
for quick debugging of fxi
'''

import sys
import datetime
sys.path.append('gen-py')
sys.path.append('/System/Library/Frameworks/Python.framework/Versions/2.7/lib/python2.7/site-packages')
from thrift.transport import TSocket
from thrift.protocol import TBinaryProtocol
from servant import FunServant

t1 = datetime.datetime.now()
sock = TSocket.TSocket('localhost', 9001)
sock.open()
protocol = TBinaryProtocol.TBinaryProtocol(sock)

client = FunServant.Client(protocol)

# ping
#=====
r = client.ping()
delta = datetime.datetime.now() - t1
print '[Client] received from rpc server:', r, delta.microseconds, 'us'

print client.mc_set('hello', 'world 世界', 120)
print client.mc_get('hello')

try:
    #print client.mc_get('hello-non-exist')
    print 'ha'
except:
    pass

print client.lc_set('error_tag', 'abcdefg')
print client.lc_get('error_tag')
