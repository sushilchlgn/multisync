import { useState } from "react";

function App() {
  const [url, setUrl] = useState("https://example.com");

  const devices = [
    { name: "Desktop", width: 1200, height: 700 },
    { name: "Tablet", width: 768, height: 1024 },
    { name: "Android", width: 412, height: 915 },
    { name: "iPhone", width: 390, height: 844 },
  ];

  return (
    <div>
      <h1>MultiSync Testing Platform</h1>

      <input
        style={{ width: "500px" }}
        value={url}
        onChange={(e) => setUrl(e.target.value)}
      />

      <div
        style={{
          display: "grid",
          gridTemplateColumns: "1fr 1fr",
          gap: "20px",
          marginTop: "20px",
        }}
      >
        {devices.map((device) => (
          <div key={device.name}>
            <h3>{device.name}</h3>

            <iframe
              src={url}
              width={device.width}
              height={device.height}
              title={device.name}
            />
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;