from concurrent import futures
from queue import Queue
import threading
import typing
import grpc

from regex import E

from Proto import iot_prototype_pb2_grpc, iot_prototype_pb2
from gate_control import gate_control_publisher, gate_control_request

GATE_ID = 1
FIRE_ALARM_ID = 1


class IotControlPrototypeServiceServicer(iot_prototype_pb2_grpc.IotControlPrototypeServiceServicer):
    """Provides methods that implement functionality of Iot Gate Server."""
    
    def __init__(self, gate_lock : threading.Lock, 
            fire_alarm_lock : threading.Lock, 
            gate_publisher: gate_control_publisher.GateControlPublisher):
       
        # Gate attributes
        self.gate_lock = gate_lock
        self.gate_publisher = gate_publisher
        self.gate_state_event = threading.Event()
        self.gate_request_queue = Queue()

        # Fire Alarm attributes
        self.fire_alarm_lock = fire_alarm_lock


    def GetGateState(self, request : iot_prototype_pb2.Gate, context) -> iot_prototype_pb2.GateState:
        print("Received Request:", request)
        print("request id: ", request.id)
        if request.id == GATE_ID:
            # Get initial current state
            state = iot_prototype_pb2.GateState(
                id=request.id, 
            )
            self.gate_lock.acquire()
            state.state = self.gate_publisher.gate_state
            self.gate_lock.release()
            yield state

            # Subscribe to gate control event
            self.gate_publisher.subscribe(threading.get_ident(), self.gate_state_event, self.gate_request_queue)
            yield from self.ListenToGateEvents(request.id)

            # Unsubscribe from gate control event
            self.gate_publisher.unsubscribe()
        else:
            # No such stepper exists, error
            state = iot_prototype_pb2.GateState(
                id=request.id, 
                state=iot_prototype_pb2.GateState.ERROR
            )
            yield state

    def SetGateState(self,
        request_iterator: typing.Iterator[iot_prototype_pb2.GateState],
        context: grpc.ServicerContext,
    ) -> typing.Iterator[iot_prototype_pb2.GateState]:
        print("connected")
        print(request_iterator)
        # for init_request in request_iterator:
        init_request = request_iterator.__next__()
        print("init_request:", init_request)
        if init_request.id == GATE_ID:
            # Get initial current state
            state = iot_prototype_pb2.GateState(
                id=init_request.id, 
            )
            self.gate_lock.acquire()
            state.state = self.gate_publisher.gate_state
            self.gate_lock.release()
            yield state
            
            # Start a new thread to listen to the actual client requests
            client_listener_thread = threading.Thread(
                target = self.ListenToClientRequests, 
                args = (request_iterator, self.gate_state_event)
            )
            client_listener_thread.start() 
            # Waiting for both the threads to finish executing 
            # Subscribe to gate control event
            self.gate_publisher.subscribe(threading.get_ident(), self.gate_state_event, self.gate_request_queue)
            yield from self.ListenToGateEvents(init_request.id)
            # Unsubscribe from gate control event
            self.gate_publisher.unsubscribe(threading.get_ident())

        else:
            # No such stepper exists, error
            state = iot_prototype_pb2.GateState(
                id=init_request.id, 
                state=iot_prototype_pb2.GateState.ERROR
            )
            yield state

    def GetFireAlarmState(self, request, context):
        print("Received Request:", request)
        print("request id: ", request.id)
        if request.id == GATE_ID:
            # Get initial current state
            state = iot_prototype_pb2.GateState(
                id=request.id, 
            )
            self.gate_lock.acquire()
            state.state = self.gate_publisher.gate_state
            self.gate_lock.release()
            yield state

            # Subscribe to gate control event
            self.gate_publisher.subscribe(threading.get_ident(), self.gate_state_event)
            yield from self.ListenToGateEvents(request.id)

            # Unsubscribe from gate control event
            self.gate_publisher.unsubscribe()
        else:
            # No such stepper exists, error
            state = iot_prototype_pb2.GateState(
                id=request.id, 
                state=iot_prototype_pb2.GateState.ERROR
            )
            yield state
        return iot_prototype_pb2.FireAlarmState()

    def ListenToGateEvents(self, gate_id):
        try:
            while(True):
                self.gate_state_event.wait()
                self.gate_state_event.clear()

                while not self.gate_request_queue.empty():
                    # Get the next gate event in the queue
                    next_request : gate_control_request = self.gate_request_queue.get()
                    if next_request.request_type == gate_control_request.GateControlRequest.MODIFY_REQUEST_KEYWORD:
                        print("modifying gate state", next_request)
                        self.gate_publisher.open_close_gate(next_request)
                    else:
                        state = iot_prototype_pb2.GateState(
                            id=gate_id, 
                        )
                        self.gate_lock.acquire()
                        state.state = self.gate_publisher.gate_state
                        self.gate_lock.release()
                        yield state
        except:
            print("excepted")

    def ListenToClientRequests(self, 
        request_iterator: typing.Iterator[iot_prototype_pb2.GateState], 
        gate_event : threading.Event):
        for request in request_iterator:
            new_request = gate_control_request.GateControlRequest(
                request_type=gate_control_request.GateControlRequest.MODIFY_REQUEST_KEYWORD,
                new_gate_state=request.state)
            self.gate_request_queue.put(new_request)
            gate_event.set()