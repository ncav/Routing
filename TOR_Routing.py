import socks
import socket
import stem.process
from stem.util import term

# Start a Tor proxy on the default port
def start_tor():
    tor_process = stem.process.launch_tor_with_config(
        config={
            'SocksPort': '9050',
        },
        init_msg_handler=print,
    )
    return tor_process

# Make a request using the Tor proxy
def make_request(url):
    socks.setdefaultproxy(socks.PROXY_TYPE_SOCKS5, '127.0.0.1', 9050)
    socket.socket = socks.socksocket
    response = requests.get(url)
    return response

# Start the Tor proxy
tor_process = start_tor()
print(term.format('Tor proxy started', term.Color.GREEN))

# Make a request using the Tor proxy
response = make_request('https://example.com')
print(response.content)

# Stop the Tor proxy
tor_process.kill()
print(term.format('Tor proxy stopped', term.Color.GREEN))
