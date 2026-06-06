import { useEffect, useRef, useState } from "react";

function App() {
  const [url, setUrl] = useState("https://example.com");
  const [socket, setSocket] = useState(null);
  const socketRef = useRef(null);

  const devices = [
    { name: "Desktop", width: 1200, height: 700 },
    { name: "Tablet", width: 768, height: 1024 },
    { name: "Android", width: 412, height: 915 },
    { name: "iPhone", width: 390, height: 844 },
  ];

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");

    ws.onopen = () => {
      console.log("WebSocket connected");
    };

    ws.onmessage = (event) => {
      const msg = JSON.parse(event.data);

      if (msg.type === "URL_CHANGE") {
        setUrl(msg.data);
      }
    };

    socketRef.current = ws;

    return () => ws.close();
  }, []);

  const sendUrl = (newUrl) => {
    setUrl(newUrl);

    socketRef.current?.send(
      JSON.stringify({
        type: "URL_CHANGE",
        data: newUrl,
      })
    );
  };
  return (
    <div style={{ padding: 20 }}>
      <h1>MultiSync MVP</h1>

      <input
        style={{ width: 400 }}
        value={url}
        onChange={(e) => sendUrl(e.target.value)}
      />

      <div
        style={{
          display: "grid",
          gridTemplateColumns: "1fr 1fr",
          gap: 20,
          marginTop: 20,
        }}
      >
        {devices.map((d) => (
          <div key={d.name}>
            <h3>{d.name}</h3>
            <iframe src={url} width={d.width} height={d.height} />
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;