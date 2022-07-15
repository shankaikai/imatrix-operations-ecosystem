from Proto import iot_prototype_pb2

class GateControlRequest:
    STATE_CHANGE_KEYWORD = "STATE_CHANGE"
    MODIFY_REQUEST_KEYWORD = "MODIFY_REQUEST"

    def __init__(self, request_type : str, new_gate_state : iot_prototype_pb2.GateState = None) -> None:
        self.request_type = request_type
        self.new_gate_state = new_gate_state
