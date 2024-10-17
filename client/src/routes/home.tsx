import { useEffect } from "react";
import type { Marker } from "../types";
import { useLeafletContext } from "../components/Leaflet/context/LeafletContext";

const HomeRoute = () => {
  const { setMarkers } = useLeafletContext();

  useEffect(() => {
    const fetchData = () => {
      const markerData: Marker[] = [
        { name: "Location 1", lat: 51.505, lng: -0.09 },
        { name: "Location 2", lat: 51.51, lng: -0.1 },
        { name: "Location 3", lat: 51.515, lng: -0.11 },
        { name: "Location 4", lat: 51.52, lng: -0.12 },
      ];
      setMarkers(markerData);
    };

    fetchData();
  }, []);

  return <div>root</div>;
};

export default HomeRoute;
