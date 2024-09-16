from __future__ import print_function

import logging

import grpc
# from proto import deep_pb2_grpc as proto, deep_pb2 as pt # proto type
import proto.deep_pb2 as pt
import proto.deep_pb2_grpc as proto


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel("localhost:6969") as channel:
        stub = proto.DeepSrvStub(channel)
        response = stub.AddTwoNum(pt.NumReq(a=1, b=2))
    print("AddTwoNum response received: %s" % response)
    print("Will try to greet world ...")


if __name__ == "__main__":
    logging.basicConfig()
    run()
