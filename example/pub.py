import sys
import argparse
import zmq
from random import randrange


def _pub(args):
    #  Prepare our context and sockets
    context = zmq.Context()
    socket = context.socket(zmq.PUB)
    socket.connect("tcp://localhost:5569")

    while True:
        topic = randrange(1, 100000)
        temperature = randrange(-80, 135)
        relhumidity = randrange(10, 60)
        socket.send_string("%i %i %i" % (topic, temperature, relhumidity))


def _parser_args(params):
    """解析命令行参数."""
    parser = argparse.ArgumentParser()
    parser.add_argument('--url', type=str, default="tcp://localhost:5559", help="指定连接到哪个组件")
    parser.set_defaults(func=_client)
    args = parser.parse_args(params)
    args.func(args)


def main(argv=sys.argv[1:]):
    u"""服务启动入口.

    设置覆盖顺序`命令行参数`>`'-c'指定的配置文件`>`项目启动位置的配置文件`>默认配置.
    """
    _parser_args(argv)


if __name__ == '__main__':
    main()
