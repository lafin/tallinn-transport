import React, { useEffect, useState } from "react";
import { MapContainer, TileLayer, CircleMarker } from "react-leaflet";
import { preparePointsToMap, getData } from "./providers";
import { Point } from "./interfaces";
import "./App.css";
import "leaflet/dist/leaflet.css";

const App: React.FunctionComponent = () => {
  const [init] = useState({
    lat: 59.436962,
    lng: 24.753574,
    zoom: 13,
  });
  const [points, setPoints] = useState<Point[]>([]);

  useEffect(() => {
    addData();
    setInterval(() => addData(), 2000);
  }, []);

  async function addData() {
    setPoints(await preparePointsToMap(getData()));
  }

  function getColor(type: number) {
    switch (type) {
      case 1:
        return "blue";
      case 2:
        return "green";
      case 3:
        return "red";
      case 4:
        return "orange";
    }
  }

  return (
    <MapContainer center={[init.lat, init.lng]} zoom={init.zoom}>
      <TileLayer
        attribution='Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors, <a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, Imagery © <a href="https://www.mapbox.com/">Mapbox</a>'
        url="https://api.mapbox.com/styles/v1/{id}/tiles/{z}/{x}/{y}?access_token={accessToken}"
        id="mapbox/light-v10"
        accessToken="pk.eyJ1IjoibGFmaW4iLCJhIjoiY2tjZ203MXljMGM2ajJ5cWszMG1wY2ExdCJ9.cPdi63ptWIqYt6MRLMseVg"
      />
      <>
        {points.map((point) => (
          <CircleMarker
            key={`marker-${point.vehicleType}-${point.vehicleNumber}`}
            center={[point.latitude, point.longitude]}
            radius={3}
            fill={true}
            weight={1}
            color={getColor(point.vehicleType)}
          />
        ))}
      </>
    </MapContainer>
  );
};

export default App;
