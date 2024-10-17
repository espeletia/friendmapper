import { useState } from "react";
import { MapContainer } from "react-leaflet";
import { Marker } from "../../types";
import SidebarInfo from "./components/SidebarInfo";
import { useLeafletContext } from "./context/LeafletContext";
import MapView from "./MapView";

const Leaflet = () => {
  const { markers, updateMarkersDebounce } = useLeafletContext();

  const [activeMarker, setActiveMarker] = useState<Marker | null>(null);

  const handleClick = (marker: Marker) => {
    setActiveMarker(marker);
  };

  const closeCallback = () => {
    setActiveMarker(null);
  };

  return (
    <MapContainer
      center={[50.209722, 15.830473]}
      zoom={13}
      style={{ height: "100%", width: "100%" }}
      bounds={[
        [50.24807831738873, 15.963821411132814],
        [50.16446904866322, 15.70426940917969],
      ]}
    >
      {activeMarker && (
        <SidebarInfo closeCallback={closeCallback} marker={activeMarker} />
      )}
      <MapView
        handleClick={handleClick}
        markers={markers}
        updateMarkers={updateMarkersDebounce}
      />
    </MapContainer>
  );
};

export default Leaflet;
