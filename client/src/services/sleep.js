import { reactive } from 'vue';
import { step_state } from './step';
import { ip } from 'src/ip_adress';
export const sleep_state = reactive({
  connected: false,
  data: [],
});

class WebSocketService {
  constructor() {
    this.socket = null;
  }

  connect(url, messageHandler) {
    this.messageHandler = messageHandler;
    this.socket = new window.WebSocket(url)
    this.socket.onopen = () => {
      sleep_state.connected = true
      console.log('WebSocket connection established');
    };

    this.socket.onmessage = (event) => {
      sleep_state.data.push(JSON.parse(event.data))
    };

    this.socket.onclose = () => {
      step_state.connected = false
      console.log('WebSocket connection closed');
    };

    this.socket.onerror = (error) => {
      console.error('WebSocket error:', error);
    };
  }

  send(message) {
    if (this.socket && this.socket.readyState === WebSocket.OPEN) {
      this.socket.send(message);
    } else {
      console.error('WebSocket connection is not open');
    }
  }

  close() {
    if (this.socket) {
      this.socket.close();
    }
  }
}


function handleMessage(message) {
  console.log('Received message:', message);

}
var webSocket = new WebSocketService();
webSocket.connect('ws://' + ip + ':3000/ws/sleep/1', handleMessage);


