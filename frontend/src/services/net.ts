export class NetService {
    private websocket!: WebSocket;
    private textDecoder = new TextDecoder();

    connect() {
        this.websocket = new WebSocket("ws://localhost:3000/ws");
        this.websocket.onopen = () => {
            console.log("connected");
        }

        this.websocket.onmessage = async (event: MessageEvent) => {
            const arrayBuffer = event.data.arrayBuffer();
            const bytes = new Uint8Array(arrayBuffer);
            const packetId = bytes[0];

            const packet = JSON.parse(this.textDecoder.decode(bytes.subarray(1)));
        }
    }
}