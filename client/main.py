from __future__ import print_function

import logging
import time

import grpc
# from proto import deep_pb2_grpc as proto, deep_pb2 as pt # proto type
import proto.deep_pb2 as pt
import proto.deep_pb2_grpc as proto


def pass_num_with_one_second_delay():
    nums = [pt.Num(a=i+1) for i in range(10)]
    for num in nums:
        time.sleep(1)
        yield num

def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel("localhost:6969") as channel:
        stub = proto.DeepSrvStub(channel)
        response = stub.AddTwoNum(pt.NumReq(a=1, b=2))
        print("AddTwoNum response received: %s" % response)
        print("="*100)
        
        respose = stub.AddAllTheseNum(pass_num_with_one_second_delay())
        print("AddAllTheseNum response received: %s" % respose)
    


if __name__ == "__main__":
    logging.basicConfig()
    run()
