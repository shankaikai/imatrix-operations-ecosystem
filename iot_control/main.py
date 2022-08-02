from parse_args import parse_arguments
import grpcServer

def main():
    configs = parse_arguments()
    print("configs", configs)
    grpcServer.serve(configs.port)

if __name__ == "__main__":
    main()