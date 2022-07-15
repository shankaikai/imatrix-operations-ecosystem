import argparse

def parse_arguments():
    parser = argparse.ArgumentParser(description='IoT Gate Control Prototype')
    parser.add_argument('--port', type=str, default="9099",
                        help='Port of the IoT Gate Server')
    args = parser.parse_args()
    return args