import { MapContainer } from "react-leaflet";
import { useLeafletContext } from "./context/LeafletContext";
import MapView from "./MapView";

const Leaflet = () => {
  const { markers } = useLeafletContext();

  return (
    <MapContainer
      center={[51.505, -0.09]}
      zoom={2}
      style={{ height: "100%", width: "100%" }}
    >
      <MapView markers={markers} />
    </MapContainer>
  );
};

export default Leaflet;
