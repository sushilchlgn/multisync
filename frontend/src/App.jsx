import { useEffect, useRef, useState } from "react";

function App() {
  const [url, setUrl] = useState("https://example.com");
  const socketRef = useRef(null);

  const safeUrl = url || "https://example.com";

  const devices = [
    { name: "Desktop", width: 1200, height: 700 },
    { name: "Tablet", width: 768, height: 1024 },
    { name: "Android", width: 412, height: 915 },
    { name: "iPhone", width: 390, height: 844 },
  ];

  // ---------------------------
  // WebSocket connection
  // ---------------------------
  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");

    ws.onopen = () => {
      console.log("WebSocket connected");
    };

    ws.onclose = () => {
      console.log("WebSocket closed");
    };

    ws.onerror = (err) => {
      console.log("WebSocket error", err);
    };

    ws.onmessage = (event) => {
      try {
        const msg = JSON.parse(event.data);

        switch (msg.type) {
          case "URL_CHANGE":
            setUrl(msg.data);
            break;

          case "INPUT":
            setUrl(msg.data.value);
            break;

          case "SCROLL":
            window.scrollTo(
              msg.data.x || 0,
              msg.data.y || 0
            );
            break;

          case "CLICK":
            console.log("click event:", msg.data);
            break;

          default:
            break;
        }
      } catch (err) {
        console.log("Invalid WS message", err);
      }
    };

    socketRef.current = ws;

    return () => ws.close();
  }, []);

  // ---------------------------
  // Safe send function
  // ---------------------------
  const sendMessage = (msg) => {
    const socket = socketRef.current;

    if (!socket || socket.readyState !== WebSocket.OPEN) return;

    socket.send(JSON.stringify(msg));
  };

  // ---------------------------
  // URL INPUT SYNC
  // ---------------------------
  const handleUrlChange = (value) => {
    setUrl(value);

    sendMessage({
      type: "INPUT",
      data: { value },
    });
  };

  // ---------------------------
  // SCROLL SYNC (THROTTLED)
  // ---------------------------
  useEffect(() => {
    let lastTime = 0;

    const handleScroll = () => {
      const now = Date.now();

      // throttle to ~20fps
      if (now - lastTime < 50) return;
      lastTime = now;

      sendMessage({
        type: "SCROLL",
        data: {
          x: window.scrollX,
          y: window.scrollY,
        },
      });
    };

    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  // ---------------------------
  // CLICK SYNC
  // ---------------------------
  useEffect(() => {
    const handleClick = (e) => {
      sendMessage({
        type: "CLICK",
        data: {
          x: e.clientX,
          y: e.clientY,
        },
      });
    };

    window.addEventListener("click", handleClick);
    return () => window.removeEventListener("click", handleClick);
  }, []);

  return (
    <div style={{ padding: 20 }}>
      <h1>MultiSync MVP</h1>

      <input
        style={{ width: 400 }}
        value={url}
        onChange={(e) => handleUrlChange(e.target.value)}
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

            <iframe
              src={safeUrl}
              width={d.width}
              height={d.height}
              title={d.name}
              style={{ border: "1px solid #ccc" }}
            />
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;