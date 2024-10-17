import { MapContainer } from "react-leaflet";
import { useLeafletContext } from "./context/LeafletContext";
import MapView from "./MapView";

const Leaflet = () => {
  const { markers, updateMarkersDebounce } = useLeafletContext();

  return (
    <MapContainer
      center={[50.209722, 15.830473]}
      zoom={13}
      style={{ height: "100%", width: "100%" }}
      bounds={[[50.24807831738873, 15.963821411132814], [50.16446904866322, 15.70426940917969]]}
    >
      <MapView markers={markers} updateMarkers={updateMarkersDebounce} />
    </MapContainer>
  );
};

export default Leaflet;
