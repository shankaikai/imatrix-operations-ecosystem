from __future__ import annotations
from queue import Queue

from . import gate_control_request

import threading

from Proto import iot_prototype_pb2

# Publisher that publishes changes in the gate state
class GateControlPublisher:
    # A dictionary of all subscribers
    # key: thread id
    # value: tuple(event objects for notification, event queue)
    subscribers : dict[int, tuple[threading.Event, Queue]] = {}

    # The initial state of the gate
    gate_state = iot_prototype_pb2.GateState.CLOSED
    
    def __init__(self, gate_lock : threading.Lock) -> None:
        self.gate_lock = gate_lock

    def subscribe(self, thread_id : int, event_object, queue : Queue):
        self.subscribers[thread_id] = (event_object, queue)

    def unsubscribe(self, thread_id):
        del self.subscribers[thread_id]

    def publish(self):
        self.gate_lock.acquire()
        # Let all the subscribers know that there is a state change.
        for event, queue in self.subscribers.values():
            new_request = gate_control_request.GateControlRequest(
                request_type=gate_control_request.GateControlRequest.STATE_CHANGE_KEYWORD,
                new_gate_state=self.gate_state)
            queue.put(new_request)
            # Notify
            event.set()
        self.gate_lock.release()
        
    
    def open_close_gate(self, gate_control_request : gate_control_request.GateControlRequest):
        print("changing gate to", gate_control_request.new_gate_state)
        self.gate_lock.acquire()
        self.gate_state = gate_control_request.new_gate_state
        self.gate_lock.release()
        print("new state", self.gate_state)
        self.publish()
