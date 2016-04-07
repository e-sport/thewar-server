import websocket
import time
import struct

ws = websocket.WebSocket()
ws.connect("ws://localhost:8080")
ws.send(struct.pack("!i", 1003) + 'clientaaaaaaaaaaaaaaaaaaaaaa')
for i in range(1, 10):
	result =  ws.recv()
	print result
	# ws.send('clientaaaaaaaaaaaaaaaaaaaaaa')


ws.close()