const ws = new WebSocket("ws://localhost:8080/ws/cat021-track");

ws.onopen = () => {
  console.log("connected to websocket localhost:8080");
};

ws.onmessage = (event) => {
  console.log("message received", event.data);
};

ws.onerror = (error) => {
  console.log("WebSocket error: " + error.message);
};

ws.onclose = () => {
  console.log("Disconnected from WebSocket server");
};
